package main

import (
	"fmt"
	"os"

	"github.com/hofstadter-io/cinful"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "list" {
		cinful.PrintVendors()
		return
	}

	vendor := cinful.Info()
	if vendor != nil {
		fmt.Println(vendor)
	}
}

