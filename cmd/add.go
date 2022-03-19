package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add mode you can create new host",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln == "" {

			pathDir := fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain)
			if _, err := os.Stat(pathDir); os.IsNotExist(err) {
				GenerateHost := cfg.GenerateHost(Domain)
				if GenerateHost != nil {
					log.Fatal(err)
				}
				fmt.Printf("\n[+] Host %s is Created at ./%s", Domain, cfg.DirConfig.HostsIdentityDir)
			} else {
				fmt.Printf("\n[?] Warning! Host Identity for %s is already exist at ./%s\n", Domain, cfg.DirConfig.HostsIdentityDir)
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Are you sure you want to replace the current Host Identity %s (y/N): ", Domain)
				text, _, _ := reader.ReadRune()
				resultText := fmt.Sprintf("%c", unicode.ToLower(text))
				if resultText == "n" {

				} else if resultText == "y" {
					GenerateHost := cfg.GenerateHost(Domain)
					if GenerateHost != nil {
						log.Fatal(err)
					}
					fmt.Printf("\n[+] Host %s is Created at ./%s", Domain, cfg.DirConfig.HostsIdentityDir)
				}
			}
		}

		if Domain != "" && Vuln != "" {
			pathDir := fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain)
			if _, err := os.Stat(pathDir); os.IsNotExist(err) {
				fmt.Printf("\n[*] Warning! Host Identity for %s is not exist at ./%s", Domain, cfg.DirConfig.HostsIdentityDir)
			}
			GenerateVuln := cfg.GenerateVuln(Domain, Vuln)
			if GenerateVuln != nil {
				log.Fatal("ERROR GENERATE VULNERABILITY")
			}
			fmt.Printf("\n[+] File %s.yml is generated at ./%s, Happy Hacking!", Vuln, cfg.DirConfig.VulnDir)
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}
