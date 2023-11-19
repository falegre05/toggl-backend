package query

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
)

var Question = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: resolvers.GetQuestionByID,
}

var Questions = &graphql.Field{
	Type: graphql.NewList(output.Question),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetAllQuestions()
	},
}
