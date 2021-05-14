package main

import (
	"fmt"

	quotev2 "rsc.io/quote/v2"
	"rsc.io/quote/v3"
	// "github.com/cmwylie19/go-practice/hello"
)

func main() {
	fmt.Printf("%q\n", quote.GlassV3())
	fmt.Printf("%q", quotev2.OptV2())
}

// func main() {
// 	fmt.Println(hello.Hello())
// 	fmt.Println(hello.Proverb())
// }
