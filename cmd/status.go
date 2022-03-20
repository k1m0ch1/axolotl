package cmd

import (
	"fmt"

	"github.com/k1m0ch1/axolotl/utils"
	"github.com/spf13/cobra"
)

func init() {
	statusCmd.AddCommand(acceptedCmd)
	statusCmd.AddCommand(reportedCmd)
	statusCmd.AddCommand(approvedCmd)
	statusCmd.AddCommand(fixedCmd)
	statusCmd.AddCommand(validatedCmd)
	statusCmd.AddCommand(duplicatedCmd)
	statusCmd.AddCommand(holdCmd)
	statusCmd.AddCommand(rejectedCmd)
	statusCmd.AddCommand(reviewedCmd)
	statusCmd.AddCommand(closedCmd)
	statusCmd.AddCommand(completedCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Simple stats of the Vulnerabilities",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		fmt.Println("WOW")
	},
}

var acceptedCmd = &cobra.Command{
	Use:   "accepted",
	Short: "Timestamp with status ACCEPTED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("created", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var reviewedCmd = &cobra.Command{
	Use:   "reviewed",
	Short: "Timestamp with status REVIEWED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("reviewed", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var reportedCmd = &cobra.Command{
	Use:   "reported",
	Short: "Timestamp with status REPORTED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("reported", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var approvedCmd = &cobra.Command{
	Use:   "approved",
	Short: "Timestamp with status APPROVED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("approved", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var fixedCmd = &cobra.Command{
	Use:   "fixed",
	Short: "Timestamp with status FIXED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("fixed", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var validatedCmd = &cobra.Command{
	Use:   "validated",
	Short: "Timestamp with status VALIDATED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("validated", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var duplicatedCmd = &cobra.Command{
	Use:   "duplicated",
	Short: "Timestamp with status DUPLICATED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("duplicated", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var holdCmd = &cobra.Command{
	Use:   "hold",
	Short: "Timestamp with status HOLD",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("hold", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var rejectedCmd = &cobra.Command{
	Use:   "rejected",
	Short: "Timestamp with status REJECTED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("rejected", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var closedCmd = &cobra.Command{
	Use:   "closed",
	Short: "Timestamp with status CLOSED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("closed", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Timestamp with status COMPLETED",
	Long:  `Generate the new project include with dirs and config.yml`,
	Run: func(cmd *cobra.Command, args []string) {
		var cfg utils.UserConfig
		cfg.Load("config.yml")

		if Domain != "" && Vuln != "" {
			err := utils.Stamp("completed", cfg, Domain, Vuln)
			if err != nil {
				fmt.Println(err)
			}
		}

		if Domain == "" && Vuln != "" {
			fmt.Println("you must add hostname to create vulnerability record")
		}
	},
}
