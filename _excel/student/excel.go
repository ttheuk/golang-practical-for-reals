package student

import (
	"context"
	pb "rpc"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

type ExcelRepository struct {
	db *gorm.DB
}

func NewExcelRepository() *ExcelRepository {
	return &ExcelRepository{}
}

func (repo *ExcelRepository) ExportXLSX(r *pb.XlsxRequest, conn *grpc.ClientConn) error {
	client := pb.NewExcelClient(conn)
	ctx := context.Background()

	// Lấy data từ server chính
	list, err := client.GetAllStudent(ctx, &pb.Empty{})

	// Tạo file excel
	f := excelize.NewFile()

	// Set value for fist row.
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Age")
	f.SetCellValue("Sheet1", "D1", "Created at")
	f.SetCellValue("Sheet1", "E1", "Updated at")

	for i, obj := range list.Students {

		index := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", "A"+index, obj.Id)
		f.SetCellValue("Sheet1", "B"+index, obj.Name)
		f.SetCellValue("Sheet1", "C"+index, obj.Age)
		f.SetCellValue("Sheet1", "D"+index, (time.Unix(0, obj.CreatedAt).Format("02/01/2006, 15:04:05")))
		f.SetCellValue("Sheet1", "E"+index, (time.Unix(0, obj.UpdatedAt).Format("02/01/2006, 15:04:05")))
	}

	// Set active sheet
	f.SetActiveSheet(0)

	// Save xlsx file by the given path
	if r.Path != "" {
		r.Path = r.Path + "/"
	}

	actualPath := r.Path + r.FileName
	err = f.SaveAs(actualPath)
	return err
}
