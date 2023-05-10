package user

import (
	"errors"
	"fmt"

	db "PersonalizedRecommendationSystem/database"

	"PersonalizedRecommendationSystem/model"
)

type CommonUserModel struct {
	model.BaseModel
	Name     string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Phone    string `json:"phone" gorm:"column:phone;" binding:"required"`
	Password string `json:"password" gorm:"column:password;" binding:"required"`
}

type StaffUserModel struct {
	model.BaseModel
	Name            string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Phone           string `json:"phone" gorm:"column:phone;" binding:"required"`
	Password        string `json:"password" gorm:"column:password;" binding:"required"`
	Permissionlevel int    `json:"permissionlevel" gorm:"column:permissionlevel;" binding:"required"`
}

type MemberUserModel struct {
	model.BaseModel
	Name     string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Phone    string `json:"phone" gorm:"column:phone;" binding:"required"`
	Password string `json:"password" gorm:"column:password;" binding:"required"`
}

func (u *CommonUserModel) TableName() string {
	return "tbl_user_common"
}

func (u *CommonUserModel) Create() error {
	return db.DB.
		Table("tbl_user_common").
		Create(u).Error
}

func (u *CommonUserModel) Save() (err error) {
	return db.DB.
		Table("tbl_user_common").
		Save(u).Error
}

func (u *StaffUserModel) TableName() string {
	return "tbl_staff"
}

func (u *StaffUserModel) Create() error {
	return db.DB.
		Table("tbl_staff").
		Create(u).Error
}

func (u *StaffUserModel) Save() (err error) {
	return db.DB.
		Table("tbl_staff").
		Save(u).Error
}

func (u *MemberUserModel) TableName() string {
	return "tbl_member"
}

func (u *MemberUserModel) Create() error {
	return db.DB.
		Table("tbl_member").
		Create(u).Error
}

func (u *MemberUserModel) Save() (err error) {
	return db.DB.
		Table("tbl_member").
		Save(u).Error
}

func GetCommonUserByPhoneAndPassword(phone string, password string) (*CommonUserModel, error) {
	u := &CommonUserModel{}
	d := db.DB.
		Table("tbl_user_common").
		Where("phone = ? AND password = ?", phone, password).First(u)
	return u, d.Error
}

func GetStaffUserByPhoneAndPassword(phone string, password string) (*StaffUserModel, error) {
	u := &StaffUserModel{}
	d := db.DB.
		Table("tbl_staff").
		Where("phone = ? AND password = ?", phone, password).First(u)
	return u, d.Error
}

func GetMemberUserByPhoneAndPassword(phone string, password string) (*MemberUserModel, error) {
	u := &MemberUserModel{}
	d := db.DB.
		Table("tbl_member").
		Where("phone = ? AND password = ?", phone, password).First(u)
	return u, d.Error
}

// JudgeCommonUserPhoneExist 通过手机号来判断手机号是非被注册
func JudgeCommonUserPhoneExist(phone string) error {
	var userModel CommonUserModel
	db.DB.
		Table("tbl_user_common").
		Where("phone = ?", phone).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this phone has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}

// JudgeStaffUserPhoneExist 通过手机号来判断手机号是非被注册
func JudgeStaffUserPhoneExist(phone string) error {
	var userModel StaffUserModel
	db.DB.
		Table("tbl_staff").
		Where("phone = ?", phone).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this phone has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}

// JudgeMemberUserPhoneExist 通过手机号来判断手机号是非被注册
func JudgeMemberUserPhoneExist(phone string) error {
	var userModel MemberUserModel
	db.DB.
		Table("tbl_member").
		Where("phone = ?", phone).First(&userModel)
	if userModel.BaseModel.ID > 0 {
		fmt.Println("this phone has been registered")
		return errors.New("this phone has been registered")
	} else {
		return nil
	}
}
