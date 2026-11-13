package main

import (
	"flag"
	"fmt"

	"github.com/rifkisururi/belajar-golang/internal/greeting"
)

func main() {
	name := flag.String("name", "Gopher", "name to greet")
	flag.Parse()
	fmt.Println(greeting.Message(*name))
}
