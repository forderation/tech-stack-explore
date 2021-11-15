package eslab

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/forderation/es-2/data"
	"github.com/olivere/elastic/v7"
)

func (es *EsInstance) GetESClient() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("ES initialized successfully")
	es.Instance = client
}

func (es *EsInstance) AddDataToIndex(payload interface{}) {
	ctx := context.Background()
	dataJSON, err := json.Marshal(payload)
	newStudentPayload := string(dataJSON)
	if err != nil {
		panic(err)
	}
	_, err = es.Instance.Index().Index("students").BodyJson(newStudentPayload).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("[Elastic][Insert Product] Insertion Successful")
}

func (es *EsInstance) GetStudentByName(name string) {
	ctx := context.Background()
	var students []data.Student

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name", name))

	queryStr, err1 := searchSource.Source()
	queryJson, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}

	// print query block
	fmt.Println("[esclient]Final ESQuery=\n", string(queryJson))

	searchService := es.Instance.Search().Index("students").SearchSource(searchSource)
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	for _, hit := range searchResult.Hits.Hits {
		var student data.Student
		err := json.Unmarshal(hit.Source, &student)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}
		students = append(students, student)
	}

	for _, s := range students {
		fmt.Printf("Student found Name: %s, Age: %d, Score: %f \n", s.Name, s.Age, s.AverageScore)
	}
}
