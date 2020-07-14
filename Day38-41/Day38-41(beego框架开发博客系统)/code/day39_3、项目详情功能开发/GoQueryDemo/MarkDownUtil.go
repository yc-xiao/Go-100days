package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
)

func main() {

	fileread, _ := ioutil.ReadFile("./test.md")
	//转换Markdown语法，如将"#"转换为"<h1></h1>"
	subHtml := blackfriday.MarkdownCommon(fileread)
	subHtmlStr := string(subHtml)
	fmt.Println(subHtmlStr)
}
