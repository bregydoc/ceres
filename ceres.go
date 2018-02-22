package main

// ConfigFile ...
type ConfigFile struct {
	Project struct {
		Name    string `json:"name"`
		Author  string `json:"author"`
		Version string `json:"version"`
	} `json:"project"`

	Types struct {
		Definitions string `json:"definitions"`
	} `json:"types"`

	Db struct {
		Path string `json:"path"`
	} `json:"db"`
}
