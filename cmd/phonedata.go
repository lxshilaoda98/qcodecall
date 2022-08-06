package main

import (
	"fmt"
	"github.com/lxshilaoda98/phonedata"
)

func main() {

	pr, err := phonedata.Find("17600082595", "", nil)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Print(pr)
}
