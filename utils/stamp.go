package utils

import (
	"fmt"
	"os"
	"time"

	yaml "github.com/goccy/go-yaml"
)

func Stamp(mode string, cfg UserConfig, host string, vuln string) error {
	fmt.Printf("[+] I will stamp %s to %s.yml", mode, vuln)
	var vulns Finding
	fullPath := fmt.Sprintf("./%s/%s/%s.yml", cfg.DirConfig.VulnDir, host, vuln)
	vulns.Load(fullPath)

	if mode == "created" {
		vulns.Status.Created = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Found the Finding",
		}
	}

	if mode == "reviewed" {
		vulns.Status.Reviewed = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Reviewed",
		}
	}

	if mode == "reported" {
		vulns.Status.Reported = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Reported",
		}
	}

	if mode == "approved" {
		vulns.Status.Approved = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Approved",
		}
	}

	if mode == "fixed" {
		vulns.Status.Fixed = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Fixed",
		}
	}

	if mode == "validated" {
		vulns.Status.Validated = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Validated",
		}
	}

	if mode == "duplicated" {
		vulns.Status.Duplicated = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Duplicated",
		}
	}

	if mode == "hold" {
		vulns.Status.Hold = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Hold",
		}
	}

	if mode == "rejected" {
		vulns.Status.Rejected = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Rejected",
		}
	}

	if mode == "closed" {
		vulns.Status.Closed = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Closed",
		}
	}

	if mode == "completed" {
		vulns.Status.Completed = StatusFields{
			By:   cfg.ProjectOwner,
			Time: time.Now().Format(time.RFC3339),
			Desc: "Finding is Completed",
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
