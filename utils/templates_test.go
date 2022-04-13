package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/thanhpk/randstr"
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

func TestGenerateVuln(t *testing.T) {
	randomdir := randstr.Hex(6)
	var uc UserConfig
	uc.DirConfig.VulnDir = fmt.Sprintf("temporarydir/%s/vulns", randomdir)

	err := uc.GenerateVuln("example.com", "IDOR")
	if err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("./%s/example.com/IDOR.yml", uc.DirConfig.VulnDir))

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(content), "desc: Found the Finding")
	assert.Contains(t, string(content), "domain: example.com")
}

func TestGenerateHost(t *testing.T) {
	randomdir := randstr.Hex(6)
	var uc UserConfig
	uc.DirConfig.HostsIdentityDir = fmt.Sprintf("temporarydir/%s/hosts", randomdir)

	err := uc.GenerateHost("example.com")
	if err != nil {
		t.Fatal(err)
	}

	content, err := ioutil.ReadFile(fmt.Sprintf("./%s/example.com.yml", uc.DirConfig.HostsIdentityDir))

	if err != nil {
		t.Fatal(err)
	}

	assert.Contains(t, string(content), "id: example.com")
	assert.Contains(t, string(content), "url: example.com")
}
