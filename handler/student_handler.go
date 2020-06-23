package handler

import (
	"context"
	"encoding/json"
	"entity"
	"fmt"
	"repository/student"
	pb "rpc"

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
	service student.Service
	conn    *grpc.ClientConn
}

type XlsxRequest struct {
	Type     string
	Path     string
	FileName string
	Data     string
}

func NewStudentHandler(repo student.StudentRepository, rpc *grpc.ClientConn) *StudentHandler {
	return &StudentHandler{
		service: *student.NewService(repo),
		conn:    rpc,
	}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	// Tạo client để gọi phương thức từ RPC
	client := pb.NewStudentClient(h.conn)
	ctx := context.Background()

	// for search
	studentRequest := pb.StudentRequest{
		Keyword: c.Query("keyword"),
	}

	// Gọi hàm tìm kiếm student từ RPC
	studentResponse, err := client.SearchStudent(ctx, &studentRequest)

	// Lấy danh sách student trong database theo các id nhận được
	data, err := h.service.GetAllById(studentResponse.Ids)

	if failOnError(c, err) {
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	student := entity.Student{}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	err := h.service.Create(&student)

	if failOnError(c, err) {
		return
	}

	indexRequest := pb.IndexStudentRequest{
		Name: student.Name,
		Age:  (int32)(student.Age),
		Id:   fmt.Sprint(student.ID),
	}

	// Tạo client để gọi phương thức từ RPC
	client := pb.NewStudentClient(h.conn)
	ctx := context.Background()

	// Gọi hàm tìm kiếm student từ RPC
	indexResponse, err := client.IndexStudent(ctx, &indexRequest)
	fmt.Println(indexResponse)

	c.JSON(200, gin.H{
		"code": 200,
		"data": student,
	})
}

func (h *StudentHandler) ExportXLSX(c *gin.Context) {
	path := c.Query("path")
	fileName := c.Query("file-name")

	if fileName == "" {
		fileName = "demo_project_file.xlsx"
	}

	//------------------------------------------
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if failOnError(c, err) {
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if failOnError(c, err) {
		return
	}

	defer ch.Close()

	err = ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)

	if failOnError(c, err) {
		return
	}

	body := XlsxRequest{
		Type:     "export_xlsx",
		Path:     path,
		FileName: fileName,
	}
	jsonBody, err := json.Marshal(body)

	if failOnError(c, err) {
		return
	}

	err = ch.Publish("logs", "", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonBody,
		})

	if failOnError(c, err) {
		return
	}
	// ----------------------------------------------

	c.JSON(200, gin.H{
		"code": 200,
		"data": "sent request successful",
	})
}
