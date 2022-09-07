package cmd

import (
	"github.com/NotReeceHarris/project-wizard/v1/pkg/pwizard"
	"github.com/NotReeceHarris/project-wizard/v1/pkg/registry"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pwizard",
	Version: registry.CLIVersion,
	Short:   "🧙 Increase your workflow productivity with Project Wizard",
	Long: `
A fast and easy way to create new projects with a single command.
Create a new production-ready project with templates for your favorite,
Language, Framework, and Tech stack.`,
}

func init() {
	rootCmd.SetOut(pwizard.Stdout)
	rootCmd.SetErr(pwizard.Stderr)
}

func Execute() {
	_ = rootCmd.Execute()
}
