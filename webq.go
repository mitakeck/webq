package main

import (
  "flag"
  "fmt"
  "os"
  "strings"
  "io/ioutil"
  "github.com/PuerkitoBio/goquery"
)

func check(e error)  {
  if e != nil {
    panic(e)
  }
}

func main() {
  flag.Parse()

  body, err := ioutil.ReadAll(os.Stdin)
  check(err)
  stringReader := strings.NewReader(string(body))

  switch flag.NArg() {
  case 1:
    doc, err := goquery.NewDocumentFromReader(stringReader)
    check(err)
    doc.Find(flag.Arg(0)).Each(func(_ int, s *goquery.Selection) {
      fmt.Println(s.Html())
    })
  default:
    fmt.Println(string(body))
  }

}
