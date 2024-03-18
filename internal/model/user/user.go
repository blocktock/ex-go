package user

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	Id          uint64                `gorm:"column:id" json:"-"`
	EnShortName string                `gorm:"column:en_short_name" json:"en_short_name"`
	CnName      string                `gorm:"column:cn_name" json:"cn_name"`
	Ctime       time.Time             `gorm:"column:ctime;autoCreateTime" json:"-"`
	MTime       time.Time             `gorm:"column:mtime;autoUpdateTime" json:"-"`
	Deleted     soft_delete.DeletedAt `gorm:"column:deleted;softDelete:flag" json:"-"`
}
