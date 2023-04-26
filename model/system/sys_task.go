package system

import (
	"server/global"
	"time"
)

type SysTask struct {
	global.GVA_MODEL
	Name        string        `json:"name" gorm:"comment:作业名称"`        // 作业名称
	Description string        `json:"description" gorm:"comment:作业描述"` // 作业描述
	StartTime   time.Time     `json:"start_time" gorm:"comment:开始时间"`  // 开始时间
	EndTime     time.Time     `json:"end_time" gorm:"comment:截止时间"`    // 截止时间
	Score       uint          `json:"score" gorm:"comment:分数"`         // 分数
	Practice    []SysPractice `json:"practices" gorm:"foreignKey:TaskID"`
}

func (SysTask) TableName() string {
	return "sys_tasks"
}
