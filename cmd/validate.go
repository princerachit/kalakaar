package cmd

import (
	"fmt"
	"os"

	"github.com/princerachit/kalakaar/pkg/converter"
	"github.com/princerachit/kalakaar/pkg/input_config"
	"github.com/spf13/cobra"
)

func init() {
	var config string
	// validateCmd represents the validate command
	var validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Validate the config file",
		Long: `This command will validate the config file
and will return errors if any.`,
		Run: func(cmd *cobra.Command, args []string) {
			data, err := os.ReadFile(config)
			if err != nil {
				fmt.Println("Error reading the configuration file:", err)
				return
			}

			k, err := input_config.Parse(data)
			if err != nil {
				fmt.Println("Error parsing the configuration file:", err)
			}

			p, err := converter.InputConfigToConfiguration(k)
			if err != nil {
				fmt.Println("Error converting the configuration file:", err)
			}
			errs := p.Validate()
			if len(errs) > 0 {
				fmt.Println("Validation failed:")
				for _, err := range errs {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Validation successful.")
			}
		},
	}
	rootCmd.AddCommand(validateCmd)

	// todo: add support for a flag that accepts some token and validate the config further
	// e.g. it can validate if the current plan of the user would support the features used in the config
	validateCmd.Flags().StringVarP(&config, "config", "f", "kaarkala.yaml", "Input Configuration File")
}
