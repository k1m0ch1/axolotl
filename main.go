package main

import(
	"os"
	"log"
	"fmt"
	"flag"
	"bufio"
	"errors"
	"strings"
	"unicode"
	
	utils "github.com/k1m0ch1/axolotl/utils"
)

func main(){
	
	var cfg utils.UserConfig
	cfg.Load("config.yml")

	HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range HostFile {
		var h utils.HostIdentity
		h.Load(f)
	}

	VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("\n==================================")
	fmt.Printf("\n|| Axolotl ")
	fmt.Printf("\n|| Ez pentest findings management")
	fmt.Printf("\n|| %d Host and %d Vuln Available", len(HostFile), len(VulnFile)) 
	fmt.Printf("\n==================================\n")

	lookupMode := flag.Bool("l", true, "Lookup Mode")
	insertMode := flag.Bool("i", false, "Insert Mode")

	hostArg := flag.String("host", "", "Hostname")
	vulnArg := flag.String("vn", "", "Attack name or vulnerability name")

	techStackArg := flag.String("ts", "", "Tech Stack to search")
	tagsArg := flag.String("tag", "", "Tag to search")
	portArg := flag.String("port", "", "Port to search")

	flag.Parse()
	otherArg := flag.Args()

	if *techStackArg == "" && *tagsArg == "" && *portArg == ""{
		fmt.Println("add argument -h to see the help command")
	}

	if len(otherArg) > 0 {
		if otherArg[0] == "init" {
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
		}
	}

	if *lookupMode == true {
		found := 0
		if *vulnArg != "" {
			VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range VulnFile {
				var v utils.Finding
				v.Load(f)
				ListVulnRaw := strings.ReplaceAll(v.VulnInfo.VulnType, " ", "")
				ListVulns := strings.Split(ListVulnRaw, ",")
				for _, f := range ListVulns {
					if f == *vulnArg {
						fmt.Printf("\n[w00t] %s is have %s vuln type with finding %s", v.VulnInfo.Domain, *vulnArg, v.ID)
						found = found + 1
					}					
				}
			}
		}

		HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}
		
		for _, f := range HostFile {
			var h utils.HostIdentity
			h.Load(f)
			if *techStackArg != "" {
				ListStacksRaw := strings.ReplaceAll(h.Info.TechStacks, " ", "")
				ListStacks := strings.Split(ListStacksRaw, ",")
				for _, f := range ListStacks {
					tSA := *techStackArg
					if strings.Contains(tSA, ":") == true {
						parseLagi := strings.Split(ListStacksRaw, ",")
						tSA = parseLagi[0]
					}
					if f == tSA {
						fmt.Printf("\n[w00t] %s is used %s stack", h.ID, f)
						found = found + 1
					}
				}
			}

			if *tagsArg != "" {
				ListTagsRaw := strings.ReplaceAll(h.Info.Tag, " ", "")
				ListTags := strings.Split(ListTagsRaw, ",")
				for _, f := range ListTags {
					tA := *tagsArg
					if f == tA {
						fmt.Printf("\n[w00t] %s is have %s tag", h.ID, f)
						found = found + 1
					}
				}
			}

			if *portArg != "" {
				ListPortRaw := strings.ReplaceAll(h.Info.OpenPorts, " ", "")
				ListPorts := strings.Split(ListPortRaw, ",")
				for _, f := range ListPorts {
					splitPorts := strings.Split(f, "/")
					for _, p := range splitPorts {
						if p == *portArg {
							fmt.Printf("\n[w00t] %s is have %s port", h.ID, p)
							found = found + 1
						}
					}					
				}
			}
		}

		fmt.Printf("\n\n%d Result", found)
	}

	if *insertMode == true{
		*lookupMode = false

		if *hostArg != "" && *vulnArg == "" {
			pathDir := fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, *hostArg)
			if _, err := os.Stat(pathDir); os.IsNotExist(err) {
				GenerateHost := cfg.GenerateHost(*hostArg)
				if GenerateHost != nil {
					log.Fatal(err)
				}
				fmt.Printf("\n[+] Host %s is Created at ./%s", *hostArg, cfg.DirConfig.HostsIdentityDir)
			}else{
				fmt.Printf("\n[?] Warning! Host Identity for %s is already exist at ./%s\n", *hostArg, cfg.DirConfig.HostsIdentityDir)
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("Are you sure you want to replace the current Host Identity %s (y/N): ", *hostArg)
				text, _, _ := reader.ReadRune()
				resultText := fmt.Sprintf("%c", unicode.ToLower(text))
				if resultText == "n" {
					
				}else if resultText == "y" {
					GenerateHost := cfg.GenerateHost(*hostArg)
					if GenerateHost != nil {
						log.Fatal(err)
					}
					fmt.Printf("\n[+] Host %s is Created at ./%s", *hostArg, cfg.DirConfig.HostsIdentityDir)
				}
			}
		}

		if *vulnArg != "" && *hostArg == "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}else if *vulnArg != "" && *hostArg != "" {
			pathDir := fmt.Sprintf("./%s/%s.yml", cfg.DirConfig.HostsIdentityDir, *hostArg)
			if _, err := os.Stat(pathDir); os.IsNotExist(err) {
				fmt.Printf("\n[*] Warning! Host Identity for %s is not exist at ./%s", *hostArg, cfg.DirConfig.HostsIdentityDir)
			}
			GenerateVuln := cfg.GenerateVuln(*hostArg, *vulnArg)
			if GenerateVuln != nil {
				log.Fatal(err)
			}
			fmt.Printf("\n[+] File %s.yml is generated at ./%s, Happy Hacking!", *vulnArg, cfg.DirConfig.VulnDir)
		}
	}

	fmt.Println("\n\nBye!\n")
}

