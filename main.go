package main

import(
	"os"
	"log"
	"fmt"
	"flag"
	"sort"
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
	fmt.Printf("\n==================================\n\n")

	lookupMode := flag.Bool("l", true, "Lookup Mode")
	insertMode := flag.Bool("i", false, "Insert Mode")
	createdStamp := flag.Bool("created", false, "timestamp to be created (when the first time you found the vuln")
	reviewedStamp := flag.Bool("reviewed", false, "timestamp to be reviewed (when you give the finding to review")
	reportedStamp := flag.Bool("reported", false, "timestamp to be reported (when you send the report to vendor)")
	approvedStamp := flag.Bool("approved", false, "timestamp to be approved (when the vuln is accepted by the vendor")
	fixedStamp := flag.Bool("fixed", false, "timestamp to be fixed (when the vuln is fixed by the vendor")
	validatedStamp := flag.Bool("validated", false, "timestamp to be validated (when the vuln is fixed and validated by bug founder")
	duplicatedStamp := flag.Bool("duplicated", false, "timestamp to be duplicated (when the vuln is responded duplicate by the vendor")
	holdStamp := flag.Bool("hold", false, "timestamp to be hold (when the vuln is hold by the vendor")
	rejectedStamp := flag.Bool("rejected", false, "timestamp to be rejected (when the vuln is rejected by the vendor")
	closedStamp := flag.Bool("closed", false, "timestamp to be closed (when the vuln is closed or stop without any progress by the vendor")
	completedStamp := flag.Bool("completed", false, "timestamp to be completed (when the vuln is done and complete")

	hostArg := flag.String("host", "", "Hostname")
	vulnArg := flag.String("vn", "", "Attack name or vulnerability name")

	techStackArg := flag.String("ts", "", "Tech Stack to search")
	tagsArg := flag.String("tag", "", "Tag to search")
	portArg := flag.String("port", "", "Port to search")

	haveArg := false

	flag.Parse()
	otherArg := flag.Args()

	var looking = []bool{
		*lookupMode, *insertMode, *createdStamp, *reviewedStamp, *reportedStamp,
		*approvedStamp, *fixedStamp, *validatedStamp, *duplicatedStamp,
		*holdStamp, *rejectedStamp, *closedStamp, *completedStamp, 
	}
	countTrue := 0
	for  _, v := range looking {
		if v == true {
			countTrue = countTrue + 1
		}
	}

	if countTrue > 1 {
		*lookupMode = false
	}

	if len(otherArg) > 0 {

		if otherArg[0] == "stat" {
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
								Type: value,
								ListOfVuln: []string{v.ID},
							},
						)
					}else{						
						checkVuln := sort.SearchStrings(vulnTypeStat[currIndex].ListOfVuln, v.ID)
						if checkVuln >= len(vulnTypeStat[currIndex].ListOfVuln) {
							vulnTypeStat[currIndex].ListOfVuln = append(vulnTypeStat[currIndex].ListOfVuln, v.ID)
						}else{
							if v.ID==vulnTypeStat[currIndex].ListOfVuln[checkVuln] {
								fmt.Println(v.ID, "Already exist at", vulnTypeStat[currIndex].Type, "with data", vulnTypeStat[currIndex].ListOfVuln, "with index", checkVuln)
							}else{
								vulnTypeStat[currIndex].ListOfVuln = append(vulnTypeStat[currIndex].ListOfVuln, v.ID)
							}
						}
						
					}
				}
			}

			fmt.Println("Vulnerability Type")
			currMin := 0
			currMax := 0
			for index, value := range vulnTypeStat{
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
			fmt.Printf("\nand %s as the least you found (%d vuln)", vulnTypeStat[currMin].Type, len(vulnTypeStat[currMin].ListOfVuln) )
		}
		
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

		haveArg = true
	}

	if *lookupMode == true && haveArg == false{
		
		if *techStackArg == "" && *tagsArg == "" && *portArg == "" {
			fmt.Println("add argument -h to see the help command")
		}else{
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

	if *createdStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("created", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *reviewedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("reviewed", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *reportedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("reported", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *approvedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("approved", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *fixedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("fixed", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *validatedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("validated", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *duplicatedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("duplicated", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *holdStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("hold", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *rejectedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("rejected", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *closedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("closed", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	if *completedStamp == true {
		if *vulnArg != "" && *hostArg != "" {
			utils.Stamp("completed", cfg, *hostArg, *vulnArg)
		}else{
			fmt.Println("[!] Please define -host and -vn")
		}
	}

	fmt.Println("\n\nBye!\n")
}

