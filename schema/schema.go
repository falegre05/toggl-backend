package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/togglhire/backend-homework/schema/mutation"
	"github.com/togglhire/backend-homework/schema/query"
)

func Get() (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    query.Query,
		Mutation: mutation.Mutation,
	})
}
