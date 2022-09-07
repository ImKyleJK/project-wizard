package cmd

import (
	"github.com/NotReeceHarris/project-wizard/pkg/pwizard"
	"github.com/spf13/cobra"
)

var (
	backend, frontend, proxy string                 // define project variables
	inventory, playbook      map[string]interface{} // define template variables
	options, proxyList       []string               // define options, proxy list
	useCustomTemplate        bool                   // define custom templates
	askBecomePass            bool                   // install Ansible roles, ask become pass
)

var rootCmd = &cobra.Command{
	Use:     "pwizard",
	Version: "1.0.0",
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
