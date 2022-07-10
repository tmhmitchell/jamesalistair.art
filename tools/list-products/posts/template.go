package posts

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

var postTemplate *template.Template

const templateContent = `---
layout: post
title: {{ .Title }}
description: {{ .Description }}
date: {{ .Date.Format "2006-01-02 15:04:05" }}
hiQualPath: {{ .ImageSrc }}
loQualPath: {{ .ImageSrc }}
productId: {{ .ShopifyId }}
---`

func init() {
	err := error(nil)
	postTemplate, err = template.New("post").Parse(templateContent)
	if err != nil {
		panic(err)
	}
}

type Post struct {
	Title       string
	Description string
	Date        time.Time
	ImageSrc    string
	ShopifyId   int64
}

func (p Post) Render() ([]byte, error) {
	buf := bytes.Buffer{}

	// XXX Need to be fixed, tags in the description are messing up the
	// posts/YAML
	var builder strings.Builder
	for _, line := range strings.Split(p.Description, "\n") {
		builder.WriteString(line)
	}
	p.Description = builder.String()

	err := postTemplate.Execute(&buf, p)
	return buf.Bytes(), err
}
