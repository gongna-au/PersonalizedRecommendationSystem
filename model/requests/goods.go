package requests

type GetGoodsInStockBySupplierRequest struct {
	Supplier string `json:"supplier,omitempty" binding:"required"`
} //@name GetGoodsInStockBySupplierRequest

type GetGoodsInStockByNumRequest struct {
	Num int `json:"num,omitempty" binding:"required"`
} //@name GetGoodsInStockByNumRequest

type GetResourceByUid struct {
	Uid int `json:"uid,omitempty" gorm:"column:uid;" binding:"required"`
} //@name GetResourceByKey
