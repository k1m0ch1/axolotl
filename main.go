package main

import(
	"os"
	"log"
	"fmt"		
	
	cmd "github.com/k1m0ch1/axolotl/cmd"
	utils "github.com/k1m0ch1/axolotl/utils"
)

func main(){

	cmd.Execute()

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

	os.Exit(0)
	
	fmt.Printf("\n==================================")
	fmt.Printf("\n|| Axolotl ")
	fmt.Printf("\n|| Ez pentest findings management")
	fmt.Printf("\n|| %d Host and %d Vuln Available", len(HostFile), len(VulnFile)) 
	fmt.Printf("\n==================================\n\n")
	
	fmt.Println("\n\nBye!\n")
}