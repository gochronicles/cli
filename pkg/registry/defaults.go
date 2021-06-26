
package registry

import (
	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Autobot CLI.
const CLIVersion string = "1.0.0"
const EMPTY = ""
const FIBER_TEMPLATE string = "github.com/gochronicles/monorepo-fiber-neo4j"
const FIBER_TEMPLATE_CREATION_SUCCESS = "Project to be created with Fiber as a framework and Neo4j as Backend"
const NO_TEMPLATE_CHOSEN = "No Template was chosen"

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Backend       string
	Proxy         string
	AgreeCreation bool `survey:"agree"`
}

var (

	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"fiber",
					"fastApi",
					"n/a",
				},
				Default:  "fiber",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "Are you sure that we can proceed ? ;)",
				Default: true,
			},
		},
	}

	// CustomCreateQuestions survey's questions for `create -c` command.
	CustomCreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom backend repository:",
				Help:    "No need to specify `http://` or `https://` protocol.",
			},
			Validate: survey.Required,
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

)
