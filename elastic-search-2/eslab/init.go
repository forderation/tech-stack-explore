package eslab

import "github.com/olivere/elastic/v7"

type EsInstance struct {
	Instance *elastic.Client
}

type EsClientRepo interface {
	GetESClient()
	AddDataToIndex(payload interface{})
	GetStudentByName(string)
}
