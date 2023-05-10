package goods

import (
	db "PersonalizedRecommendationSystem/database"

	"PersonalizedRecommendationSystem/model"
)

type GetGoodsInStockByNameModel struct {
	model.BaseModel
	Name string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
}

type GetGoodsInStockBySupplierModel struct {
	model.BaseModel
	Supplier string `json:"supplier,omitempty" gorm:"column:supplier;" binding:"required"`
}

type GetGoodsInStockByNumModel struct {
	model.BaseModel
	Num int `json:"num,omitempty" gorm:"column:num;" binding:"required"`
}

type GoodsInStockModel struct {
	model.BaseModel
	Barcode       string  `json:"barcode,omitempty" gorm:"column:barcode;" binding:"required"`
	Name          string  `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Unit          string  `json:"unit,omitempty" gorm:"column:unit;" binding:"required"`
	Supplier      string  `json:"supplier,omitempty" gorm:"column:supplier;" binding:"required"`
	Brand         string  `json:"brand,omitempty" gorm:"column:brand;" binding:"required"`
	Specification string  `json:"specification,omitempty" gorm:"column:specification;" binding:"required"`
	Purchaseprice float64 `json:"purchaseprice,omitempty" gorm:"column:purchaseprice;" binding:"required"`
	Retailprice   float64 `json:"retailprice,omitempty" gorm:"column:retailprice;" binding:"required"`
	Num           int     `json:"num,omitempty" gorm:"column:num;" binding:"required"`
	Time          string  `json:"time,omitempty" gorm:"column:time;" binding:"required"`
}

func (u *GetGoodsInStockByNameModel) GetGoodsInStockByName() ([]GoodsInStockModel, error) {
	goodsList := make([]GoodsInStockModel, 0)
	query := db.DB.Table("tbl_storehouse").
		Where("name = ?", u.Name)
	if err := query.Find(&goodsList).Error; err != nil {
		return nil, err
	}
	return goodsList, nil
}

func (u *GetGoodsInStockBySupplierModel) GetGoodsInStockByName() ([]GoodsInStockModel, error) {
	goodsList := make([]GoodsInStockModel, 0)
	query := db.DB.Table("tbl_storehouse").
		Where("supplier = ?", u.Supplier)
	if err := query.Find(&goodsList).Error; err != nil {
		return nil, err
	}
	return goodsList, nil

}

func (u *GetGoodsInStockByNumModel) GetGoodsInStockByName() ([]GoodsInStockModel, error) {
	goodsList := make([]GoodsInStockModel, 0)
	query := db.DB.Table("tbl_storehouse").
		Where("num <= ?", u.Num)
	if err := query.Find(&goodsList).Error; err != nil {
		return nil, err
	}
	return goodsList, nil
}
