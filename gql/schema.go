package gql

import "github.com/graphql-go/graphql"

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var AuthorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

var TutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				Type: AuthorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(CommentType),
			},
		},
	},
)

var FieldsTutorial = graphql.Fields{
	"tutorial": &graphql.Field{
		Type:        TutorialType,
		Description: "Get tutorial by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: TutorialResolver,
	},
	"list": &graphql.Field{
		Type:        graphql.NewList(TutorialType),
		Description: "Get tutorial list",
		Resolve:     ListTutorialResolver,
	},
}

func GetSchemaTutorial() (graphql.Schema, error) {
	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: FieldsTutorial,
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	return graphql.NewSchema(schemaConfig)
}
