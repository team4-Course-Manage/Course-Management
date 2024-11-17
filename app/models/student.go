package models

type Student struct {
	StudentID string `gorm:"column:student_id;not null"`
	Name      string `gorm:"column:name;not null"`
	Institute string `gorm:"column:institute"`
	Password  string `gorm:"column:password;not null"`
}

//gorm默认复数表名，对应数据库，可以重写TableName方法
func (Student) TableName() string {
	return "student"
}
