package cmd

import (
	"os/exec"

	"github.com/NotReeceHarris/project-wizard/pkg/registry"
	"github.com/sanbornm/go-selfupdate/selfupdate"
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

	var updater = &selfupdate.Updater{
		CurrentVersion: registry.CLIVersion,
		CmdName:        "pwizard", // app name
	}

	if updater != nil {
		go updater.BackgroundRun()
	}

	exec.Command("go install github.com/NotReeceHarris/project-wizard/cmd/pwizard@latest").Output()
	return nil
}
