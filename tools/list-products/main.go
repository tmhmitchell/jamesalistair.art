package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/tmhmitchell/jamesalistair/tools/list-products/posts"
	"github.com/tmhmitchell/jamesalistair/tools/list-products/products"
)

func kebabCase(s string) string {
	lowerWords := make([]string, 0)
	for _, word := range strings.Split(s, " ") {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}

	return strings.Join(lowerWords, "-")
}

func productToPostFileName(p products.Product) string {
	return p.CreatedAt.Format("2006-01-02") + "-" + kebabCase(p.Title) + ".md"
}

func main() {
	var postsDirPath string
	{
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		postsDirPath = path.Join(cwd, "_posts")
	}

	postFileNames := make(map[string]struct{})
	{

		entries, err := os.ReadDir(postsDirPath)
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			postFileNames[entry.Name()] = struct{}{}
		}
	}

	products, err := products.GetProducts()
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		postFileName := productToPostFileName(product)

		if _, ok := postFileNames[postFileName]; ok {
			continue
		}

		postFilePath := path.Join(postsDirPath, postFileName)
		postFileData := posts.Post{
			Title:       product.Title,
			Description: product.BodyHTML,
			Date:        product.CreatedAt,
			ImageSrc:    product.Image.Src,
		}

		if len(product.Variants) == 0 {
			panic(fmt.Sprintf("product \"%s\" has no variants", product.Title))
		}

		if product.Variants[0].InventoryQuantity != 0 {
			postFileData.ShopifyId = product.Variants[0].ID
		}

		renderedPostFileData, err := postFileData.Render()
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(postFilePath, renderedPostFileData, 0644); err != nil {
			panic(err)
		}
	}
}
