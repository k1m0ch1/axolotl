package utils

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGenerateConfig(t *testing.T) {
	var uc UserConfig
	err := uc.GenerateConfig("test")
	if err != nil {
		t.Fatal(err)
	}
	content, err := ioutil.ReadFile("./config.yml")

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(content), "project-owner: test")
}
