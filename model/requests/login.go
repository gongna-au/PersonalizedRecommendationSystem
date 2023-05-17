package requests

type LoginByEmailRequest struct {
	Email    string `json:"email,omitempty" valid:"email"`
	Password string `json:"password,omitempty" valid:"password"`
} //@name LoginByPhoneRequest

type UpdateUserRequest struct {
	Email       string `json:"email,omitempty" valid:"email"`
	Password    string `json:"password,omitempty" valid:"password"`
	NewPassword string `json:"newpassword,omitempty" valid:"newpassword"`
} //@name UpdateUserRequest

type GetUserByIdRequest struct {
	Id int `json:"id,omitempty" valid:"id"`
} //@name LoginByPhoneRequest
