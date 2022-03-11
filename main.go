package main

import(
	"os"
	"log"
	"fmt"
	"flag"
	
	utils "github.com/k1m0ch1/axolotl/utils"
)

func main(){
	
	var cfg utils.UserConfig
	cfg.Load("config.yml")

	HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
	for _, f := range HostFile {
		var h utils.HostIdentity
		h.Load(f)
		// fmt.Printf("\nHost: %s", h.Info.URL)
		// fmt.Printf("\nTechStack: %s\n", h.Info.TechStacks)
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
	hostArg := flag.String("host", "", "Hostname")
	vulnArg := flag.String("vn", "", "Attack name or vulnerability name")

	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("\ntype -h to see the available commands")
	}

	if *insertArg == true{
		*searchArg = false
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

	// GenerateHost := cfg.GenerateHost("mantap.com")
	// if GenerateHost != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("\n\nBye!\n")
}

