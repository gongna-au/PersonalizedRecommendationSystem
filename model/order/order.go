package order

import (
	"fmt"
	"math/rand"
	"time"

	db "PersonalizedRecommendationSystem/database"
)

type OrderModel struct {
	Id        int     `json:"id,omitempty" gorm:"column:id;" binding:"required"`
	Phone     string  `json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	SaleId    int     `json:"sid,omitempty" gorm:"column:sid;" binding:"required"`
	Ordercode string  `json:"ordercode,omitempty" gorm:"column:ordercode;" binding:"required"`
	Barcode   string  `json:"barcode,omitempty" gorm:"column:barcode;" binding:"required"`
	Saleprice float64 `json:"saleprice,omitempty" gorm:"column:saleprice;" binding:"required"`
	Discount  float64 `json:"discount,omitempty" gorm:"column:discount;" binding:"required"`
	Payprice  float64 `json:"payprice,omitempty" gorm:"column:payprice;" binding:"required"`
} //@name OrderModel

type OrderReturnModel struct {
	Id          int     `json:"id,omitempty" gorm:"column:id;" binding:"required"`
	Phone       string  `json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	SaleId      int     `json:"sid,omitempty" gorm:"column:sid;" binding:"required"`
	Ordercode   string  `json:"ordercode,omitempty" gorm:"column:ordercode;" binding:"required"`
	Barcode     string  `json:"barcode,omitempty" gorm:"column:barcode;" binding:"required"`
	Saleprice   float64 `json:"saleprice,omitempty" gorm:"column:saleprice;" binding:"required"`
	Discount    float64 `json:"discount,omitempty" gorm:"column:discount;" binding:"required"`
	Payprice    float64 `json:"payprice,omitempty" gorm:"column:payprice;" binding:"required"`
	Returnprice float64 `json:"returnprice ,omitempty" gorm:"column:returnprice;" binding:"required"`
} //@name OrderReturnModel

func (u *OrderModel) Create() error {
	return db.DB.
		Table("tbl_order").
		Create(u).Error
}

func (u *OrderReturnModel) Create() error {
	return db.DB.
		Table("tbl_order_return").
		Create(u).Error
}

func (u *OrderModel) GetOrdersByUser() ([]OrderModel, error) {
	result := make([]OrderModel, 0)
	query := db.DB.Table("tbl_order").
		Where("phone = ?", u.Phone)
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *OrderModel) GetOrdersByTime(start string, end string) ([]OrderModel, error) {
	result := make([]OrderModel, 0)
	query := db.DB.Table("tbl_order").
		Where("time >= ? And time <= ?", start, end)
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *OrderModel) GetOrdersProfitByTime(start string, end string) (float64, error) {
	//"2022-11-18 20:00:00"
	fmt.Println(start)
	fmt.Println(end)
	result := make([]OrderModel, 0)
	query := db.DB.Table("tbl_order").
		Where("time >= ? And time <= ?", fmt.Sprintf("%s", start), fmt.Sprintf("%s", end))
	if err := query.Find(&result).Error; err != nil {
		return 0, err
	}
	var sum float64
	for _, v := range result {
		sum = sum + (v.Payprice - v.Saleprice)
	}
	return sum, nil
}

func (u *OrderModel) GetOrdersProfitByTimeAndName(start string, end string, barcode string) (float64, error) {

	result := make([]OrderModel, 0)
	query := db.DB.Table("tbl_order").
		Where("time >= ? And time <= ? and barcode = ? ", start, end, barcode)
	if err := query.Find(&result).Error; err != nil {
		return 0, err
	}
	var sum float64
	for _, v := range result {
		sum = sum + (v.Payprice - v.Saleprice)
	}

	return sum, nil
}

func (u *OrderModel) GetOrdersByBarcode() (OrderModel, error) {
	result := OrderModel{}
	query := db.DB.Table("tbl_order").
		Where("barcode = ? ", u.Barcode)
	if err := query.First(&result).Error; err != nil {
		return OrderModel{}, err
	}
	return result, nil
}

func (u *OrderModel) GetOrdersByOrdercode() (OrderModel, error) {
	result := OrderModel{}
	query := db.DB.Table("tbl_order").
		Where("ordercode = ?", u.Ordercode)
	if err := query.First(&result).Error; err != nil {
		return OrderModel{}, err
	}
	return result, nil
}

func GetTimeTick64() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetTimeTick32() int32 {
	return int32(time.Now().Unix())
}

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

// 基础做法 日期20191025时间戳1571987125435+3位随机数
func GenerateCode() string {
	date := GetFormatTime(time.Now())
	r := rand.Intn(1000)
	code := fmt.Sprintf("%s%d%03d", date, GetTimeTick64(), r)
	return code
}
