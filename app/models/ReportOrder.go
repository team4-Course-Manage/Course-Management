package models

type ReportOrder struct {
	ID       uint     `gorm:"primaryKey"`
	Week     uint     `gorm:"column:week;not null"`
	Order    uint     `gorm:"column:order;not null"` // 顺序编号 1-4
	Student  *Student `gorm:"foreignKey:StudentID"`
	StudentID *string `gorm:"column:student_id"` // 可为空，表示未选择
}

func (ReportOrder) TableName() string {
	return "report_order"
}