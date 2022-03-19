package cmd

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stat",
	Short: "Simple statistic of the Vulnerabilities",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		// get many vuln-type found
		VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}

		var vulnTypeStat []utils.VulnType

		for _, f := range VulnFile {
			var v utils.Finding
			v.Load(f)
			ListVulnRaw := strings.ReplaceAll(v.VulnInfo.VulnType, " ", "")
			ListVulns := strings.Split(ListVulnRaw, ",")
			for _, value := range ListVulns {
				countType := 0
				currIndex := 999
				for vulnIndex, vi := range vulnTypeStat {
					if value == vi.Type {
						countType = countType + 1
						currIndex = vulnIndex
					}
				}
				if countType == 0 {
					vulnTypeStat = append(vulnTypeStat,
						utils.VulnType{
							Type:       value,
							ListOfVuln: []string{v.ID},
						},
					)
				} else {
					checkVuln := sort.SearchStrings(vulnTypeStat[currIndex].ListOfVuln, v.ID)
					if checkVuln >= len(vulnTypeStat[currIndex].ListOfVuln) {
						vulnTypeStat[currIndex].ListOfVuln = append(vulnTypeStat[currIndex].ListOfVuln, v.ID)
					} else {
						if v.ID == vulnTypeStat[currIndex].ListOfVuln[checkVuln] {
							fmt.Println(v.ID, "Already exist at", vulnTypeStat[currIndex].Type, "with data", vulnTypeStat[currIndex].ListOfVuln, "with index", checkVuln)
						} else {
							vulnTypeStat[currIndex].ListOfVuln = append(vulnTypeStat[currIndex].ListOfVuln, v.ID)
						}
					}

				}
			}
		}

		fmt.Println("Vulnerability Type")
		currMin := 0
		currMax := 0
		for index, value := range vulnTypeStat {
			fmt.Printf("\n%s with %d vuln", value.Type, len(value.ListOfVuln))
			if len(value.ListOfVuln) < len(vulnTypeStat[currMin].ListOfVuln) {
				currMin = index
			}

			if len(value.ListOfVuln) > len(vulnTypeStat[currMax].ListOfVuln) {
				currMax = index
			}
		}
		fmt.Printf("\n")
		fmt.Printf("\nwith the %s as the most vulnerability you found (%d vuln)", vulnTypeStat[currMax].Type, len(vulnTypeStat[currMax].ListOfVuln))
		fmt.Printf("\nand %s as the least you found (%d vuln)", vulnTypeStat[currMin].Type, len(vulnTypeStat[currMin].ListOfVuln))
	},
}
