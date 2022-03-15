package utils

import (
	"os"
	"fmt"
	"time"

	yaml "github.com/goccy/go-yaml"
)

func Stamp(mode string, cfg UserConfig, host string, vuln string) error{
	fmt.Printf("[+] I will stamp %s to %s.yml", mode, vuln)
	var vulns Finding
	fullPath := fmt.Sprintf("./%s/%s/%s.yml", cfg.DirConfig.VulnDir, host, vuln)
	vulns.Load(fullPath)

	if mode == "created"{
		vulns.Status.Created = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "reviewed"{
		vulns.Status.Reviewed = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "reported"{
		vulns.Status.Reported = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "approved"{
		vulns.Status.Approved = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "fixed"{
		vulns.Status.Fixed = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "validated"{
		vulns.Status.Validated = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "duplicated"{
		vulns.Status.Duplicated = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "hold"{
		vulns.Status.Hold = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "rejected"{
		vulns.Status.Rejected = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "closed"{
		vulns.Status.Closed = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}

	if mode == "completed"{
		vulns.Status.Completed = StatusFields{
			By: cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "found finding",
		}
	}
	

	bytes, err := yaml.Marshal(vulns)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}