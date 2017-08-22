package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

var imageURLs = []string{
	"https://img.wenhairu.com/image/As5IP",
	"https://img.wenhairu.com/image/Asp1h",
	"https://img.wenhairu.com/image/As1mu",
	"https://img.wenhairu.com/image/As9BI",
	"https://img.wenhairu.com/image/AsuR6",
	"https://img.wenhairu.com/image/Asawp",
	"https://img.wenhairu.com/image/AsE3X",
	"https://img.wenhairu.com/image/AsDgq",
	"https://img.wenhairu.com/image/AsWEH",
	"https://img.wenhairu.com/image/AsJhd",
	"https://img.wenhairu.com/image/AsCIf",
	"https://img.wenhairu.com/image/Asw1o",
	"https://img.wenhairu.com/image/AsfA3",
	"https://img.wenhairu.com/image/AsNBK",
	"https://img.wenhairu.com/image/AsFQg",
	"https://img.wenhairu.com/image/Asqfj",
	"https://img.wenhairu.com/image/AsA3U",
	"https://img.wenhairu.com/image/As8c0",
	"https://img.wenhairu.com/image/As6Dv",
	"https://img.wenhairu.com/image/AivhG",
	"https://img.wenhairu.com/image/Ail4T",
	"https://img.wenhairu.com/image/Aih1A",
	"https://img.wenhairu.com/image/AiLAn",
	"https://img.wenhairu.com/image/AizeB",
	"https://img.wenhairu.com/image/AitQN",
	"https://img.wenhairu.com/image/AirfR",
	"https://img.wenhairu.com/image/Ai22s",
	"https://img.wenhairu.com/image/AiscC",
	"https://img.wenhairu.com/image/AiiDS",
	"https://img.wenhairu.com/image/AibLt",
	"https://img.wenhairu.com/image/AiB4D",
	"https://img.wenhairu.com/image/AieGP",
	"https://img.wenhairu.com/image/AiX8h",
	"https://img.wenhairu.com/image/AiTeu",
	"https://img.wenhairu.com/image/AiOHI",
	"https://img.wenhairu.com/image/AiIf6",
	"https://img.wenhairu.com/image/Ai02p",
	"https://img.wenhairu.com/image/AiZSX",
	"https://img.wenhairu.com/image/AiPDq",
	"https://img.wenhairu.com/image/AigLH",
	"https://img.wenhairu.com/image/Aic0d",
	"https://img.wenhairu.com/image/AiSGf",
	"https://img.wenhairu.com/image/Ai78o",
	"https://img.wenhairu.com/image/AikX3",
	"https://img.wenhairu.com/image/AiRHK",
	"https://img.wenhairu.com/image/AiQYg",
	"https://img.wenhairu.com/image/Aiysj",
	"https://img.wenhairu.com/image/AinSU",
	"https://img.wenhairu.com/image/Ai5W0",
	"https://img.wenhairu.com/image/Ai1jv",
	"https://img.wenhairu.com/image/AiG0G",
	"https://img.wenhairu.com/image/Ai99T",
	"https://img.wenhairu.com/image/Aiu8A",
	"https://img.wenhairu.com/image/AiUXn",
	"https://img.wenhairu.com/image/AiEyB",
	"https://img.wenhairu.com/image/AiDYN",
	"https://img.wenhairu.com/image/AiKsR",
	"https://img.wenhairu.com/image/AiJ7s",
	"https://img.wenhairu.com/image/AiCWC",
	"https://img.wenhairu.com/image/AifjS",
	"https://img.wenhairu.com/image/AiYZt",
	"https://img.wenhairu.com/image/AiN9D",
	"https://img.wenhairu.com/image/AiF6P",
	"https://img.wenhairu.com/image/AimVh",
	"https://img.wenhairu.com/image/AiAyu",
	"https://img.wenhairu.com/image/Ai8NI",
	"https://img.wenhairu.com/image/Adxs6",
	"https://img.wenhairu.com/image/Adv7p",
	"https://img.wenhairu.com/image/AdlKX",
	"https://img.wenhairu.com/image/AdLjq",
	"https://img.wenhairu.com/image/AdjZH",
	"https://img.wenhairu.com/image/Adzud",
	"https://img.wenhairu.com/image/Adt6f",
	"https://img.wenhairu.com/image/Ad3Vo",
	"https://img.wenhairu.com/image/Ad2n3",
}

func extractImageInfo(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %d", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s failed: %v", url, err)
	}

	extract := func(node *html.Node) {
		if node != nil && node.Type == html.ElementNode && node.Data == "meta" {
			for _, attr := range node.Attr {
				if attr.Val == "og:title" {
					fmt.Printf("\n%s ===> ", node.Attr[1].Val)
				}
				if attr.Val == "og:image" {
					fmt.Printf("%s\n", node.Attr[1].Val)
				}
			}
		}
	}
	forEachNode(doc, extract, nil)

	return nil
}

// forEachNode 遍历给定HTML文档节点的每个子节点，对每个节点执行预处理操作和后处理操作
func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}

func main() {
	for _, url := range imageURLs {
		if err := extractImageInfo(url); err != nil {
			fmt.Printf("got err: %v\n", err)
		}
	}
}
