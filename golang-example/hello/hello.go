package main // import "github.com/you/hello"

import (
	"fmt"

	"rsc.io/quote"
	"github.com/hashicorp/go-getter"
)

func main() {
	fmt.Println(quote.Hello())
}
