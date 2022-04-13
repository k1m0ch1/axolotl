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
		domainList := make(map[string]int)

		for _, f := range VulnFile {
			var v utils.Finding
			v.Load(f)

			domainList[v.VulnInfo.Domain] += 1

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
							// fmt.Println(v.ID, "Already exist at", vulnTypeStat[currIndex].Type, "with data", vulnTypeStat[currIndex].ListOfVuln, "with index", checkVuln)
						} else {
							vulnTypeStat[currIndex].ListOfVuln = append(vulnTypeStat[currIndex].ListOfVuln, v.ID)
						}
					}

				}
			}
		}

		fmt.Println("Top 10 Host with Vulnerability Finding")

		type kv struct {
			Key   string
			Value int
		}

		var rankedDomain []kv
		for k, v := range domainList {
			rankedDomain = append(rankedDomain, kv{k, v})
		}

		sort.Slice(rankedDomain, func(i, j int) bool {
			return rankedDomain[i].Value > rankedDomain[j].Value
		})

		maxSlice := len(rankedDomain)
		if len(rankedDomain) > 10 {
			maxSlice = 10
		}

		for index, value := range rankedDomain[0:maxSlice] {
			fmt.Printf("\n%d. %s (%d finding)", index+1, value.Key, value.Value)
		}

		fmt.Println("\n\nTop 10 Vulnerability Type Finding")
		maxSlice = len(vulnTypeStat)
		if len(vulnTypeStat) > 10 {
			maxSlice = 10
		}

		var rankedVuln []kv
		for _, v := range vulnTypeStat {
			rankedVuln = append(rankedVuln, kv{v.Type, len(v.ListOfVuln)})
		}

		sort.Slice(rankedVuln, func(i, j int) bool {
			return rankedVuln[i].Value > rankedVuln[j].Value
		})

		for index, value := range rankedVuln[0:maxSlice] {
			fmt.Printf("\n%d. `%s` with %d vuln", index+1, value.Key, value.Value)
		}
	},
}
