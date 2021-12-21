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

	versions := []string{"https://github.com/tmukherjee13/go-concourse", "https://github.com/tmukherjee13/yii2-reverse-migration"}

	file, _ := os.OpenFile("my-resource/repos.json", os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(versions)

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
