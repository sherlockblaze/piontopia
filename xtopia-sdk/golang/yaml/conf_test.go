package yaml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestA struct {
	Key1 string `yaml:"key1"`
	Key2 Key2   `yaml:"key2"`
}

type Key2 struct {
	SubKey1 string   `yaml:"subkey1"`
	SubKey2 []string `yaml:"subkey2"`
}

func TestParse(t *testing.T) {
	s1 := &TestA{}
	err := LoadYAML("./test/test.yaml", &s1)
	assert.NoError(t, err)

	s2 := &TestA{
		Key1: "value1",
		Key2: Key2{
			SubKey1: "sub_value1",
			SubKey2: []string{"sub_value2", "sub_value3"},
		},
	}
	assert.EqualValues(t, s2, s1)
}
