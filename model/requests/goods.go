package requests

type GetGoodsInStockByNameRequest struct {
	Name string `json:"name,omitempty" binding:"required"`
} //@name GetGoodsInStockByNameRequest

type GetGoodsInStockBySupplierRequest struct {
	Supplier string `json:"supplier,omitempty" binding:"required"`
} //@name GetGoodsInStockBySupplierRequest

type GetGoodsInStockByNumRequest struct {
	Num int `json:"num,omitempty" binding:"required"`
} //@name GetGoodsInStockByNumRequest
