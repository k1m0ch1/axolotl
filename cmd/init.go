package cmd

import (
	"os"
	"fmt"
	"errors"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)


var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate the new project",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n[*] Generating new template")
		if _, err := os.Stat("./config.yml"); err == nil {
			var uC utils.UserConfig
			fmt.Println("[+] File config.yml existed, use the current config")
			uC.Load("./config.yml")
			fmt.Printf("\n[+] Hi %s", uC.ProjectOwner)

			dirs := []string{
				uC.DirConfig.HostsIdentityDir, 
				uC.DirConfig.VulnDir,
				uC.DirConfig.ToolsReports,
				uC.DirConfig.PocDir,
				uC.DirConfig.OutputReportsDir,
				uC.DirConfig.TemplatesReportDir,
			}
			
			for _, v := range dirs {
				path := fmt.Sprintf("./%s", v)
				_, err := utils.CheckDirAndCreate(path)
				if err != nil {
					fmt.Println("[X] Weird error occured ", err)
				}
			}

		} else if errors.Is(err, os.ErrNotExist) {
			var uC utils.UserConfig
			var name string
			fmt.Println("[!] File config.yml not exist, generate the template config")
			fmt.Printf("\n[+] Tell me your name/nick: ")
			fmt.Scanln(&name)
			uC.GenerateConfig(name)

			uC.Load("./config.yml")
			
			dirs := []string{
				uC.DirConfig.HostsIdentityDir, 
				uC.DirConfig.VulnDir,
				uC.DirConfig.ToolsReports,
				uC.DirConfig.PocDir,
				uC.DirConfig.OutputReportsDir,
				uC.DirConfig.TemplatesReportDir,
			}
			
			for _, v := range dirs {
				path := fmt.Sprintf("./%s", v)
				_, err := utils.CheckDirAndCreate(path)
				if err != nil {
					fmt.Println("[X] Weird error occured ", err)
				}
			}
		} else {
			// Schrodinger: file may or may not exist. See err for details.
			// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
			panic(err)
		}
	},
}