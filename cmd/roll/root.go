package roll

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gochronicles/pkg/registry"
)

var (
	backend,backendMessage                            string                 // define project variables
	options                            []string               // define options, proxy list
	useCustomTemplate                  bool                   // define custom templates
	createAnswers, customCreateAnswers registry.CreateAnswers // define answers variable for `create` command

	surveyIconsConfig = func(icons *survey.IconSet) {
		icons.Question.Format = "cyan"
		icons.Question.Text = "Q:"
		icons.Help.Format = "blue"
		icons.Help.Text = "Help ->"
		icons.Error.Format = "red"
		icons.Error.Text = "Note ->"
	}
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     "autobots",
	Version: registry.CLIVersion,
	Short:   "A powerful CLI To Create Microservices Monorepos based on domain",
	Long: `
A powerful CLI To Create Microservices Monorepos based on domain.

Create a new production-ready project with backend in Go/Python`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}
