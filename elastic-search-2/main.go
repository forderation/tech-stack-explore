package main

import (
	"github.com/forderation/es-2/data"
	"github.com/forderation/es-2/eslab"
)

func main() {
	es := eslab.EsInstance{}
	ExecuteEs(&es)
}

func ExecuteEs(esRepo eslab.EsClientRepo) {
	// get connection
	esRepo.GetESClient()

	// add data
	newStudent := data.Student{
		Name:         "Gopher Due",
		Age:          10,
		AverageScore: 99.9,
	}
	esRepo.AddDataToIndex(newStudent)

	//querying data
	esRepo.GetStudentByName("Gopher Due")
}
