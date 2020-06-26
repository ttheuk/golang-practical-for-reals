package student

import (
	"entity"
	pb "rpc"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/jinzhu/gorm"
)

type ExcelRepository struct {
	db *gorm.DB
}

func NewExcelRepository() *ExcelRepository {
	return &ExcelRepository{}
}

func (repo *ExcelRepository) ExportXLSX(r *pb.XlsxRequest) error {
	client := pb.NewStudentClient(h.conn)
	ctx := context.Background()
	
	// Thay bằng code lấy data
	s1 := entity.Student{Name: "the", Age: 20}
	s2 := entity.Student{Name: "thanh", Age: 22}
	s3 := entity.Student{Name: "nguyen", Age: 21}

	list := entity.ListStudent{s1, s2, s3}

	f := excelize.NewFile()
	// Set value for fist row.
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "Age")

	for i, obj := range list {
		index := strconv.Itoa(i + 2)
		// f.SetCellValue("Sheet1", "A"+index, obj.ID)
		f.SetCellValue("Sheet1", "B"+index, obj.Name)
		f.SetCellValue("Sheet1", "C"+index, obj.Age)
	}

	// Set active sheet
	f.SetActiveSheet(0)

	// Save xlsx file by the given path
	if r.Path != "" {
		r.Path = r.Path + "/"
	}

	actualPath := r.Path + r.FileName
	err := f.SaveAs(actualPath)
	return err
}
