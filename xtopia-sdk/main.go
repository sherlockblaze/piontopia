package main

type TestA struct {
	Key1 string
	Key2 struct {
		Sub_key1 string
		Sub_key2 []string
	}
}

func main() {
	s := TestA{}
	conf.Parse("test.yaml", s)
}
