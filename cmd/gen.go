package cmd

import (
	"os"
	"fmt"
	"log"
	"bytes"
	"strings"
	"os/exec"
	"io/ioutil"
	"text/template"

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
				var ReportVar utils.ReportVar
				var buf bytes.Buffer
				var vulns []utils.Finding
				domain.Load(fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, Domain))
				pathVuln := fmt.Sprintf("./%s/%s/", cfg.DirConfig.VulnDir, domain.ID)

				if _, err := os.Stat(pathVuln); os.IsNotExist(err) {
					fmt.Println("0 Result Vulnerability")
					os.Exit(0)
				}else{
					// get many vuln-type found
					VulnFile, err := utils.WalkMatch(pathVuln, "*.yml")
					if err != nil {
						fmt.Println(err)
					}
	
					for _, f := range VulnFile {
						var v utils.Finding
						v.Load(f)
						vulns = append(vulns, v)
					}
				}

				ReportVar.Host = domain
				ReportVar.User = cfg
				ReportVar.Vulns = vulns
				content, _ := ioutil.ReadFile(Template)

				sContent := string(content)

				tmpl, err := template.New("test").Funcs(template.FuncMap{
					"add": func(a, b int) int {
						return a + b
					},
				}).Parse(sContent)
				if err != nil {
					log.Fatal(err)
				}
				err = tmpl.Execute(&buf, ReportVar)
				if err != nil {
					log.Fatal(err)
				}

				result := buf.String()
				
				tmpFile, err := ioutil.TempFile("", "axolotl-templates*.md")
				if err != nil {
					log.Fatal(err)
				}

				if _, err := tmpFile.Write([]byte(result)); err != nil {
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
