package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"sinhasoftware.solutions/kalakaar/pkg/configuration"
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
			data, err := ioutil.ReadFile(config)
			if err != nil {
				fmt.Println("Error reading the configuration file:", err)
				return
			}

			k, err := configuration.Parse(data)
			if err != nil {
				fmt.Println("Error parsing the configuration file:", err)
			}

			errs := configuration.Validate(k)
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

	validateCmd.Flags().StringVarP(&config, "config", "f", "kaarkala.yaml", "Input Configuration File")
}
