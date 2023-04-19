package system

import (
	"github.com/satori/go.uuid"
	"server/global"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:学生UUID"`                                                     // 学生UUID
	StuNumber   string         `json:"userName" gorm:"index;comment:学生学号"`                                                   // 学生登录名
	Password    string         `json:"-"  gorm:"comment:学生登录密码"`                                                             // 学生登录密码
	StuName     string         `json:"stuName" gorm:"default:系统学生;comment:学生姓名"`                                             // 学生昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:学生侧边主题"`                                          // 学生侧边主题
	HeaderImg   string         `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:学生头像"` // 学生头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:学生角色ID"`                                        // 学生角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:学生角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:学生手机号"`                     // 学生手机号
	Email       string         `json:"email"  gorm:"comment:学生邮箱"`                      // 学生邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:学生是否被冻结 1正常 2冻结"` //学生是否被冻结 1正常 2冻结
	Profession  string         `json:"profession" gorm:"comment:所属专业全称"`
	Grade       string         `json:"grade" gorm:"所属年级"`
	Class       string         `json:"class" gorm:"班级简称"`
}

func (SysUser) TableName() string {
	return "sys_users"
}
