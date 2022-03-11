package main

import(
	"os"
	"log"
	"fmt"
	"flag"
	"bufio"
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
	fmt.Printf("\n|| Ez pentest documentation")
	fmt.Printf("\n|| %d Host and %d Vuln Available", len(HostFile), len(VulnFile)) 
	fmt.Printf("\n==================================\n")

	insertArg := flag.Bool("i", false, "Insert Mode")
	searchArg := flag.Bool("s", true, "Search Mode")
	techStackArg := flag.String("ts", "python", "Tech Stack to search")
	hostArg := flag.String("host", "", "Hostname")
	vulnArg := flag.String("vn", "", "Attack name or vulnerability name")

	flag.Parse()

	if *hostArg == "" && *vulnArg == "" && *techStackArg == "" {
		fmt.Println("\ntype -h to see the available commands")
	}

	if *searchArg == true {
		HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
		if err != nil {
			log.Fatal(err)
		}
		found := 0
		for _, f := range HostFile {
			var h utils.HostIdentity
			h.Load(f)
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

		fmt.Printf("\n\n%d Result", found)
	}

	if *insertArg == true{
		*searchArg = false

		if *hostArg != "" {
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

