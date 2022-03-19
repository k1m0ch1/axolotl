package main

import(
	"log"
	"fmt"		
	
	cmd "github.com/k1m0ch1/axolotl/cmd"
	utils "github.com/k1m0ch1/axolotl/utils"
	figure "github.com/common-nighthawk/go-figure"
)

func main(){

	var cfg utils.UserConfig
	cfg.Load("config.yml")

	var app utils.App
	app.Version = "0.1.0-alpha"

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

	figure.NewFigure("AXOLOTL", "starwars", true).Print()

	fmt.Printf("\n                    Axolotl - Ez pentest vuln record          ")
	fmt.Printf("\n                        v%s by k1m0ch1", app.Version)
	fmt.Printf("\n                      %d Host and %d Vuln Recorded", len(HostFile), len(VulnFile))
	fmt.Printf("\n                 Info: https://github.com/k1m0ch1/axolotl\n\n") 

	cmd.Execute()

}