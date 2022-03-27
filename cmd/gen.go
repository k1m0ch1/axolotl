package cmd

import (
	"os"
	"fmt"
	"log"
	"strings"
	"os/exec"
	"io/ioutil"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

var (
	Output string
	Template string
)

func init(){
	genCmd.PersistentFlags().StringVarP(&Output, "out", "O", "", "File Output")
	genCmd.PersistentFlags().StringVarP(&Template, "template", "T", "", "Markdown Template")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate report from Markdown Template into PDF Format",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		_, err := exec.LookPath("pandoc")
		if err != nil {
			log.Println("[!] you need to install `pandoc` on your machine")
			log.Println("[!] here is how to install https://pandoc.org/installing.html")
			log.Println("[!] or you can run the file install-pandoc.sh in here:")
			log.Println("[!] https://github.com/k1m0ch1/axolotl/blob/master/install-pandoc.sh")
			log.Fatal(err)
		}
		
		out, err := exec.Command("pandoc", "-v").Output()
		if err != nil {
			log.Fatal(err)
		}

		Pandoc := strings.Split(string(out),"\n")
		pandorDirConfig := strings.ReplaceAll(strings.Split(Pandoc[3], ":")[1], " ", "")
		templateIsExist := fmt.Sprintf("%s/templates/eisvogel.latex", pandorDirConfig)
		if _, err := os.Stat(templateIsExist); os.IsNotExist(err) {
			log.Printf("[!] File template eisvogel does not exist at %s/templates", pandorDirConfig)
			log.Println("[!] you need to download file templates from here :")
			log.Println("[!] https://github.com/Wandmalfarbe/pandoc-latex-template/blob/master/eisvogel.tex")
			log.Printf("[!] and put the file as %s/templates/eisvogel.latex", pandorDirConfig)
		}else{
			if Output != "" && Template != "" {
				var domain utils.HostIdentity
				domain.Load(fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain))

				content, _ := ioutil.ReadFile(Template)

				sContent := string(content)

				sContent = strings.ReplaceAll(sContent, "{{Host.Info.URL}}", domain.Info.URL)
				sContent = strings.ReplaceAll(sContent, "{{Host.Info.Desc}}", domain.Info.Desc)
				sContent = strings.ReplaceAll(sContent, "{{Host.Info.HostIP}}", domain.Info.HostIP)
				sContent = strings.ReplaceAll(sContent, "{{UserConfig.ProjectOwner}}", cfg.ProjectOwner)
				sContent = strings.ReplaceAll(sContent, "{{UserConfig.Email}}", cfg.Email)
				
				tmpFile, err := ioutil.TempFile("", "axolotl-templates*.md")
				if err != nil {
					log.Fatal(err)
				}

				if _, err := tmpFile.Write([]byte(sContent)); err != nil {
					fmt.Println(err)
				}
			
				_, err = ioutil.ReadFile(tmpFile.Name())
				if err != nil {
					fmt.Println(err)
				}				
			
				_, err = exec.Command("pandoc", tmpFile.Name(), "-o", Output, "--from", "markdown+yaml_metadata_block+raw_html", "--template", "eisvogel", "--table-of-contents", "--toc-depth", "6", "--number-sections", "--top-level-division=chapter", "--highlight-style", "breezedark").Output()
				if err != nil {
					log.Fatal(err)
				}

				defer os.Remove(tmpFile.Name())
			}
		}
	},
}
