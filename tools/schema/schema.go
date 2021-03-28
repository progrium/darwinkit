package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

type Class struct {
	Name            string
	Description     string
	InstanceMethods []InstanceMethod
	TypeMethods     []TypeMethod
	Properties      []Property
	Frameworks      []string
	Platforms       []string
	URL             string
}

type Property struct {
	Name string
	URL  string
}

type InstanceMethod struct {
	Name string
	URL  string
}

type TypeMethod struct {
	Name string
	URL  string
}

type Topic struct {
	Name string
	URL  string
}

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var class Class
	class.URL = os.Args[1]

	if err := chromedp.Run(ctx,
		chromedp.Navigate(class.URL),
		chromedp.WaitVisible(`main`),
		chromedp.Text(`main div.topictitle h1.title`, &class.Name),
		chromedp.Text(`main div.container div.description div.abstract.content`, &class.Description),
		textList(`main div.summary div.frameworks ul li span`, &class.Frameworks),
		textList(`main div.summary div.availability ul li span`, &class.Platforms),
	); err != nil {
		log.Fatal(err)
	}

	topics := "#topics div.contenttable-section div.section-content div.topic a.link:not(.deprecated)"
	for _, n := range nodes(ctx, topics) {
		var t Topic
		var ok bool
		if err := chromedp.Run(ctx, chromedp.TextContent(n.FullXPathByID(), &t.Name)); err != nil {
			log.Fatal(err)
		}
		if err := chromedp.Run(ctx, chromedp.AttributeValue(n.FullXPathByID(), "href", &t.URL, &ok)); err != nil {
			log.Fatal(err)
		}

		// Handle related consts/enums later
		if strings.HasPrefix(t.Name, "NS") ||
			strings.HasPrefix(t.Name, "CG") ||
			strings.HasPrefix(t.Name, "UI") ||
			strings.HasPrefix(t.Name, "WK") {
			continue
		}

		// Manual fix for less than perfect selector
		if strings.HasPrefix(t.Name, "API Reference") {
			continue
		}

		if t.URL != "" {
			t.URL = fmt.Sprintf("https://developer.apple.com%s", t.URL)
		}

		if strings.HasPrefix(t.Name, "+ ") {
			class.TypeMethods = append(class.TypeMethods, TypeMethod{Name: t.Name[2:], URL: t.URL})
		} else if strings.HasPrefix(t.Name, "- ") {
			class.InstanceMethods = append(class.InstanceMethods, InstanceMethod{Name: t.Name[2:], URL: t.URL})
		} else {
			class.Properties = append(class.Properties, Property{Name: t.Name, URL: t.URL})
		}
	}

	b, err := json.MarshalIndent(class, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func nodes(ctx context.Context, sel string) []*cdp.Node {
	var nodes []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(sel, &nodes)); err != nil {
		log.Fatal(err)
	}
	return nodes
}

func textList(sel string, lst *[]string) chromedp.Tasks {
	var nodes []*cdp.Node
	return chromedp.Tasks{
		chromedp.Nodes(sel, &nodes),
		chromedp.ActionFunc(func(ctx context.Context) error {
			for _, n := range nodes {

				txt := innerText(n)
				if txt != "" {
					*lst = append(*lst, txt)
				}
			}
			return nil
		}),
	}
}

func innerText(node *cdp.Node) string {
	var t []string
	for _, c := range node.Children {
		switch c.NodeType {
		case cdp.NodeTypeText:
			t = append(t, strings.Trim(c.NodeValue, " \n"))
		case cdp.NodeTypeElement:
			t = append(t, innerText(c))
		}
	}
	return strings.Trim(strings.Join(t, ""), " \n")
}
