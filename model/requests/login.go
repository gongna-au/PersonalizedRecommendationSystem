package requests

type LoginByPhoneRequest struct {
	Email    string `json:"email,omitempty" valid:"email"`
	Password string `json:"password,omitempty" valid:"password"`
} //@name LoginByPhoneRequest

type GetUserByIdRequest struct {
	Id int `json:"id,omitempty" valid:"id"`
} //@name LoginByPhoneRequest
