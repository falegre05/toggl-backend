package mutation

import "github.com/graphql-go/graphql"

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// users
		"addUser": AddUser,

		// questions
		"addQuestion":    AddQuestion,
		"updateQuestion": UpdateQuestion,
		"deleteQuestion": DeleteQuestion,

		// options
		"addOptionToQuestion": AddOptionToQuestion,
		"updateOption":        UpdateOption,
		"deleteOption":        DeleteOption,
	},
})
