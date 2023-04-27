// This software is licensed under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl

package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/oracle-quickstart/oci-ocihpc/stacks"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available stacks",
	Long: `
Example command: ocihpc list
	`,

	Run: func(cmd *cobra.Command, args []string) {
		localStackCatalogs := ""

		localStackConfigPath, _ := cmd.Flags().GetString("f")

		if localStackConfigPath != "" {
			localStackConfigFile, err := os.Open(localStackConfigPath)
			helpers.FatalIfError(err)

			defer localStackConfigFile.Close()

			var localStackConfig map[string]any
			if err := json.NewDecoder(localStackConfigFile).Decode(&localStackConfig); err != nil {
				log.Fatal(err)
			}
			for key := range localStackConfig {
				localStackCatalogs += key + "\r\n"
			}
		}

		defaultCatalogs, err := stacks.ConfigFS.ReadFile("catalog")
		helpers.FatalIfError(err)

		fmt.Printf("\nList of available stacks:\n\n")
		fmt.Println(localStackCatalogs + string(defaultCatalogs))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("f", "f", "", "Local stack configuration file.")
}
