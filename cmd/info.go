package cmd

import (
	"os"
	"fmt"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about host or vuln",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln == "" {
			var domain utils.HostIdentity
			domain.Load(fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain))

			fmt.Printf("Info Result of the Domain `%s`\n", Domain)
			fmt.Printf("\nDomain `%s` %s (%s)", domain.Info.URL, domain.Info.HostIP, domain.Info.OpenPorts)
			fmt.Printf("\n%s\n", domain.Info.Desc)
			fmt.Printf("\nTechnology : %s ", domain.Info.TechStacks)

			fmt.Printf("\n\nCurrent Vulnerability : ")

			path := fmt.Sprintf("./%s/%s/", cfg.DirConfig.VulnDir, domain.ID)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("0 Result Vulnerability")
			}else{
				// get many vuln-type found
				VulnFile, err := utils.WalkMatch(path, "*.yml")
				if err != nil {
					fmt.Println(err)
				}

				for index, f := range VulnFile {
					var v utils.Finding
					v.Load(f)

					fmt.Printf("\n%d. %s\n   %s (%s)", index+1, v.ID, v.VulnInfo.RiskRatings.Severity, v.VulnInfo.VulnType)
				}
			}
		}
	},
}
