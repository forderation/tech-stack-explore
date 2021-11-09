package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

var EsClient *elasticsearch.Client

func ReadText(reader *bufio.Scanner, prompt string) string {
	fmt.Print(prompt + ": ")
	reader.Scan()
	return reader.Text()
}

func Exit() {
	fmt.Println("Goodbye!")
	os.Exit(0)
}

const BASE_API = "http://stapi.co/api/v1/rest/spacecraft"

func LoadData() {
	var spaceCrafts []map[string]interface{}
	pageNumber := 0
	for {
		response, err := http.Get(BASE_API + "/search?pageSize=100&pageNumber=" + strconv.Itoa(pageNumber))
		if err != nil {
			fmt.Printf("error while request: %s\n", err.Error())
			break
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		page := result["page"].(map[string]interface{})
		totalPages := int(page["totalPages"].(float64))
		crafts := result["spacecrafts"].([]interface{})

		for _, craftInterface := range crafts {
			craft := craftInterface.(map[string]interface{})
			spaceCrafts = append(spaceCrafts, craft)
		}

		pageNumber++
		if pageNumber >= totalPages {
			break
		}
	}

	for _, data := range spaceCrafts {
		uid, _ := data["uid"].(string)
		jsonString, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("error while parsing data spacecraft %s", err.Error())
			break
		}
		request := esapi.IndexRequest{
			Index:      "stsc",
			DocumentID: uid,
			Body:       strings.NewReader(string(jsonString)),
		}
		_, err = request.Do(context.Background(), EsClient)
		if err != nil {
			fmt.Println("error while indexing result request: ", err.Error())
		}
	}

	print(len(spaceCrafts), " spacecraft read \n")
}

func PrintSpaceCraft(spaceCraft map[string]interface{}) {
	name := spaceCraft["name"]
	status := ""
	if spaceCraft["status"] != nil {
		status = "- " + spaceCraft["status"].(string)
	}
	registry := ""
	if spaceCraft["registry"] != nil {
		registry = "- " + spaceCraft["registry"].(string)
	}
	class := ""
	if spaceCraft["spacecraftClass"] != nil {
		class = "- " + spaceCraft["spacecraftClass"].(map[string]interface{})["name"].(string)
	}
	fmt.Println(name, registry, class, status)
}

func Search(reader *bufio.Scanner, queryType string) {
	key := ReadText(reader, "Enter key")
	value := ReadText(reader, "Enter value")
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			queryType: map[string]interface{}{
				key: value,
			},
		},
	}
	json.NewEncoder(&buffer).Encode(query)
	response, err := EsClient.Search(EsClient.Search.WithIndex("stsc"), EsClient.Search.WithBody(&buffer))
	if err != nil {
		fmt.Println("error while search", err.Error())
	}
	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	log.Println(result)
	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		craft := hit.(map[string]interface{})["_source"].(map[string]interface{})
		PrintSpaceCraft(craft)
	}
}

func Get(reader *bufio.Scanner) {
	id := ReadText(reader, "Enter spacecraft ID")
	request := esapi.GetRequest{Index: "stsc", DocumentID: id}
	response, err := request.Do(context.Background(), EsClient)
	if err != nil {
		fmt.Println("error while read index: ", err.Error())
	}
	var results map[string]interface{}
	json.NewDecoder(response.Body).Decode(&results)
	if results["_source"] == nil {
		fmt.Printf("search for %s not found \n", id)
		return
	}
	PrintSpaceCraft(results["_source"].(map[string]interface{}))
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var err error
	EsClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load spacecraft")
		fmt.Println("2) Get spacecraft")
		fmt.Println("3) Search spacecraft by key and value")
		fmt.Println("4) Search spacecraft by key and prefix")
		option := ReadText(reader, "Enter option")
		switch option {
		case "0":
			Exit()
		case "1":
			LoadData()
		case "2":
			Get(reader)
		case "3":
			Search(reader, "match")
		case "4":
			Search(reader, "prefix")
		default:
			fmt.Println("Invalid option")
		}
	}
}
