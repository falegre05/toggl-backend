package mutation

import "github.com/graphql-go/graphql"

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addUser": AddUser,

		"addQuestion":    AddQuestion,
		"updateQuestion": UpdateQuestion,
		"deleteQuestion": DeleteQuestion,
	},
})
