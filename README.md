# go-scrapbox
Go library for accessing the Scrapbox API https://scrapbox.io/help-jp/API

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