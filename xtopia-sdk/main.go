package main

import "piontopia/xtopia-sdk/golang/yaml"

type TestA struct {
	Key1 string
	Key2 struct {
		Sub_key1 string
		Sub_key2 []string
	}
}

func main() {
	s := TestA{}
	yaml.LoadYAML("test.yaml", s)
}
