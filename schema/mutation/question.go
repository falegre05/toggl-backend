package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/resolvers"
	"github.com/togglhire/backend-homework/schema/output"
)

var AddQuestion = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.AddQuestion,
}

var UpdateQuestion = &graphql.Field{
	Type: output.Question,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"body": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.UpdateQuestion,
}

var DeleteQuestion = &graphql.Field{
	Type: graphql.Boolean,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: resolvers.DeleteQuestion,
}
