package handler

import (
	"context"
	"fmt"
	"net/http"
	pb "source/mix/protobuf"
	"source/mix/server/entity"

	"github.com/gin-gonic/gin"
)

// Lấy danh sách students
func GetStudents(c *gin.Context) {
	// Kết nối đến RPC
	conn, _ := ConnectRPC()
	defer conn.Close()

	// Tạo client để gọi phương thức từ RPC
	client := pb.NewStudentClient(conn)
	ctx := context.Background()

	// for search
	studentRequest := pb.StudentRequest{
		Keyword: c.Query("keyword"),
	}

	// Gọi hàm tìm kiếm student từ RPC
	studentResponse, err := client.SearchStudent(ctx, &studentRequest)

	// Lấy danh sách student trong database theo các id nhận được
	listStudents := DB.ListStudents{}
	err = listStudents.GetAllByIDs(studentResponse.Ids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": listStudents,
	})
}

// Tạo 1 dòng dữ liệu students
func CreateStudent(c *gin.Context) {
	student := DB.Student{}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := student.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	// Kết nối đến RPC
	conn, _ := ConnectRPC()
	defer conn.Close()

	// for search
	indexRequest := pb.IndexStudentRequest{
		Name: student.Name,
		Age:  (int32)(student.Age),
		Id:   fmt.Sprint(student.ID),
	}

	// Tạo client để gọi phương thức từ RPC
	client := pb.NewStudentClient(conn)
	ctx := context.Background()

	// Gọi hàm tìm kiếm student từ RPC
	indexResponse, err := client.IndexStudent(ctx, &indexRequest)
	fmt.Println(indexResponse)

	// Tạo thành công
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": student,
	})
}
