package models

type Manifest struct {
	Repos []Repo `yaml:"repos"`
}

type Repo struct {
	Url  string `yaml:"url"`
	Name string `yaml:"name"`
}

type Version struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type SourceIn struct {
	Source Source `json:"source"`
}

type Source struct {
	Uri string `json:"uri"`
}

type Out struct {
	Version Version `json:"version"`
}
