package main

import(
	"log"
	"fmt"
	
	utils "github.com/k1m0ch1/axolotl/utils"
)

func main(){
	
	var cfg utils.UserConfig
	cfg.Load("config.yml")

	HostFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.HostsIdentityDir), "*.yml")
	for _, f := range HostFile {
		var h utils.HostIdentity
		h.Load(f)
		fmt.Printf("\nHost: %s", h.Info.URL)
		fmt.Printf("\nTechStack: %s\n", h.Info.TechStacks)
	}

	VulnFile, err := utils.WalkMatch(fmt.Sprintf("./%s/", cfg.DirConfig.VulnDir), "*.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(HostFile), " Host Available")
	fmt.Println(len(VulnFile), " Vuln Available")

	GenerateHost := cfg.GenerateHost("mantap.com")
	if GenerateHost != nil {
		log.Fatal(err)
	}

	GenerateVuln := cfg.GenerateVuln("mantap.com", "the-attack-on-titan")
	if GenerateVuln != nil {
		log.Fatal(err)
	}
}

