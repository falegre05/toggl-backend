package query

import "github.com/graphql-go/graphql"

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"userLogin": UserLogin,

		"question":  Question,
		"questions": Questions,
	},
})
