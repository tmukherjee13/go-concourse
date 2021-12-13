package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}
	fmt.Printf("%s", bytes)

	versions := []string{"HEAD"}
	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
