package handler

import (
	"context"
	"encoding/json"
	"fmt"
	pb "rpc"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func failOnError(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
		return true
	}
	return false
}

//-----------------------------------

type StudentHandler struct {
	rabbitConn  *amqp.Connection
	elasticConn *grpc.ClientConn
	localConn   *grpc.ClientConn
}

func NewStudentHandler(rabbit *amqp.Connection, elasticRPC *grpc.ClientConn, localRPC *grpc.ClientConn) *StudentHandler {
	return &StudentHandler{
		rabbitConn:  rabbit,
		elasticConn: elasticRPC,
		localConn:   localRPC,
	}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	client := pb.NewExcelClient(h.localConn)
	client2 := pb.NewStudentClient(h.elasticConn)
	ctx := context.Background()

	studentRequest := pb.StudentRequest{
		Keyword: c.Query("keyword"),
	}

	studentResponse, err := client2.SearchStudent(ctx, &studentRequest)
	if failOnError(c, err) {
		return
	}

	list, err := client.GetAllById(ctx, studentResponse)
	if failOnError(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": list.Students,
	})
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	client := pb.NewExcelClient(h.localConn)
	client2 := pb.NewStudentClient(h.elasticConn)
	ctx := context.Background()

	student := pb.StudentStruct{}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	data, err := client.Create(ctx, &student)
	if failOnError(c, err) {
		return
	}

	indexRequest := pb.IndexStudentRequest{
		Name: data.Name,
		Age:  (int32)(data.Age),
		Id:   fmt.Sprint(data.Id),
	}

	// Gọi hàm tìm kiếm student từ RPC
	indexResponse, err := client2.IndexStudent(ctx, &indexRequest)
	fmt.Println(indexResponse)

	c.JSON(200, gin.H{
		"code": 200,
		"data": student,
	})
}

func (h *StudentHandler) ExportXLSX(c *gin.Context) {
	path := strings.TrimSpace(c.Query("path"))
	fileName := strings.TrimSpace(c.Query("file-name"))
	fileName += ".xlsx"

	if fileName == ".xlsx" {
		now := time.Now()
		timeStamp := now.UnixNano()
		fileName = "student_" + fmt.Sprint(timeStamp) + ".xlsx"
	}

	ch, err := h.rabbitConn.Channel()
	if failOnError(c, err) {
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	if failOnError(c, err) {
		return
	}

	r := pb.XlsxRequest{
		Path:     path,
		FileName: fileName,
	}

	body, err := json.Marshal(r)
	if failOnError(c, err) {
		return
	}

	err = ch.Publish("logs", "", false, false, amqp.Publishing{ContentType: "text/plain", Body: body})
	if failOnError(c, err) {
		return
	}
	// ----------------------------------------------

	c.JSON(200, gin.H{
		"code":    200,
		"message": "File saved at: " + r.Path + "/" + r.FileName,
	})
}
