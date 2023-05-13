package user

import (
	"errors"
	"fmt"

	db "PersonalizedRecommendationSystem/database"

	"PersonalizedRecommendationSystem/model"
)

type CommonUserModel struct {
	model.BaseModel
	UserName string `json:"username,omitempty" gorm:"column:username;" binding:"required"`
	Email    string `json:"email" gorm:"column:email;" binding:"required"`
	Password string `json:"password" gorm:"column:password;" binding:"required"`
}

func (u *CommonUserModel) TableName() string {
	return "users"
}

func (u *CommonUserModel) Create() error {
	return db.DB.
		Table(u.TableName()).
		Create(u).Error
}

func (u *CommonUserModel) Save() (err error) {
	return db.DB.
		Table(u.TableName()).
		Save(u).Error
}

func GetCommonUserByPhoneAndPassword(email string, password string) (*CommonUserModel, error) {
	u := &CommonUserModel{}
	d := db.DB.
		Table(u.TableName()).
		Where("email = ? AND password = ?", email, password).First(u)
	return u, d.Error
}

func GetCommonUserById(id int) (*CommonUserModel, error) {
	u := &CommonUserModel{}
	d := db.DB.
		Table(u.TableName()).
		Where("id = ?  ?", id).First(u)
	return u, d.Error
}

// JudgeCommonUserPhoneExist 通过邮箱来判断手机号是非被注册
func JudgeCommonUserEmailExist(email string) error {
	var userModel CommonUserModel
	db.DB.
		Table("users").
		Where("email = ?", email).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this emial has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}
