package practice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)
func Index() {
	// https://github.com/elastic/go-elasticsearch/tree/main/_examples 列子
	es, err := elasticsearch.NewDefaultClient()
	if err!=nil{
		log.Fatalf(err.Error())
	}
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title": "张三李四？",
		"content": "外面的世界真的很精彩",
	}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf(err.Error())
	}
	res, err := es.Index("demo", &buf, es.Index.WithDocumentType("doc"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func Query()  {
	es, err := elasticsearch.NewDefaultClient()
	if err!=nil{
		log.Fatalf(err.Error())
	}
	// search - highlight
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "世界",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags" : []string{"<font color='red'>"},
			"post_tags" : []string{"</font>"},
			"fields" : map[string]interface{}{
				"title" : map[string]interface{}{},
			},
		},
	}
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf(err.Error())
	}
	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("demo"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func DeleteByQuery() {
	// new client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// DeleteByQuery deletes documents matching the provided query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
	}
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf(err.Error())
	}
	index := []string{"demo"}
	res, err := es.DeleteByQuery(index, &buf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func Delete() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Delete removes a document from the index
	res, err := es.Delete("demo", "POcKSHIBX-ZyL96-ywQO")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func Create() {
	addresses := []string{"http://127.0.0.1:9203"}
	config := elasticsearch.Config{
		Addresses:             addresses,
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Create creates a new document in the index.
	// Returns a 409 response when a document with a same ID already exists in the index.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"title": "你看到外面的世界是什么样的？",
		"content": "外面的世界真的很精彩",
	}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf(err.Error())
	}
	res, err := es.Create("demo", "esd02", &buf, es.Create.WithDocumentType("doc"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}

func Get() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses:             addresses,
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	res, err := es.Get("demo", "esd")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func Update() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses:             addresses,
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// Update updates a document with a script or partial document.
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"doc": map[string]interface{}{
			"title": "更新你看到外面的世界是什么样的？",
			"content": "更新外面的世界真的很精彩",
		},
	}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf(err.Error())
	}
	res, err := es.Update("demo", "esd", &buf, es.Update.WithDocumentType("doc"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
func UpdateByQuery() {
	addresses := []string{"http://127.0.0.1:9200", "http://127.0.0.1:9201"}
	config := elasticsearch.Config{
		Addresses:             addresses,
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
	}
	// new client
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// UpdateByQuery performs an update on every document in the index without changing the source,
	// for example to pick up a mapping change.
	index := []string{"demo"}
	var buf bytes.Buffer
	doc := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "外面",
			},
		},
		// 根据搜索条件更新title
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source['title']='更新你看到外面的世界是什么样的？'",
		   },
		*/
		// 根据搜索条件更新title、content
		/*
		   "script": map[string]interface{}{
		       "source": "ctx._source=params",
		       "params": map[string]interface{}{
		           "title": "外面的世界真的很精彩",
		           "content": "你看到外面的世界是什么样的？",
		       },
		       "lang": "painless",
		   },
		*/
		// 根据搜索条件更新title、content
		"script": map[string]interface{}{
			"source": "ctx._source.title=params.title;ctx._source.content=params.content;",
			"params": map[string]interface{}{
				"title": "看看外面的世界真的很精彩",
				"content": "他们和你看到外面的世界是什么样的？",
			},
			"lang": "painless",
		},
	}
	if err = json.NewEncoder(&buf).Encode(doc); err != nil {
		log.Fatalf(err.Error())
	}
	res, err := es.UpdateByQuery(
		index,
		es.UpdateByQuery.WithDocumentType("doc"),
		es.UpdateByQuery.WithBody(&buf),
		es.UpdateByQuery.WithContext(context.Background()),
		es.UpdateByQuery.WithPretty(),
	)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}