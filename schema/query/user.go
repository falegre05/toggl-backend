package query

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/security"
)

var UserLogin = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: security.Check(security.PermissionsPublic, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.UserLogin(p.Args)
	}),
}
