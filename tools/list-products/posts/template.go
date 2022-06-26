package posts

import (
	"text/template"
	"time"
)

var Template *template.Template

type PostData struct {
	Title       string
	Description string
	Date        time.Time
	ImageSrc    string
}

const templateContent = `---
layout: post
title: {{ .Title }}
description: {{ .Description }}
date: {{ .Date.Format "2006-01-02 15:04:05" }}
hiQualPath: {{ .ImageSrc }}
loQualPath: {{ .ImageSrc }}
---`

func init() {
	err := error(nil)
	Template, err = template.New("post").Parse(templateContent)
	if err != nil {
		panic(err)
	}
}
