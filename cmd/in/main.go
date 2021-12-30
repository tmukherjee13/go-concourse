package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tmukherjee13/go-concourse/models"
	"gopkg.in/yaml.v2"
)

func main() {
	in, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("incomplete list of argumetns")
	}

	var input models.SourceIn
	json.Unmarshal(in, &input)

	destDir := os.Args[1]
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		log.Fatal("unable to create directory")
	}

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
	var version models.Version
	for _, repo := range manifest.Repos {
		version = models.Version(repo)
		versions = append(versions, version)
	}

	// file, _ := os.OpenFile("my-resource/repos.json", os.O_CREATE, os.ModePerm)
	// defer file.Close()
	// encoder := json.NewEncoder(file)
	// encoder.Encode(versions)

	data, _ := json.Marshal(versions)
	destFile := filepath.Join(destDir, "repos.json")
	ioutil.WriteFile(destFile, data, 0755)

	err = json.NewEncoder(os.Stdout).Encode(models.Out{Version: version})
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
