package go_queryset

import (
	"time"
)

//go:generate goqueryset -in $GOFILE
// gen:qs
// 素材标签使用统计
type IdeaTagStat struct {
	Id          int64     `gorm:"primaryKey;comment:主键"`
	CreatedAt   time.Time `gorm:"not null;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"not null;comment:更新时间"`
	TagId       int64     `gorm:"not null;comment:标签ID"`
	UseCount    int64     `gorm:"not null; comment:使用次数"`
	SearchCount int64     `gorm:"not null; comment:搜索次数"`
}

func (s *IdeaTagStat) TableName() string {
	return "irm_idea_tag_stat"
}
