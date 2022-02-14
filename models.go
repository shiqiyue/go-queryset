package go_queryset

import (
	"github.com/lib/pq"
	"time"
)

//go:generate goqueryset -in $GOFILE
// gen:qs
// 标签
type Tag struct {
	ID             int64          `gorm:"primarykey; comment:ID"`
	CreatedAt      time.Time      `gorm:"not null; comment:创建时间"`
	UpdatedAt      time.Time      `gorm:"not null; comment:更新时间"`
	Name           string         `gorm:"not null; index; comment:名称"`
	Aliases        pq.StringArray `gorm:"type:text[]; index; comment:别名"`
	Disabled       bool           `gorm:"not null; index; comment:是否禁用"`
	DisabledReason *string        `gorm:"comment:禁用理由"`
	DisabledAt     *time.Time     `gorm:"comment:禁用时间"`
	DisabledBy     *int64         `gorm:"comment:禁用操作人"`
}

func (p *Tag) TableName() string {
	return "sec_tag"
}
