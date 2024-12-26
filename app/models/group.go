package models

type Group struct {
	GroupId           int    `gorm:"column:group_id;primaryKey;autoIncrement"`
	GroupName         string `gorm:"column:group_name"`
	MaxMembers        int    `gorm:"column:max_members"`
	CurrentMembers    int    `gorm:"column:current_members;default:1"`
	GroupMemberNumber int    `gorm:"column:current_members"`
}
