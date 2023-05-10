package goods

import (
	"PersonalizedRecommendationSystem/model/goods"
	"PersonalizedRecommendationSystem/model/requests"
	"PersonalizedRecommendationSystem/model/response"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Get Goods In Stock By Name
// @Description  通过商品名字查询商品的库存情况
// @Tags         goods
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetGoodsInStockByNameRequest  true  "Name--商品名"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/goods/name   [post]
// GetGoodsInStockByName 通过商品名字查询商品的库存情况
func GetGoodsInStockByName(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetGoodsInStockByNameRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := goods.GetGoodsInStockByNameModel{
		Name: request.Name,
	}
	data, err := model.GetGoodsInStockByName()
	if err != nil {
		response.Abort500(c, err.Error())
	}
	response.Data(c, data)
}

// ShowAccount godoc
// @Summary      Get Goods In Stock By Supplier
// @Description  通过商品供应商查询商品的库存情况
// @Tags         goods
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetGoodsInStockBySupplierRequest  true  "Supplier--供应商"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router        /api/v1/goods/supplier  [post]
// GetGoodsInStockBySupplier 通过商品供应商查询商品的库存情况
func GetGoodsInStockBySupplier(c *gin.Context) {

	// 1. 验证表单
	request := requests.GetGoodsInStockBySupplierRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := goods.GetGoodsInStockBySupplierModel{
		Supplier: request.Supplier,
	}
	data, err := model.GetGoodsInStockByName()
	if err != nil {
		response.Abort500(c, err.Error())
	}
	response.Data(c, data)
}

// ShowAccount godoc
// @Summary      Get Goods In Stock By Num
// @Description  通过商品的剩余量供查询商品的库存情况
// @Tags         goods
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetGoodsInStockByNumRequest  true  "Num--数量"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/goods/num   [post]
// GetGoodsInStockByNum  通过商品的剩余量供查询商品的库存情况
func GetGoodsInStockByNum(c *gin.Context) {
	// 1. 验证表单
	request := requests.GetGoodsInStockByNumRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := goods.GetGoodsInStockByNumModel{
		Num: request.Num,
	}
	data, err := model.GetGoodsInStockByName()
	if err != nil {
		response.Abort500(c, err.Error())
	}
	response.Data(c, data)
}
