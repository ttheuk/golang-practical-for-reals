package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	pb "rpc"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

const (
	host      = "http://localhost:9200"
	indexName = "students"
	_type     = "_doc"
)

var (
	client *elastic.Client
	ctx    context.Context
	err    error
)

// Tạo client
func NewClient() (*elastic.Client, error) {
	// khai báo một số option của client
	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(true),
		elastic.SetURL(host), // nếu không có dòng này thì mặc định là 127.0.0.1:9200
		elastic.SetHealthcheckInterval(5 * time.Second),
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
	exist, err := client.IndexExists(indexName).Do(ctx)
	if !exist || err != nil {
		return []uint64{}
	}

	var ids []uint64
	// Search theo trường name
	matchQuery := elastic.NewMatchQuery("name", keyword)
	generalQuery := elastic.NewBoolQuery().Should(matchQuery)

	// nếu keyword là số thì check thêm trường age
	if id, err := strconv.Atoi(keyword); err == nil {
		matchQuery2 := elastic.NewMatchQuery("age", id)
		generalQuery.Should(matchQuery2)
	}

	res, err := client.Search().
		Index(indexName).
		Query(generalQuery).
		From(0).Size(10000).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}

	// Lấy danh sách id ép kiểu về uint64
	for _, hit := range res.Hits.Hits {
		id, err := strconv.Atoi(hit.Id)
		if err != nil {
			log.Print("error id:", id)
			continue
		}

		ids = append(ids, uint64(id))
	}

	return ids
}

// Export excel
// func ExportXLSX(info *pb.XlsxRequest) (*pb.XlsxResponse, error) {
// 	// Thay bằng code lấy data
// 	s1 := entity.Student{Name: "the", Age: 20}
// 	s2 := entity.Student{Name: "thanh", Age: 22}
// 	s3 := entity.Student{Name: "nguyen", Age: 21}

// 	list := entity.ListStudent{s1, s2, s3}

// 	f := excelize.NewFile()
// 	// Create a new sheet.
// 	// index := f.NewSheet("Sheet2")

// 	// Set value for fist row.
// 	f.SetCellValue("Sheet1", "A1", "ID")
// 	f.SetCellValue("Sheet1", "B1", "Name")
// 	f.SetCellValue("Sheet1", "C1", "Age")

// 	for i, obj := range list {
// 		index := strconv.Itoa(i + 2)
// 		// f.SetCellValue("Sheet1", "A"+index, obj.ID)
// 		f.SetCellValue("Sheet1", "B"+index, obj.Name)
// 		f.SetCellValue("Sheet1", "C"+index, obj.Age)
// 	}

// 	// Set active sheet
// 	f.SetActiveSheet(0)

// 	// Save xlsx file by the given path.
// 	if info.Path != "" {
// 		info.Path = info.Path + "/"
// 	}

// 	actualPath := info.Path + info.FileName
// 	if err := f.SaveAs(actualPath); err != nil {
// 		return &pb.XlsxResponse{
// 				Path:    actualPath,
// 				Message: "fail",
// 			},
// 			err
// 	}

// 	return &pb.XlsxResponse{
// 			Path:    actualPath,
// 			Message: "successful",
// 		},
// 		nil
// }
