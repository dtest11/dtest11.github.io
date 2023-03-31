package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"os"
	"path/filepath"
	"strings"
)

func renderLinkNode(f string) {
	var str strings.Builder
	file, err := os.Open(f)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "{{#include") {
			str.WriteString(scanner.Text())
			str.WriteString("\n")
			continue
		}
		linkNode := parseInclude(scanner.Text(), filepath.Dir(f))
		content := linkNode.GetContent()
		str.WriteString(content)
		str.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)
	}

	// 写入到outputDir 下面
	dir := filepath.Join(ScanDir, OutputDir, f)
	makeDir(dir)
	content := str.String()
	err = os.WriteFile(dir, []byte(content), os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
}

func makeDir(path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		logrus.Fatal(err)
	}
}

type LinkNode struct {
	parentDir string
	Name      string `json:"name"`
	Start     int    `json:"start"`
	End       int    `json:"end"`
}

func (l *LinkNode) GetContent() string {
	path := filepath.Join(l.parentDir, l.Name)
	file, err := os.Open(path)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()
	var line int
	var content strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if l.End == 0 && l.Start == 0 {
			content.WriteString(scanner.Text())
			content.WriteString("\n")
			continue
		}
		if line >= l.Start && line < l.End {
			content.WriteString(scanner.Text())
			content.WriteString("\n")

		}
		line++
	}

	if err := scanner.Err(); err != nil {
		logrus.Fatal(err)
	}
	return content.String()
}

func parseInclude(s string, dir string) LinkNode {
	s = strings.ReplaceAll(s, "{{", "")
	s = strings.ReplaceAll(s, "}}", "")
	s = strings.ReplaceAll(s, "#include", "")
	var res LinkNode = LinkNode{parentDir: dir}
	if strings.Contains(s, ":") {
		index := strings.Split(s, ":")
		res = LinkNode{
			parentDir: dir,
			Name:      index[0],
			Start:     cast.ToInt(index[1]),
			End:       cast.ToInt(index[2]),
		}
	} else {
		res.Name = s
		res.parentDir = dir
	}
	res.Name = strings.TrimSpace(res.Name)
	return res
}

func RenderMd(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			renderLinkNode(path)
			logrus.Infof("render md: %s", path)

		}
		return nil
	})

	if err != nil {
		logrus.Fatal(err)
	}
}
