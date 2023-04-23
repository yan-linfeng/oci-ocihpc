// This software is licensed under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl.

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/oracle-quickstart/oci-ocihpc/stacks"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a stack for deployment",
	Long: `
Example command: ocihpc init --stack ClusterNetwork
	`,
	Run: func(cmd *cobra.Command, args []string) {
		stack, _ := cmd.Flags().GetString("stack")
		fmt.Printf("\nDownloading stack %s...", stack)
		stackInit(stack)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("stack", "s", "", "Name of the stack you want to deploy.")
	initCmd.MarkFlagRequired("stack")
}

func stackInit(stack string) {
	stackZipFile, err := stacks.ConfigFS.Open(fmt.Sprintf("%s/%s.zip", stack, stack))
	helpers.FatalIfError(err)
	defer stackZipFile.Close()

	zipfile := stack + ".zip"
	stackZipErr := copyFile(stackZipFile, zipfile)
	helpers.FatalIfError(stackZipErr)

	configFile, readConfigRrr := stacks.ConfigFS.Open(fmt.Sprintf("%s/config.zip", stack))
	helpers.FatalIfError(readConfigRrr)
	defer configFile.Close()

	configFilePath := filepath.Join(getWd(), "config.json")
	configFileErr := copyFile(configFile, configFilePath)
	helpers.FatalIfError(configFileErr)

	fmt.Println("\n\nDownloaded stack " + stack)
	fmt.Printf("\nIMPORTANT: Edit the contents of the %s file before running ocihpc deploy command\n\n", configFilePath)
}
