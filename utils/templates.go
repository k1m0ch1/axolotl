package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	yaml "github.com/goccy/go-yaml"
)

func (cfg *UserConfig) GenerateConfig(name string) error {
	var uc UserConfig

	fmt.Println("[?] Generate file Config ./config.yml")

	defaultUserConfig := UserConfig{
		ProjectOwner: name,
		Group:        "",
		Teams: []Team{
			Team{
				TeamName: "",
				Members: []Member{
					Member{
						MemberName: "",
						Role:       "",
						Email:      "",
					},
				},
			},
		},
		Email: "",
		DirConfig: DirConfig{
			HostsIdentityDir:   "hosts",
			VulnDir:            "vulns",
			ToolsReports:       "reports",
			PocDir:             "poc",
			OutputReportsDir:   "outputs",
			TemplatesReportDir: "templates",
		},
	}
	uc = defaultUserConfig

	bytes, err := yaml.Marshal(uc)
	if err != nil {
		return err
	}

	err = os.WriteFile("./config.yml", bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (cfg *UserConfig) GenerateHost(host string) error {
	var h HostIdentity

	defaultHostInfo := HostInfo{
		URL:          host,
		TechStacks:   "",
		Tag:          "",
		OpenPorts:    "",
		FilteredPort: "",
		HostIP:       "",
		Country:      "",
		City:         "",
		Organization: "",
		Others: []Other{
			Other{
				Source: "",
				Desc:   "",
				URL:    "",
			},
		},
		Desc: "",
	}

	defaultRecons := []Recon{
		Recon{
			Tools:    "",
			Type:     "",
			Report:   "",
			FileType: "",
			Desc:     "",
			Images: []Image{
				Image{
					Path:    "",
					Caption: "",
				},
			},
		},
	}

	defaultVulnScans := []VulnScan{
		VulnScan{
			Tools:    "",
			Type:     "",
			Report:   "",
			FileType: "",
			Desc:     "",
			Images: []Image{
				Image{
					Path:    "",
					Caption: "",
				},
			},
		},
	}

	h.ID = host
	h.Info = defaultHostInfo
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

	defaultVulnInfo := VulnInfo{
		FindingName:     "",
		URL:             "",
		Domain:          host,
		Author:          "",
		Team:            "",
		Email:           "",
		Tags:            "",
		ReportTemplate:  "",
		EksternalReport: "",
		RiskRatings: RiskRating{
			Severity: "",
			CVSS:     "",
			OWASP:    "",
		},
	}

	defaultPoC := []ProofOfConcept{
		ProofOfConcept{
			Path: "",
			StepsToReproduce: []Step{
				Step{
					Desc: "",
					Images: []Image{
						Image{
							Path:    "",
							Caption: "",
						},
					},
				},
			},
			Images: []Image{
				Image{
					Path:    "",
					Caption: "",
				},
			},
			NucleiTemplate: "",
			Exploit:        "",
			Desc:           "",
		},
	}

	defaultHTF := []HowToFix{
		HowToFix{
			Information: "",
			Desc:        "",
			URL:         "",
			StepsToFix: []Step{
				Step{
					Desc: "",
					Images: []Image{
						Image{
							Path:    "",
							Caption: "",
						},
					},
				},
			},
		},
	}

	datetime := time.Now()

	f.ID = nameVuln
	f.VulnInfo = defaultVulnInfo
	f.ProofOfConcept = defaultPoC
	f.HowToFix = defaultHTF
	f.Bounty = Bounty{
		BugLink:       "",
		WriteUp:       "",
		ApprovedScore: "",
		Reward:        "",
		Desc:          "",
	}
	f.Status = Status{
		Created: StatusFields{
			By:   cfg.ProjectOwner,
			Time: datetime.Format(time.RFC3339),
			Desc: "Found the Finding",
		},
	}

	bytes, err := yaml.Marshal(f)
	if err != nil {
		return err
	}
	pathDir := fmt.Sprintf("./%s/%s", cfg.DirConfig.VulnDir, host)
	if _, err := os.Stat(pathDir); os.IsNotExist(err) {
		err := os.Mkdir(pathDir, 644)
		if err != nil {
			return err
		}
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
