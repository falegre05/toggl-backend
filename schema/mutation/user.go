package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
	"github.com/togglhire/backend-homework/security"
)

var AddUser = &graphql.Field{
	Type: output.User,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: security.Check(security.PermissionsPublic, func(p graphql.ResolveParams) (interface{}, error) {
		return resolvers.AddUser(p.Args)
	}),
}
