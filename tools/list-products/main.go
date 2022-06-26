package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Variants struct {
	ProductID            int64       `json:"product_id"`
	ID                   int64       `json:"id"`
	Title                string      `json:"title"`
	Price                string      `json:"price"`
	Sku                  string      `json:"sku"`
	Position             int         `json:"position"`
	InventoryPolicy      string      `json:"inventory_policy"`
	CompareAtPrice       interface{} `json:"compare_at_price"`
	FulfillmentService   string      `json:"fulfillment_service"`
	InventoryManagement  interface{} `json:"inventory_management"`
	Option1              string      `json:"option1"`
	Option2              interface{} `json:"option2"`
	Option3              interface{} `json:"option3"`
	CreatedAt            time.Time   `json:"created_at"`
	UpdatedAt            time.Time   `json:"updated_at"`
	Taxable              bool        `json:"taxable"`
	Barcode              string      `json:"barcode"`
	Grams                int         `json:"grams"`
	ImageID              interface{} `json:"image_id"`
	Weight               float64     `json:"weight"`
	WeightUnit           string      `json:"weight_unit"`
	InventoryItemID      int64       `json:"inventory_item_id"`
	InventoryQuantity    int         `json:"inventory_quantity"`
	OldInventoryQuantity int         `json:"old_inventory_quantity"`
	RequiresShipping     bool        `json:"requires_shipping"`
	AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
}

type Options struct {
	ProductID int64    `json:"product_id"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Position  int      `json:"position"`
	Values    []string `json:"values"`
}

type Images struct {
	ProductID         int64         `json:"product_id"`
	ID                int64         `json:"id"`
	Position          int           `json:"position"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Alt               interface{}   `json:"alt"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	Src               string        `json:"src"`
	VariantIds        []interface{} `json:"variant_ids"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
}

type Image struct {
	ProductID         int64         `json:"product_id"`
	ID                int64         `json:"id"`
	Position          int           `json:"position"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	Alt               interface{}   `json:"alt"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	Src               string        `json:"src"`
	VariantIds        []interface{} `json:"variant_ids"`
	AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
}

type Products struct {
	ID                int64      `json:"id"`
	Title             string     `json:"title"`
	BodyHTML          string     `json:"body_html"`
	Vendor            string     `json:"vendor"`
	ProductType       string     `json:"product_type"`
	CreatedAt         time.Time  `json:"created_at"`
	Handle            string     `json:"handle"`
	UpdatedAt         time.Time  `json:"updated_at"`
	PublishedAt       time.Time  `json:"published_at"`
	TemplateSuffix    string     `json:"template_suffix"`
	Status            string     `json:"status"`
	PublishedScope    string     `json:"published_scope"`
	Tags              string     `json:"tags"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id"`
	Variants          []Variants `json:"variants"`
	Options           []Options  `json:"options"`
	Images            []Images   `json:"images"`
	Image             Image      `json:"image"`
}

type Payload struct {
	Products []Products `json:"products"`
}

func kebabCase(s string) string {
	lowerWords := make([]string, 0)
	for _, word := range strings.Split(s, " ") {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}

	return strings.Join(lowerWords, "-")
}

func main() {
	url := "https://james-alistair-art.myshopify.com/admin/api/2021-10/products.json"

	token := os.Getenv("SHOPIFY_API_TOKEN")
	if token == "" {
		panic("shopify api token is empty")
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Shopify-Access-Token", token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
	fmt.Println()

	payload := Payload{}
	err = json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		panic(err)
	}

	for _, product := range payload.Products {
		fmt.Println(product)
		fmt.Println(product.Variants[0].InventoryQuantity)

		fmt.Printf(
			"%s-%s.md\n",
			product.CreatedAt.Format("2006-01-02"),
			kebabCase(product.Title),
		)
	}
}
