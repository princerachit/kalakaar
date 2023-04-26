package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kalakaar",
	Short: "The backend cli for kalaakar",
	Long: `Kalakaar serves multiple purposes like validating the kaarkala.yml file,
creating an xml tree for visual interpretation or even creating a jpeg image
representing the pipeline visually. It also serves as a clifor the kalaakar
backend server.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
