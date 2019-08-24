package yaml

import (
	"io/ioutil"
	"log" // 后期改成用库中的logger

	"gopkg.in/yaml.v2"
)

/**
Parse reading and parsing a yaml config file
@param path : a yaml file path
@param out : the address of yaml struct instance
@return void.
**/
func Parse(path string, out interface{}) {

	err := yaml.Unmarshal(read(path), &out)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("DEBUG | your yaml conf: %v", out)
}

func read(path string) []byte {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return yamlFile
}
