package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	ProjectName string `json:"project_name"`
}

var rootCommand = &cobra.Command{
	Use:   "cld",
	Short: "Generate monorepo project with CI/CD setup",
	Run: func(cmd *cobra.Command, args []string) {
		configFile := "config.json"
		content, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Printf("Error reading config: %v\n", err)
			os.Exit(1)
		}

		var cfg Config
		if err := json.Unmarshal(content, &cfg); err != nil {
			fmt.Printf("Error parsing config: %v\n", err)
			os.Exit(1)
		}

		if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
			fmt.Printf("Error creating project dir: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ Project scaffolded successfully.")

	},
}

func Execute() {
	cobra.CheckErr(rootCommand.Execute())
}
