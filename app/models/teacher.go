package models

type Teacher struct {
	TeacherID string `gorm:"column:teacher_ID;not null"`
	Name      string `gorm:"column:name;not null"`
	Institute string `gorm:"column:institute"`
	Password  string `gorm:"column:password;not null"`
}

func (Teacher) TableName() string {
	return "teacher"
}
