package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tmukherjee13/go-concourse/models"
	"gopkg.in/yaml.v2"
)

func main() {

	in, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	var input models.SourceIn
	json.Unmarshal(in, &input)

	// Read manifest file
	resp, err := http.Get(fmt.Sprintf("%s/raw/master/manifest.yml", input.Source.Uri))
	if err != nil {
		log.Fatal("Error reading manifest file: ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading manifest file: ", err)
	}

	var manifest models.Manifest
	yaml.Unmarshal(body, &manifest)

	var versions []models.Version
	for _, repo := range manifest.Repos {
		versions = append(versions, models.Version{Url: repo.Url, Name: repo.Name})
	}

	file, _ := os.OpenFile("repos.json", os.O_CREATE, os.ModePerm)
	defer file.Close()
	encoder := json.NewEncoder(file)
	data, _ := json.Marshal(versions)
	encoder.Encode(data)

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
