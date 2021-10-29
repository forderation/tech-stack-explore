package main

import (
	"github.com/graphql-go/graphql"
)

func HelloResolver(p graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}

func GetSchemaHello() (graphql.Schema, error) {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type:    graphql.String,
			Resolve: HelloResolver,
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}

func main() {

}
