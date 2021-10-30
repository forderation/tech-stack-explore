package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/forderation/go-se-stack/gql"
	"github.com/graphql-go/graphql"
)

type TestCaseGQL struct {
	name  string
	query string
	want  string
}

func Test_GraphqlTutorial(t *testing.T) {
	getListCase := TestCaseGQL{
		name: "query list tutorial gql",
		query: `{
				list {
					id
					title
					comments {
						body
					}
					author {
						Name
						Tutorials
					}
				}
			}`,
		want: `{"data":{"list":[{"author":{"Name":"Elliot Forbes","Tutorials":[1]},"comments":[{"body":"First Comment"}],"id":1,"title":"Go GraphQL Tutorial"}]}}`,
	}
	getIdCase := TestCaseGQL{
		name: "query get id tutorial gql +",
		query: `
		{
			tutorial(id:1) {
				title
				author {
					Name
					Tutorials
				}
			}
		}
	`,
		want: `{"data":{"tutorial":{"author":{"Name":"Elliot Forbes","Tutorials":[1]},"title":"Go GraphQL Tutorial"}}}`,
	}
	getIdCaseNeg := TestCaseGQL{
		name: "query get id not in list gql -",
		query: `
		{
			tutorial(id:3) {
				title
				author {
					Name
					Tutorials
				}
			}
		}
		`,
		want: `{"data":{"tutorial":null}}`,
	}

	testCases := []TestCaseGQL{
		getListCase, getIdCase, getIdCaseNeg,
	}

	schema, err := gql.GetSchemaTutorial()
	if err != nil {
		t.Fatalf("failed to iniate schema , error: %v", err)
	}
	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			params := graphql.Params{Schema: schema, RequestString: c.query}
			r := graphql.Do(params)
			if len(r.Errors) > 0 {
				t.Fatalf("Failed to test gql errors: %+v", r.Errors)
			}
			rJSON, _ := json.Marshal(r)
			got := string(rJSON)
			if got != c.want {
				t.Errorf("got %v want %v", jsonPrettyString(got), jsonPrettyString(c.want))
			}
		})
	}
}

func jsonPrettyString(v interface{}) string {
	pretty, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatalf("error while pretty json %v", err)
	}
	return string(pretty)
}
