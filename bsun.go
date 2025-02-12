package main

import (
	"flag"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}

func readBson() {
	var raw bson.Raw = nil
}
