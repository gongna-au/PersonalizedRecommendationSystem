package profit

import (
	"PersonalizedRecommendationSystem/model/goods"
	"PersonalizedRecommendationSystem/model/order"

	"PersonalizedRecommendationSystem/model/requests"
	"PersonalizedRecommendationSystem/model/response"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Get Profit By Time
// @Description  得到一段时间内部超市总的利润
// @Tags         profit
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetProfitByTimeRequest true "start--起始时间|| end -- 结束时间"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/profit   [post]
// GetProfitByTime 得到一段时间内部超市总的利润
func GetProfitByTime(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetProfitByTimeRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := order.OrderModel{}
	result, err := model.GetOrdersProfitByTime(request.Start, request.End)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	response.CreatedJSON(c, gin.H{
		"profit": result,
	})
}

// ShowAccount godoc
// @Summary      Get Profit By Time And Name
// @Description  得到一段时间内某个商品总的利润
// @Tags         profit
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.CreateOrderRequest true "start--起始时间 || end -- 结束时间 || name-- 商品名  "
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/profit/name  [post]
// GetProfitByTimeAndName 得到一段时间内某个商品总的利润
func GetProfitByTimeAndName(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetProfitByTimeAndNameRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := order.OrderModel{}
	goods := goods.GetGoodsInStockByNameModel{
		Name: request.Name,
	}
	r, err := goods.GetGoodsInStockByName()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	} else {
		if len(r) > 0 {
			result, err := model.GetOrdersProfitByTimeAndName(request.Start, request.End, r[0].Barcode)
			if err != nil {
				response.Abort500(c, err.Error())
				return
			}
			response.CreatedJSON(c, gin.H{
				"profit": result,
			})
		}
	}
}
