package registry

import (
	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Autobot CLI.
const CLIVersion string = "1.0.7"
const EMPTY = ""

const FIBER_NEO4J_TEMPLATE string = "github.com/gochronicles/monorepo-fiber-neo4j.git"
const FIBER_NEO4J_TEMPLATE_CREATION_SUCCESS = "Project to be created with Fiber as a framework backend and Neo4j as database"

const FASTAPI_POSTGRES_TEMPLATE string = "github.com/gochronicles/monorepo-fastapi-postgresql.git"
const FASTAPI_POSTGRES_TEMPLATE_SUCCESS = "Project to be created with FastAPI as a framework backend and PostgreSQL as database"

const FASTAPI_NEO4J_TEMPLATE string = "github.com/gochronicles/monorepo-fastapi-neo4j.git"
const FASTAPI_NEO4J_TEMPLATE_SUCCESS = "Project to be created with FastAPI as a framework backend and Neo4j as database"

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
					"fiber-neo4j",
					"fastapi-postgres",
					"fastapi-neo4j",
					"n/a",
				},
				Default:  "fiber-neo4j",
				PageSize: 4,
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
