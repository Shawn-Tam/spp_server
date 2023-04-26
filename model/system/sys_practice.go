package system

import (
	"server/global"
	"time"
)

type SysPractice struct {
	global.GVA_MODEL
	Name        string    `json:"name" gorm:"comment:作业名称"`        // 实验名称
	Description string    `json:"description" gorm:"comment:作业描述"` // 实验描述
	TaskID      uint      `json:"task_id" gorm:"comment:作业ID"`     // 作业ID
	Task        SysTask   `json:"task" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartTime   time.Time `json:"start_time" gorm:"comment:开始时间"` // 开始时间
	EndTime     time.Time `json:"end_time" gorm:"comment:截止时间"`   // 截止时间
	Score       uint      `json:"score" gorm:"comment:分数"`        // 分数
}

func (SysPractice) TableName() string {
	return "sys_practices"
}
