package gql

import "github.com/graphql-go/graphql"

var tutorials = PopulateTutorial()

func TutorialResolver(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if ok {
		for _, tutorial := range tutorials {
			if int(tutorial.ID) == id {
				return tutorial, nil
			}
		}
	}
	return nil, nil
}

func ListTutorialResolver(p graphql.ResolveParams) (interface{}, error) {
	return tutorials, nil
}
