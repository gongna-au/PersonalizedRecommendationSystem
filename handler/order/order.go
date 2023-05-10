package order

import (
	"fmt"

	"PersonalizedRecommendationSystem/model/order"

	"PersonalizedRecommendationSystem/model/requests"

	"PersonalizedRecommendationSystem/model/response"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Create Order
// @Description  创建一份订单
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.CreateOrderRequest true "Phone--付款人账户|| SaleId-- 收银员ID ||  Barcode-- 商品条形码  || Discount--折扣 "
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/orders/   [post]
// CreateOrder 创建一份订单
func CreateOrder(c *gin.Context) {
	// 1. 验证表单
	request := requests.CreateOrderRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据

	model := order.OrderModel{
		Phone:     request.Phone,
		SaleId:    request.SaleId,
		Ordercode: order.GenerateCode(),
		Barcode:   request.Barcode,
		Discount:  request.Discount,
	}
	result, err := model.GetOrdersByBarcode()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}

	model.Saleprice = result.Saleprice
	model.Payprice = result.Saleprice - model.Discount
	model.Create()
	if model.Id > 0 {
		response.Data(c, "User need pay for"+fmt.Sprintf("%.3f", model.Payprice))
	} else {
		response.Abort500(c, "创建订单，请稍后尝试~")
		return
	}
}

// ShowAccount godoc
// @Summary     Create Return Order
// @Description  创建一份退货订单
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.CreateReturnOrderRequest true "ordercode--订单号码 || sid--负责退货的员工的帐号ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/orders/return   [post]
// CreateReturnOrder 创建一份退货订单
func CreateReturnOrder(c *gin.Context) {

	// 1. 验证表单
	request := requests.CreateReturnOrderRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}

	// 2. 验证成功，创建数据

	model := order.OrderModel{
		Ordercode: request.Ordercode,
	}
	result, err := model.GetOrdersByOrdercode()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	o := order.OrderReturnModel{
		Phone:       result.Phone,
		SaleId:      request.SaleId,
		Ordercode:   order.GenerateCode(),
		Barcode:     result.Barcode,
		Saleprice:   result.Saleprice,
		Discount:    result.Discount,
		Payprice:    result.Payprice,
		Returnprice: result.Payprice,
	}
	o.Create()
	if o.Id > 0 {
		response.Data(c, "Our need pay for user "+fmt.Sprintf("%.3f", o.Payprice))
	} else {
		response.Abort500(c, "退货订单创建失败，请稍后尝试~")
		return
	}
}

// ShowAccount godoc
// @Summary      Get Orders By User
// @Description  用户获取到自己的所有的订单
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetOrdersByUserRequest true "Phone--用户账户"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router      /api/v1/orders/user   [post]
// GetOrdersByUser 用户获取到自己的所有的订单
func GetOrdersByUser(c *gin.Context) {
	// 1. 验证表单
	request := requests.GetOrdersByUserRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}

	// 2. 验证成功，创建数据

	model := order.OrderModel{
		Phone: request.Phone,
	}
	result, err := model.GetOrdersByUser()
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	response.Data(c, result)

}

// ShowAccount godoc
// @Summary      GetOrdersByTime
// @Description  管理人员可以通过这个接口获取到一段时间内部所有的订单记录
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        req  {object}  body requests.GetOrdersByTimeRequest true "start--开始时间|| end-- 结束时间"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Failure      500  {object}  response.Response
// @Router       /api/v1/orders/time   [post]
// GetOrdersByTime 获取到一段时间内的所有订单
func GetOrdersByTime(c *gin.Context) {
	// 1. 验证表单
	request := requests.GetOrdersByTimeRequest{}
	// 2. 绑定
	if err := c.Bind(&request); err != nil {
		response.Error(c, err, "请求失败")
		return
	}
	// 2. 验证成功，创建数据
	model := order.OrderModel{}

	result, err := model.GetOrdersByTime(request.Start, request.End)
	if err != nil {
		response.Abort500(c, err.Error())
		return
	}
	response.Data(c, result)
}
