package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Response struct {
	Url    string
	Commit string
}

func main() {

	_, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	versions := []Response{
		{
			Url:    "https://github.com/tmukherjee13/yii2-reverse-migration",
			Commit: "abc",
		},
	}
	file, _ := os.OpenFile("repos.json", os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(versions)

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
