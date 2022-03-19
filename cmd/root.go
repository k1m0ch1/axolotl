package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Domain     string
	Vuln       string
	TechStacks string
	Tags       string
	Ports      string

	rootCmd = &cobra.Command{
		Use:   "axolotl",
		Short: "Axolotl is the productivity tools for bug bounty hackers to document the bug finding",
		Long: `A simple bug reporting tools for hackers to input the finding Inspired 
from nuclei. Complete documentation is available at http://k1m0ch1.github.io`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", "Set the Domain name")
	rootCmd.PersistentFlags().StringVarP(&Vuln, "vuln", "v", "", "Set the Vulnerability name")
	rootCmd.PersistentFlags().StringVarP(&TechStacks, "tech-stack", "", "", "Set the Techstack name")
	rootCmd.PersistentFlags().StringVarP(&Tags, "tag", "t", "", "Set the tag name")
	rootCmd.PersistentFlags().StringVarP(&Ports, "port", "p", "", "Set the Port name")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(lookupCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(infoCmd)
}
