package requests

type SignupPhoneExistRequest struct {
	Phone string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
} //@name SignupPhoneExistRequest

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Phone    string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	Name     string `valid:"name" json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Password string `valid:"password" json:"password,omitempty" gorm:"column:password;" binding:"required"`
} //@name SignupUsingPhoneRequest
