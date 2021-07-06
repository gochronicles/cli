package roll

import (
	"fmt"
	"strings"
	"time"

	"gochronicles/pkg/autobot"
	"gochronicles/pkg/registry"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rollCmd)
	rollCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend and frontend templates",
	)
}

// rollCmd represents the `create project` command.
var rollCmd = &cobra.Command{
	Use:     "roll",
	Aliases: []string{"new"},
	Short:   "Create a new Microservices Mono Repo via Termial UI",
	Long:    "\nCreate a new Microservices Mono Repo via Termial UI",
	RunE:    runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
func runCreateCmd(cmd *cobra.Command, args []string) error {
	// Start message.
	autobot.ShowMessage(
		"",
		fmt.Sprintf(
			"Create a new project via Autobot CLI v%v...",
			registry.CLIVersion,
		),
		true, true,
	)

	// Start survey.
	if useCustomTemplate {
		// Custom survey.
		if err := survey.Ask(
			registry.CustomCreateQuestions,
			&customCreateAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return autobot.ShowError(err.Error())
		}

		// Define variables for better display.
		backend = customCreateAnswers.Backend
	} else {
		// Default survey.
		if err := survey.Ask(
			registry.CreateQuestions,
			&createAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return autobot.ShowError(err.Error())
		}

		// Define variables for better display.

		if strings.Contains(createAnswers.Backend, "fiber-neo4j") {
			backend = registry.FIBER_NEO4J_TEMPLATE
		}

		if strings.Contains(createAnswers.Backend, "fastapi-postgres") {
			backend = registry.FASTAPI_POSTGRES_TEMPLATE
		}
		if strings.Contains(createAnswers.Backend, "fastapi-neo4j") {
			backend = registry.FASTAPI_NEO4J_TEMPLATE
		}
		switch createAnswers.Backend {

		case "fiber-neo4j":
			backend = registry.FIBER_NEO4J_TEMPLATE
			backendMessage = registry.FIBER_NEO4J_TEMPLATE_CREATION_SUCCESS
		case "fastapi-postgres":
			backend = registry.FASTAPI_POSTGRES_TEMPLATE
			backendMessage = registry.FASTAPI_POSTGRES_TEMPLATE_SUCCESS
		case "fastapi-neo4j":
			backend = registry.FASTAPI_NEO4J_TEMPLATE
			backendMessage = registry.FASTAPI_NEO4J_TEMPLATE_SUCCESS

		case "n/a":
			backend = registry.EMPTY
			backendMessage = registry.NO_TEMPLATE_CHOSEN
		}
	}
	autobot.ShowMessage(
		"info",
		fmt.Sprintf(backendMessage, backend),
		true, false,
	)
	// Start timer.
	startTimer := time.Now()

	/*
		The project's backend part creation.
	*/

	// Clone backend files from git repository.
	if err := autobot.GitClone("backend", backend); err != nil {
		autobot.ShowMessage(
			"info",
			fmt.Sprintf("Project Template was not chosen !", backend),
			true, false,
		)
	}

	// Show success report.
	autobot.ShowMessage(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backend),
		true, false,
	)

	// Delete unused backend files.
	autobot.RemoveFolders("backend", []string{".git", ".github"})

	// Stop timer.
	stopTimer := autobot.CalculateDurationTime(startTimer)
	autobot.ShowMessage(
		"info",
		fmt.Sprintf("Completed in %v seconds!", stopTimer),
		true, true,
	)

	autobot.ShowMessage(
		"",
		"* Please go through the Templates readme file to understand how monorepos work",
		false, true,
	)
	autobot.ShowMessage(
		"",
		"Your Microservices MonoRepo is created !!",
		false, true,
	)

	return nil
}
