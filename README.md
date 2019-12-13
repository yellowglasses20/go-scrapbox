# go-scrapbox [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![GoDoc](https://godoc.org/github.com/yellowglasses20/go-scrapbox?status.svg)](https://godoc.org/github.com/yellowglasses20/go-scrapbox) [![Go Report Card](https://goreportcard.com/badge/yellowglasses20/go-scrapbox)](https://goreportcard.com/report/github.com/yellowglasses20/go-scrapbox) ![](https://github.com/yellowglasses20/go-scrapbox/workflows/Go/badge.svg)

## usage
```go
import (
  "github.com/yellowglasses20/go-scrapbox"
  "fmt"
)

func main() {
  // Create scprapbox client
  // Please set the project name you want to get in "help-jp"
  client, _ := qiita.NewClient("help-jp")

  // Fetch pagelist and print them
  items, _, _ := client.PageList(&ListOptions{Skip: 0, Limit: 1})
  fmt.Println(items)

  // Fetch page title and print them
  items, _, _ := client.PageTitle("API")
  fmt.Println(items)

  // Fetch page text and print them
  body, _, _ := client.PageText("API")
  fmt.Printf("%s", body)

  // Fetch page icon
  bin, _, _ := client.PageIcon("API")
}
```
