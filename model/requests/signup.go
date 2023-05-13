package requests

type SignupPhoneExistRequest struct {
	Phone string `valid:"phone" json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
} //@name SignupPhoneExistRequest

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Email    string `valid:"email" json:"email,omitempty" gorm:"column:email;" binding:"required"`
	Username string `valid:"username" json:"username,omitempty" gorm:"column:username;" binding:"required"`
	Password string `valid:"password" json:"password,omitempty" gorm:"column:password;" binding:"required"`
} //@name SignupUsingPhoneRequest
