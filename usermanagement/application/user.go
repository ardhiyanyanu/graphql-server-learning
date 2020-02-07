package application

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alterra/graphql-server/usermanagement/utility"

	"golang.org/x/crypto/bcrypt"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// loginParam struct
type loginParam struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// loginOpsParam struct
type loginOpsParam struct {
	Username string `json:"username" valid:"required"`
	Password string `json:"password" valid:"required"`
}

// userParam struct
type userParam struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
	RoleID   string `json:"roleId" valid:"required"`
}

// userParam struct
type updateUserParam struct {
	Name           string `json:"name"`
	OldPassword    string `json:"oldPassword"`
	NewPassword    string `json:"newPassword"`
	ProfilePicture string `json:"profilePicture"`
}

// approvalParam struct
type approvalParam struct {
	Approval     bool   `json:"approval"`
	CustomerType string `json:"customerType"`
	AdminFee     int    `json:"adminFee"`
	VA           bool   `json:"virtualAccount"`
}

// Login function
func Login(c *gin.Context) {
	// Read the Body content
	var dataUser loginParam
	if err := c.ShouldBindJSON(&dataUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate if required parameter is not complete in payload
	check, err := govalidator.ValidateStruct(dataUser)
	if !check {
		utility.ResponseInvalidRequest(c)
		return
	}

	// trim space in leading and trailing string of password and email
	dataUser.Password = strings.TrimSpace(dataUser.Password)
	dataUser.Email = strings.ToLower(strings.TrimSpace(dataUser.Email))

	isEmailCorrect, _ := utility.CheckRegexEmail(dataUser.Email)
	if !isEmailCorrect {
		utility.ResponseInvalidEmail(c)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(dataUser.Password), 8)
	fmt.Println(string(hashedPassword))

	// Check in your db if the user exists or not
	// userDetail, err := infrastructure.GetUserDetail(dataUser.Email, dataUser.Password)
	// if err != nil {
	// 	return ResponseJSON(c, 401, 401001, "Incorrect email or password")
	// }

	// Validate if user still active
	// if userDetail.Status == "Reject" {
	// 	return ResponseJSON(c, 401, 401002, "Banned User")
	// } else if userDetail.Status == Pending {
	// 	return ResponseJSON(c, 403, 403001, "Waiting For Approval")
	// }

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = "yanu@alterra.id" //userDetail.Email
	claims["customer_id"] = "horas"     //userDetail.CustomerID
	// claims[RoleID] = userDetail.RoleID
	// claims["role_name"] = roleDetail[0].Name
	// claims["virtual_account"] = userSetting.VA
	// claims["secret_key"] = userSetting.SecretKey
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// secretToken := viper.Sub("token").GetString("secret")
	secretToken := "secret token that long enoght"
	t, err := token.SignedString([]byte(secretToken))
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":           200001,
		"message":        "Success",
		"token":          t,
		"name":           "", //userDetail.Username,
		"profilePicture": "", //userDetail.ProfilePicture
	})
}

// RegisterNewCustomerAndUser function
func RegisterNewCustomerAndUser(c gin.Context) {
	// // Read the body content
	// var bodyBytes []byte
	// if c.Request().Body != nil {
	// 	bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	// }

	// // validate if the params i payload more than required
	// actualRegister := entity.NewRegister{}
	// e := reflect.ValueOf(&actualRegister).Elem()
	// var checkNumOfParams map[string]interface{}
	// json.Unmarshal([]byte(bodyBytes), &checkNumOfParams)
	// if len(checkNumOfParams) > e.NumField() {
	// 	return ResponseInvalidRequest(c)
	// }

	// // validate if the body JSON is empty
	// var newRegistration entity.NewRegister
	// err := json.Unmarshal([]byte(bodyBytes), &newRegistration)
	// if err != nil {
	// 	return ResponseInvalidRequest(c)
	// }

	// // validate if required parameter is not complete in payload
	// check, err := govalidator.ValidateStruct(newRegistration)
	// if !check {
	// 	return ResponseInvalidRequest(c)
	// }

	// // validate if roleID != 1 (as superadmin)
	// if newRegistration.User.RoleID != 1 {
	// 	return ResponseInvalidRequest(c)
	// }

	// // trim space in leading and trailing string of customerId, email, password, and name
	// newRegistration.Customer.CustomerID = strings.ToLower(strings.TrimSpace(newRegistration.Customer.CustomerID))
	// newRegistration.Customer.Name = strings.ToLower(strings.TrimSpace(newRegistration.Customer.Name))
	// newRegistration.User.CustomerID = strings.ToLower(strings.TrimSpace(newRegistration.User.CustomerID))
	// newRegistration.User.Email = strings.ToLower(strings.TrimSpace(newRegistration.User.Email))
	// newRegistration.User.Password = strings.TrimSpace(newRegistration.User.Password)

	// // validate if customerId in Customer not equal to User
	// if newRegistration.Customer.CustomerID != newRegistration.User.CustomerID {
	// 	return ResponseInvalidCustomerId(c)
	// }

	// isCustomerIDCorrect, _ := shared_kernel.CheckRegexCustomerID(newRegistration.Customer.CustomerID)
	// if !isCustomerIDCorrect {
	// 	return ResponseInvalidCustomerId(c)
	// }

	// isCustomerNameCorrect, _ := shared_kernel.CheckRegexCustomerName(newRegistration.Customer.Name)
	// if !isCustomerNameCorrect {
	// 	return ResponseJSON(c, 400, 400001, "Invalid Name")
	// }

	// errString := validateEmailPasswordCorrect(newRegistration.User.Email, newRegistration.User.Password)
	// if errString != "" {
	// 	return ResponseJSON(c, 400, 400001, errString)
	// }

	// // Check if customerId already reqistered in DB or not
	// isCustomerExist := infrastructure.CheckCustomerIDIsExistInCustomer(newRegistration.Customer.CustomerID)
	// if isCustomerExist {
	// 	return ResponseJSON(c, 400, 400001, "CustomerID has been used")
	// }

	// // Check if email already used in same customerID
	// isEmailExist := infrastructure.CheckEmailExist("", newRegistration.User.Email)
	// if isEmailExist {
	// 	return ResponseEmailHasBeenUsed(c)
	// }

	// // register new customer
	// newRegistration.Customer.CreatedBy = System
	// newRegistration.Customer.Status = Pending
	// err = infrastructure.InsertNewCustomer(newRegistration.Customer)
	// if err != nil {
	// 	return ResponseJSON(c, 500, 500001, "Error Create New Customer")
	// }

	// // register new user
	// newRegistration.User.CreatedBy = System
	// newRegistration.User.UpdatedBy = System
	// newRegistration.User.Status = "Active"
	// newRegistration.User.Status = Pending
	// err = infrastructure.InsertNewUser(newRegistration.User)
	// if err != nil {
	// 	return ResponseErrorCreateNewUser(c)
	// }

	// return ResponseSuccessCreated(c)
}
