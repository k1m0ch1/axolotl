package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "Mode search",
	Long:  `Need to add`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		fmt.Printf("Search Result record with ")

		found := 0
		if Vuln != "" {
			VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Vulnerability Type `%s`:\n", Vuln)
			for _, f := range VulnFile {
				var v utils.Finding
				v.Load(f)
				ListVulnRaw := strings.ReplaceAll(v.VulnInfo.VulnType, " ", "")
				ListVulns := strings.Split(ListVulnRaw, ",")
				for _, f := range ListVulns {
					if f == Vuln {
						fmt.Printf("\n[w00t] %s is have %s vuln type with finding %s", v.VulnInfo.Domain, Vuln, v.ID)
						found = found + 1
					}
				}
			}
		}

		HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}

		if TechStacks != "" {
			fmt.Printf("Technology Stack `%s`:\n", TechStacks)
		}

		if Tags != "" {
			fmt.Printf("Tag `%s`:\n", Tags)
		}

		if Ports != "" {
			fmt.Printf("Port `%s`:\n", Ports)
		}

		for _, f := range HostFile {
			var h utils.HostIdentity
			h.Load(f)
			if TechStacks != "" {
				ListStacksRaw := strings.ReplaceAll(h.Info.TechStacks, " ", "")
				ListStacks := strings.Split(ListStacksRaw, ",")
				for _, f := range ListStacks {
					tSA := TechStacks
					if strings.Contains(tSA, ":") {
						parseLagi := strings.Split(ListStacksRaw, ",")
						tSA = parseLagi[0]
					}
					if f == tSA {
						fmt.Printf("\n[w00t] %s is used %s stack", h.ID, f)
						found = found + 1
					}
				}
			}

			if Tags != "" {
				ListTagsRaw := strings.ReplaceAll(h.Info.Tag, " ", "")
				ListTags := strings.Split(ListTagsRaw, ",")
				for _, f := range ListTags {
					tA := Tags
					if f == tA {
						fmt.Printf("\n[w00t] %s is have %s tag", h.ID, f)
						found = found + 1
					}
				}
			}

			if Ports != "" {
				ListPortRaw := strings.ReplaceAll(h.Info.OpenPorts, " ", "")
				ListPorts := strings.Split(ListPortRaw, ",")
				for _, f := range ListPorts {
					splitPorts := strings.Split(f, "/")
					for _, p := range splitPorts {
						if p == Ports {
							fmt.Printf("\n[w00t] %s is have %s port", h.ID, p)
							found = found + 1
						}
					}
				}
			}
		}

		fmt.Printf("\n\n%d Result", found)
	},
}
