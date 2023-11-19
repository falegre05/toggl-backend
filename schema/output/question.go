package output

import (
	"github.com/graphql-go/graphql"
)

var Option = graphql.NewObject(graphql.ObjectConfig{
	Name: "Option",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"correct": &graphql.Field{
			Type: graphql.Boolean,
		},
		"questionId": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var Question = graphql.NewObject(graphql.ObjectConfig{
	Name: "Question",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"body": &graphql.Field{
			Type: graphql.String,
		},
		"options": &graphql.Field{
			Type: graphql.NewList(Option),
		},
		"userID": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
