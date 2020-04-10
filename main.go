package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/aman-bansal/go_perf_tracker/perf_tracker"
	"io/ioutil"
)

func main() {
	cfgFile := flag.String("config", "", "path of config file to load")
	flag.Parse()
	if len(*cfgFile) == 0 {
		fmt.Println("Config file path is not provided")
		panic(errors.New("config file path is not provided"))
	}
		cfg, err := loadFileAndFuncMap(*cfgFile)
	if err != nil {
		fmt.Println("Unable to load config, got error:", err)
		panic(err)
	}
	err = perf_tracker.GenerateCode(cfg)
	if err != nil {
		fmt.Println("Unable to generate files, got error:", err)
		panic(err)
	}
}

//config file should be a map of filepath vs list of func names
func loadFileAndFuncMap(filePath string) (map[string][]string, error) {
	byteData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var funcMap map[string][]string
	// Here, we unmarshal our byteData into 'funcMap' struct.
	err = json.Unmarshal(byteData, &funcMap)
	if err != nil {
		return nil, err
	}

	return funcMap, nil
}