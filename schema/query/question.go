package query

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
	"github.com/togglhire/backend-homework/security"
)

var Questions = &graphql.Field{
	Type: graphql.NewList(output.Question),
	Args: graphql.FieldConfigArgument{
		"page": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"pageSize": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetAllQuestions(p.Context, p.Args)
	}),
}

var Question = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},

	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.GetQuestionByID(p.Context, p.Args)
	}),
}
