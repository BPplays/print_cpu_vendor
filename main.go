package main

import (
    "fmt"
	"flag"
	. "github.com/klauspost/cpuid/v2"
)

func main() {
	vendorid := flag.Bool("v", false, "vendorid")
	family := flag.Bool("b", false, "family")
	model := flag.Bool("m", false, "model")

	if *family {
		fmt.Println(CPU.Family)
	}

	if *model {
		fmt.Println(CPU.Model)
	}

	if *vendorid {
		fmt.Println(CPU.VendorID)
	}
}
