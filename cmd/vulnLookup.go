package cmd

import (
	"fmt"
	"log"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var vulnLookupCmd = &cobra.Command{
	Use:   "vuln",
	Short: "List all vuln",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		fmt.Println("\nCurrent Available Host Identity:")

		VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}

		for i, f := range VulnFile {
			var v utils.Finding
			v.Load(f)

			fmt.Printf("\n%d. %s (%s)", i+1, v.ID, v.VulnInfo.Domain)
		}
	},
}
