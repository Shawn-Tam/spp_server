package request

import (
	"server/model/system"
)

// Register User register structure
type Register struct {
	StuNumber    string `json:"stuNumber" example:"学生名"`
	Password     string `json:"passWord" example:"密码"`
	StuName      string `json:"stuName" example:"姓名"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	Profession   string `json:"profession "`
	Grade        string `json:"grade"`
	Class        string `json:"class"`
}

// User login structure
type Login struct {
	StuNumber string `json:"stuNumber"` // 学生名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	StuName      string                `json:"stuName" gorm:"default:系统学生;comment:学生姓名"`                                             // 学生昵称
	Phone        string                `json:"phone"  gorm:"comment:学生手机号"`                                                          // 学生手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:学生邮箱"`                                                           // 学生邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:学生头像"` // 学生头像
	SideMode     string                `json:"sideMode"  gorm:"comment:学生侧边主题"`                                                      // 学生侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结学生"`                                                           //冻结学生
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
