package go_queryset

import (
	"gorm.io/gorm"
	"time"
)

//go:generate goqueryset -in $GOFILE
// gen:qs
// Contact 联系人
type Contact struct {
	ID        int64          `gorm:"primarykey; comment:ID"`
	CreatedAt time.Time      `gorm:"not null; comment:创建时间"`
	UpdatedAt time.Time      `gorm:"not null; comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index; comment:删除时间"`
	Name      string         `gorm:"not null; size:255; comment:名称"`
	Email     string         `gorm:"not null; size:255; comment:邮箱"`
	Phone     string         `gorm:"not null; size: 255; comment:手机号码"`
}

func (p *Contact) TableName() string {
	return "nc_contact"
}
