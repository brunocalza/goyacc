package main

import (
	"fmt"

	"github.com/brunocalza/goyacc/examples/phone"
)

func main() {
	fmt.Println(phone.Parse("800.555.1234"))
}
