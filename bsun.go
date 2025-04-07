package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func main() {
	rawJsonData, err := os.ReadFile(path.Join("tests", "sample.json"))

	if err != nil {
		panic(err)
	}

	writeToBson(rawJsonData)
}

func readBson(filename string) {

	dat, err := os.ReadFile(filename)
	var raw bson.Raw = dat

	if err != nil {
		panic(err)
	}

	err = raw.Validate()

	if err != nil {
		panic(err)
	}

}

/*
* This function takes in raw JSON data in the form of
* a byte array, and returns the data in BSON format
 */
func writeToBson(jsonData []byte) {

	// using this as a general type for JSON objects.
	var jsonObj map[string]interface{}

	if err := json.Unmarshal(jsonData, &jsonObj); err != nil {
		panic(err)
	}
	// TODO: figure out how to convert this jsonObj to BSON format

	rawBson, err := bson.Marshal(jsonObj)

	if err != nil {
		panic(err)
	}

	os.WriteFile(path.Join("tests", "sample.bson"), rawBson, os.ModeAppend)
}

/*
* random test code that I had in main. Will most likely remove soon.
 */

func randomJunk() {
	jsonToBson := bson.M{
		"name":    "ui",
		"private": true,
		"version": "0.0.0",
		"type":    "module",
		"scripts": bson.M{
			"dev":     "vite",
			"build":   "tsc -b && vite build",
			"lint":    "eslint .",
			"preview": "vite preview",
		},
		"dependencies": bson.M{
			"octokit":      "^4.1.0",
			"react":        "^18.3.1",
			"react-dom":    "^18.3.1",
			"react-router": "^7.1.3",
			"swr":          "^2.3.2",
		},
		"devDependencies": bson.M{
			"@eslint/js":                  "^9.17.0",
			"@types/react":                "^18.3.18",
			"@types/react-dom":            "^18.3.5",
			"@vitejs/plugin-react":        "^4.3.4",
			"eslint":                      "^9.17.0",
			"eslint-plugin-react-hooks":   "^5.0.0",
			"eslint-plugin-react-refresh": "^0.4.16",
			"globals":                     "^15.14.0",
			"typescript":                  "~5.6.2",
			"typescript-eslint":           "^8.18.2",
			"vite":                        "^6.0.5",
			"wrangler":                    "^3.103.2",
		},
	}

	dat, err := os.ReadFile(path.Join("tests", "sample.json"))

	if err != nil {
		panic(err)
	}

	name := flag.String("name", "world", "The name to greet.")
	flag.Parse()
	fmt.Printf("Hello, %s!\n", *name)

	fmt.Printf("JSON FILE: %s", dat)
	rawBson := jsonToBson.String()

	fmt.Printf("JSON: %s", rawBson)
}
