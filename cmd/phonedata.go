package main

import (
	"fmt"
	"github.com/lxshilaoda98/qcodecall"
)

func main() {

	pr, err := phonedata.Find("01065080114", "", nil)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Print(pr)
}
