package utils

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGenerateConfig(t *testing.T) {
	var uc UserConfig
	uc.GenerateConfig("test")
	content, err := ioutil.ReadFile("./config.yml")

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(content), "project-owner: test")
}
