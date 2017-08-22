package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// ExtractLinks 解析给定URL里的链接，会过滤掉重复的
func ExtractLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s failed: %v", url, err)
	}

	links := []string{}
	visited := make(map[string]bool)

	// 从HTML节点解析链接
	extractLink := func(node *html.Node) {
		if node != nil && node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					link, err := resp.Request.URL.Parse(attr.Val)
					if err != nil {
						continue
					}
					if !visited[link.String()] {
						links = append(links, link.String())
						visited[link.String()] = true
					}
				}
			}
		}
	}
	ForEachNode(doc, extractLink, nil)

	return links, nil
}

// ForEachNode 遍历给定HTML文档节点的每个子节点，对每个节点执行预处理操作和后处理操作
func ForEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

func main() {
	url := os.Args[1]
	links, err := ExtractLinks(url)
	if err != nil {
		log.Fatal("extract links failed: ", err)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
