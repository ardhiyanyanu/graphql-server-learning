package utility

import "github.com/gin-gonic/gin"

type ResponseDBS struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseJSON(c *gin.Context, statusCode int, responseCode int, message interface{}) {
	c.JSON(statusCode, ResponseDBS{
		Code:    responseCode,
		Message: message,
	})
}

func ResponseBillingSuccessfulPaid(c *gin.Context) {
	c.JSON(200, ResponseDBS{
		Code:    200001,
		Message: "Billing Sucessfully Paid",
	})
}

func ResponseSuccess(c *gin.Context) {
	c.JSON(200, ResponseDBS{
		Code:    200001,
		Message: "Success",
	})
}

func ResponseSuccessCreated(c *gin.Context) {
	c.JSON(201, ResponseDBS{
		Code:    201001,
		Message: "Success",
	})
}

func ResponsePayloadNotStandard(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400001,
		Message: "Payload Is Not Standard",
	})
}

func ResponseInvalidRequest(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400001,
		Message: "Invalid Request",
	})
}

func ResponseInvalidCustomerId(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400001,
		Message: "Invalid CustomerID",
	})
}

func ResponseInvalidEmail(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400001,
		Message: "Invalid Email",
	})
}

func ResponseInvalidPassword(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400001,
		Message: "Invalid Password",
	})
}

func ResponseEmailHasBeenUsed(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400002,
		Message: "Email has been used",
	})
}

func ResponseProductNameHasBeenUsed(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400002,
		Message: "Product Name Has Been Used",
	})
}

func ResponseBillingAlreadyPaid(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400004,
		Message: "Billing Already Paid",
	})
}

func ResponseBillingAlreadyExpired(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400005,
		Message: "Billing Already Expired",
	})
}

func ResponseBillingAmountNotMatch(c *gin.Context) {
	c.JSON(400, ResponseDBS{
		Code:    400009,
		Message: "Billing Amount Not Match",
	})
}

func ResponseDataNotFound(c *gin.Context) {
	c.JSON(404, ResponseDBS{
		Code:    404001,
		Message: "data not found",
	})
}

func ResponseErrorCreateNewUser(c *gin.Context) {
	c.JSON(500, ResponseDBS{
		Code:    500001,
		Message: "Create New User",
	})
}
