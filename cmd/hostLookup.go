package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var hostLookupCmd = &cobra.Command{
	Use:   "host",
	Short: "List all host",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		fmt.Println("\nCurrent Available Host Identity:")

		HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}

		for i, f := range HostFile {
			var h utils.HostIdentity
			h.Load(f)

			fmt.Printf("\n%d. %s ", i+1, h.Info.URL)

			path := fmt.Sprintf("./%s/%s/", cfg.DirConfig.VulnDir, h.ID)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Printf("(0 vuln)")
			} else {
				// get many vuln-type found
				VulnFile, err := utils.WalkMatch(path, "*.yml")
				if err != nil {
					fmt.Println(err)
				}

				currVuln := 0
				for _, f := range VulnFile {
					var v utils.Finding
					v.Load(f)
					currVuln += 1
				}
				fmt.Printf("(%d vuln)", currVuln)
			}
		}
	},
}
