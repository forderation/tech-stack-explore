package main

import (
	"encoding/json"
	"testing"

	"github.com/graphql-go/graphql"
)

func Test_GraphqlHello(t *testing.T) {
	testCases := []struct {
		name     string
		query    string
		expected string
	}{
		{
			name:     "hello graphql positive case",
			query:    `{hello}`,
			expected: `{"data":{"hello":"world"}}`,
		},
	}

	schema, err := GetSchemaHello()
	if err != nil {
		t.Fatalf("failed to iniate schema , error: %v", err)
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			params := graphql.Params{Schema: schema, RequestString: c.query}
			r := graphql.Do(params)
			if len(r.Errors) > 0 {
				t.Fatalf("failed to execute graphql operation errors: %+v", r.Errors)
			}
			rJSON, _ := json.Marshal(r)
			got := string(rJSON)
			if got != c.expected {
				t.Errorf("got %v want %v", got, c.expected)
			}
		})
	}
}
