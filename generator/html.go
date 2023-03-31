package main

import (
	"bytes"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func mdToHTML(md []byte) []byte {
	p := parser.NewWithExtensions(parser.CommonExtensions)
	doc := p.Parse(md)
	//htmlFlags := html.CommonFlags | html.HrefTargetBlank
	htmlFlags := html.FlagsNone

	opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: func(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
		if val, ok := node.(*ast.Link); ok {
			if bytes.Contains(val.Destination, []byte(".md")) {
				val.Destination = bytes.Replace(val.Destination, []byte(".md"), []byte(".html"), -1)
			}
		}
		return 0, false
	}}
	renderer := html.NewRenderer(opts)
	return markdown.Render(doc, renderer)
}

func RenderHTML(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			content, err := os.ReadFile(path)
			if err != nil {
				logrus.Fatal(err)
			}
			newName := strings.TrimSuffix(path, ".md")
			newName = newName + ".html"
			var temp = []byte(HEAD)

			if strings.Contains(newName, "SUMMARY") {
				ffff, _ := filepath.Split(newName)
				newName = filepath.Join(ffff, "index.html")
				temp = []byte(INDEX_HEAD)
			}
			logrus.Println(newName)

			out := mdToHTML(content)
			temp = append(temp, out...)
			temp = append(temp, END...)

			err = os.WriteFile(newName, temp, os.ModePerm)
			if err != nil {
				logrus.Fatal(err)

			}
		}
		return nil
	})

	if err != nil {
		logrus.Fatal(err)
	}
}

var INDEX_HEAD = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>hello</title>
<style>
        p {
            margin-block-start: 0;
            margin-block-end: 0;
        }
body {
    display: block;
    margin: 40px;
 font-family: "Consolas","Panic Sans","Bitstream Vera Sans Mono","Menlo","Microsoft Yahei",monospace;

}
    </style>
<body>
    `

var HEAD = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>hello</title>
<link rel="stylesheet"
      href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/styles/default.min.css">
<script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.7.0/highlight.min.js"></script>
<script>hljs.highlightAll();</script>
<style>
        p {
            margin-block-start: 0;
            margin-block-end: 0;
        }
body {
    display: block;
    margin: 40px;
 font-family: "Consolas","Panic Sans","Bitstream Vera Sans Mono","Menlo","Microsoft Yahei",monospace;

}
    </style>
<body>
    `

var END = `</body></html>`
