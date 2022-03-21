package main

import (
	"fmt"
	"log"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	cmd "github.com/k1m0ch1/axolotl/cmd"
	utils "github.com/k1m0ch1/axolotl/utils"
)

func main() {

	var app utils.App
	app.Version = "0.1.4-alpha"

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

	if len(os.Args) == 1 || os.Args[0] == "-h" {
		figure.NewFigure("AXOLOTL", "starwars", true).Print()
		fmt.Printf("\n                    Axolotl - Ez pentest vuln record          ")
		fmt.Printf("\n                        v%s by k1m0ch1", app.Version)
		fmt.Printf("\n                      %d Host and %d Vuln Recorded", len(HostFile), len(VulnFile))
		fmt.Printf("\n                 Info: https://github.com/k1m0ch1/axolotl\n\n")
	} else {
		fmt.Printf("\n──────────────────────────────────────")
		fmt.Printf("\nAxolotl v%s - Ez Vuln Record", app.Version)
		fmt.Printf("\nhttps://github.com/k1m0ch1/axolotl")
		fmt.Printf("\n──────────────────────────────────────\n\n")
	}

	err = cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n\n")
}
