package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "protobuf"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

const (
	indexName = "students"
	_type     = "_doc"
)

var (
	client *elastic.Client
	ctx    context.Context
	err    error
)

//=============================//
// Tạo client
func NewClient() (*elastic.Client, error) {
	// khai báo một số option của client
	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),         // nếu không có dòng này thì mặc định là 127.0.0.1:9200
		elastic.SetHealthcheckInterval(5 * time.Second), // ngưng kết nối sau 5 giây
	}

	// tạo client với các option trên
	return elastic.NewClient(options...)
}

func InitElastic() {
	client, err = NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("client info:", client)
	ctx = context.Background()
}

// Tạo index
func CreateIndex(newStudent *pb.IndexStudentRequest) {
	dataJSON, err := json.Marshal(newStudent)
	js := string(dataJSON)
	_, err = client.Index().
		Index(indexName).
		BodyJson(js).
		Type(_type).
		Id(newStudent.Id).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("inserted:", newStudent.Name)
}

// Search
func Search(keyword string) []uint64 {
	var ids []uint64
	// Search with a term query
	matchQuery := elastic.NewMatchQuery("name", keyword)
	matchQuery2 := elastic.NewMatchQuery("age", keyword)
	generalQuery := elastic.NewBoolQuery().Should(matchQuery).Should(matchQuery2)

	searchResult, err := client.Search().
		Index("students").
		Query(generalQuery).
		From(0).Size(1000).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	// Lấy danh sách id ép về kiểu uint64
	for _, hit := range searchResult.Hits.Hits {
		id, err := strconv.Atoi(hit.Id)
		if err != nil {
			continue
		}

		ids = append(ids, uint64(id))
	}

	return ids
}

// // Kiểm tra index tồn tại chưa
// func IndexExists(index string) {
// 	isExist, err := client.IndexExists(index).Do(ctx)
// 	if err != nil {
// 		fmt.Println("exist (error):", err)
// 	} else {
// 		fmt.Println("exist:", isExist)
// 	}
// }

// // Lấy document từ index
// func GetDocument(index, id string) {
// 	res, err := client.Get().
// 		Index(index).
// 		Id(id).
// 		Type(_type).
// 		Do(ctx)

// 	if err != nil {
// 		fmt.Println("get: ", err)
// 	} else {
// 		if res.Found {
// 			fmt.Println("get res: ", res.Id, res.Version, res.Index, res.Type)
// 		} else {
// 			fmt.Println("not found")
// 		}
// 	}
// }

// // Xóa index
// func DeleteIndex(index string) {
// 	_, err := client.DeleteIndex(index).Do(ctx)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("deleted index: ", index)
// }

// // Lấy danh sách tên index
// func ListIndexNames() {
// 	names, _ := client.IndexNames()
// 	for _, name := range names {
// 		fmt.Println(name)
// 	}
// }

// // Xóa tất cả index
// func DeleteAllIndexes() {
// 	names, _ := client.IndexNames()
// 	for _, name := range names {
// 		DeleteIndex(name)
// 	}
// }
