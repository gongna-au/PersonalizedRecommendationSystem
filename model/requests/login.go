package requests

type LoginByPhoneRequest struct {
	Phone    string `json:"phone,omitempty" valid:"phone"`
	Password string `json:"password,omitempty" valid:"password"`
} //@name LoginByPhoneRequest
