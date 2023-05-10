package requests

type GetProfitByTimeRequest struct {
	Start string `json:"start,omitempty" gorm:"column:start;" binding:"required"`
	End   string `json:"end,omitempty" gorm:"column:end;" binding:"required"`
} //@name GetProfitByTimeRequest

type GetProfitByTimeAndNameRequest struct {
	Name  string `json:"name,omitempty" gorm:"column:name;" binding:"required"`
	Start string `json:"start,omitempty" gorm:"column:start;" binding:"required"`
	End   string `json:"end,omitempty" gorm:"column:end;" binding:"required"`
} //@name GetProfitByTimeAndNameRequest
