package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"-u"},
	Short:   "Install the newest version of Project Wizard.",
	Long:    "\nInstall the newest version of Project Wizard.",
	RunE:    runUpdateCmd,
}

func runUpdateCmd(cmd *cobra.Command, args []string) error {

	exec.Command("go install github.com/NotReeceHarris/project-wizard/cmd/pwizard@latest").Output()
	return nil
}
