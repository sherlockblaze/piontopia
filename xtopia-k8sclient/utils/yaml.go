package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

/*
LoadYAML reading and parsing a yaml config file
@param path: a yaml file path
@param out: the address of yaml struct instance
@return void
**/
func LoadYAML(path string, out interface{}) error {
	err := yaml.Unmarshal(read(path), out)
	if err != nil {
		log.Fatalf("error: %v", err)
		return err
	}

	err = validator.Validate(out)
	if err != nil {
		return err
	}

	log.Printf("DEBUG | your yaml conf: %v", out)

	return nil
}

func read(path string) []byte {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return yamlFile
}
