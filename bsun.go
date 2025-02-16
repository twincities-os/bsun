package main

import (
	"flag"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)
}

func readBson(filname string) {

	dat, err := os.ReadFile(filname)
	var raw bson.Raw = dat

	if err != nil {
		panic(err)
	}

	err = raw.Validate()

	if err != nil {
		panic(err)
	}

}

func writeToBson() {}
