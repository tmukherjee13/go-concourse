package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Response struct {
	Url    string
	Commit string
}

type InResponse struct {
	Version Response `json:"version"`
}

func main() {
	_, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("Error reading from stdin: ", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("incomplete list of argumetns")
	}

	destDir := os.Args[1]
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		log.Fatal("unable to create directory")
	}

	outData := []string{
		"https://github.com/tmukherjee13/yii2-reverse-migration",
	}
	versions := &InResponse{
		Version: Response{
			Url:    "https://github.com/tmukherjee13/yii2-reverse-migration",
			Commit: "abc",
		},
	}

	data, err := json.Marshal(outData)

	// file, _ := os.OpenFile("my-resource/repos.json", os.O_CREATE, os.ModePerm)
	// defer file.Close()
	// encoder := json.NewEncoder(file)
	// encoder.Encode(versions)

	destFile := filepath.Join(destDir, "repos.json")
	ioutil.WriteFile(destFile, data, 0755)

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encoding error: %s", err)
	}
}
