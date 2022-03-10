package utils

import (
	"os"
	"fmt"

	yaml "github.com/goccy/go-yaml"
)

func (cfg *UserConfig) GenerateHost(host string) error{
	var h HostIdentity

	defaultOthers := []Other{
		Other{
			Source: "",
			Desc: "",
			URL: "",
		},
	}

	defaultRecons := []Recon{
		Recon{
			Tools: "",
			Type: "",
			Report: "",
			FileType: "",
			Desc: "",
			Images: []Image{
				Image{
					Path: "",
					Caption: "",
				},
			},
		},
	}

	defaultVulnScans := []VulnScan{
		VulnScan{
			Tools: "",
			Type: "",
			Report: "",
			FileType: "",
			Desc: "",
			Images: []Image{
				Image{
					Path: "",
					Caption: "",
				},
			},
		},
	}

	h.ID = host
	h.Info.Others = defaultOthers
	h.Recons = defaultRecons
	h.VulnScans = defaultVulnScans

    bytes, err := yaml.Marshal(h)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, host), bytes, 0644)
	if err != nil {
		return err
    }

	return nil
}

func (cfg *UserConfig) GenerateVuln(host string, nameVuln string) error {
	var f Finding
    bytes, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("./%s/%s/%s.yml", cfg.DirConfig.VulnDir, host, nameVuln), bytes, 0644)
	if err != nil {
		return err
    }
	return nil
}

func (cfg *UserConfig) GenerateUserConfig() []byte {
	var uc UserConfig
    bytes, err := yaml.Marshal(uc)
	if err != nil {
		log.Fatalf("Marshal: %v", err)
	}
	err = os.WriteFile("./config.yml", bytes, 0644)
	if err != nil {
		log.Printf("err   #%v ", err)
    }
    return bytes
}