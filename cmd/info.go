package cmd

import (
	"fmt"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)
  
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information about host or vuln",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" {
			var domain utils.HostIdentity
			domain.Load(fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain))
			
			fmt.Printf("Info Result of the Domain `%s`\n\n", Domain)
			fmt.Println(domain.Info.TechStacks)
		}
	},
}