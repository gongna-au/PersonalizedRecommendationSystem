package requests

type CreateOrderRequest struct {
	Phone    string  `json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
	SaleId   int     `json:"sid,omitempty" gorm:"column:sid;" binding:"required"`
	Barcode  string  `json:"barcode,omitempty" gorm:"column:barcode;" binding:"required"`
	Discount float64 `json:"discount" gorm:"column:discount;" binding:"required"`
} //@name CreateOrderRequest

type GetOrdersByUserRequest struct {
	Phone string `json:"phone,omitempty" gorm:"column:phone;" binding:"required"`
} //@name GetOrdersByUserRequest

type GetOrdersByTimeRequest struct {
	Start string `json:"start,omitempty" gorm:"column:start;" binding:"required"`
	End   string `json:"end,omitempty" gorm:"column:end;" binding:"required"`
} //@name GetOrdersByTimeRequest

type CreateReturnOrderRequest struct {
	Ordercode string `json:"ordercode,omitempty" gorm:"column:ordercode;" binding:"required"`
	SaleId    int    `json:"sid,omitempty" gorm:"column:sid;" binding:"required"`
} //@name CreateReturnOrderRequest
