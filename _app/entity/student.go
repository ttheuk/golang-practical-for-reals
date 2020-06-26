package entity

type Student struct {
	MyModel
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
type ListStudent []Student
