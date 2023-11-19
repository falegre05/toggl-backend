package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
	"github.com/togglhire/backend-homework/security"
)

var AddOptionToQuestion = &graphql.Field{
	Type: output.Option,
	Args: graphql.FieldConfigArgument{
		"questionID": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"correct": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.AddOptionToQuestion(p.Context, p.Args)
	}),
}

var UpdateOption = &graphql.Field{
	Type: output.Option,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"correct": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.UpdateOption(p.Context, p.Args)
	}),
}

var DeleteOption = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: security.Check(security.PermissionsUser, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.DeleteOption(p.Context, p.Args)
	}),
}
