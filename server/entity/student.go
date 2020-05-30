package DB

type Student struct {
	MyModel
	Name string `binding:"required" json:"name,omitempty" gorm:"TEXT;not null"`
	Age  int8   `binding:"required" json:"age,omitempty" gorm:"INT;not null"`
}
type ListStudents []Student

func (list *ListStudents) GetAllByIDs(ids []uint64) error {
	db, err := ConnectDB()
	defer db.Close()

	if err != nil {
		return err
	}

	return db.Where("id IN(?)", ids).Find(list).Error
	// return nil
}

func (this *Student) Create() error {
	db, err := ConnectDB()
	defer db.Close()

	if err != nil {
		return err
	}

	return db.Create(this).Error
}
