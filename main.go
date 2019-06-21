package main

import (
    "fmt"
    "rsc.io/quote"
    "github.com/acmer1183/Hello/model"
)

func main() {
    fmt.Println(quote.Hello())
    model.Get()
}
