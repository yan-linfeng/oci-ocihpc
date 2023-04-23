// This software is licensed under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl

package cmd

import (
	"fmt"

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
		catalog, err := stacks.ConfigFS.ReadFile("catalog")
		helpers.FatalIfError(err)

		fmt.Printf("\nList of available stacks:\n\n")
		fmt.Println(string(catalog))
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
