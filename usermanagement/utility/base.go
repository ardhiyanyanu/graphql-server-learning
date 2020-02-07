package utility

import (
	"regexp"
	"strconv"
	"strings"
)

// CorrectBillingStatus variabel
var CorrectBillingStatus = map[string]bool{
	"paid":   true,
	"unpaid": true,
}

// CorrectSortWith variable
var CorrectSortWith = map[string]bool{
	"asc":  true,
	"desc": true,
}

// CorrectSortBy variable
var CorrectSortBy = map[string]bool{
	"product_id":   true,
	"pelanggan_id": true,
	"amount":       true,
	"status":       true,
	"expired_date": true,
	"created_at":   true,
	"payment_date": true,
}

// PaymentCorrectSortBy variable
var PaymentCorrectSortBy = map[string]bool{
	"product_id":   true,
	"pelanggan_id": true,
	"amount":       true,
	"status":       true,
	"payment_date": true,
	"expired_date": true,
	"invoice_date": true,
}

// PisahHariJam function
func PisahHariJam(data string) string {
	tanggal := strings.Split(data, "T")[0]
	jam := strings.Split(data, "T")[1]
	jamUTC := strings.Split(jam, "+")[0]
	return tanggal + " " + jamUTC
}

func ValidateDate(date string) bool {
	year, err := strconv.Atoi(strings.Split(date, "-")[0])
	if err != nil {
		return false
	}
	if year > 9999 || year < 1000 {
		return false
	}
	month, err := strconv.Atoi(strings.Split(date, "-")[1])
	if err != nil {
		return false
	}
	if month > 12 || month < 1 {
		return false
	}

	day, err := strconv.Atoi(strings.Split(date, "-")[2])
	if err != nil {
		return false
	}
	if month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12 {
		if day > 31 || day < 1 {
			return false
		}
	} else if month == 4 || month == 6 || month == 9 || month == 11 {
		if day > 30 || day < 1 {
			return false
		}
	} else if month == 2 {
		if isLeapYear(year) {
			if day > 29 || day < 1 {
				return false
			}
		} else {
			if day > 28 || day < 1 {
				return false
			}
		}
	}
	return true
}

// Check if the inputted year is leap year
func isLeapYear(year int) bool {
	leapFlag := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leapFlag = true
			} else {
				leapFlag = false
			}
		} else {
			leapFlag = true
		}
	} else {
		leapFlag = false
	}
	return leapFlag
}

// CheckRegexEmail function
func CheckRegexEmail(email string) (bool, error) {
	regexEmail, _ := regexp.Compile(`^([a-zA-Z0-9]+)([._-]??[a-zA-Z0-9]+)+@([a-zA-Z0-9]+)([.-]??[a-zA-Z0-9]+)+([.]{1}[a-zA-Z0-9]{1,3})$`)
	result := regexEmail.MatchString(email)
	return result, nil
}

// CheckRegexMobilePhone function
func CheckRegexMobilePhone(mobilePhone string) (bool, error) {
	regexMobilePhone, _ := regexp.Compile(`^[+62|62|0][0-9]*$`)
	result := regexMobilePhone.MatchString(mobilePhone)
	return result, nil
}

// CheckRegexPassword function
func CheckRegexPassword(password string) (bool, error) {
	if len(password) < 6 || len(password) > 16 {
		return false, nil
	}
	var lowercase = "abcdefghijklmnopqrstuvwxyz"
	var uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "0123456789"
	var signs = "~!@#$%^&*()_+"
	if strings.ContainsAny(password, lowercase) && strings.ContainsAny(password, uppercase) && strings.ContainsAny(password, numbers) && strings.ContainsAny(password, signs) {
		return true, nil
	}
	return false, nil
}

// CheckRegexCustomerID function
func CheckRegexCustomerID(customerID string) (bool, error) {
	// check CustomerId (alphanumeric with length 4-8)
	regexCustomerID, _ := regexp.Compile(`^[a-zA-Z0-9]{4,8}$`)
	result := regexCustomerID.MatchString(customerID)
	return result, nil
}

// CheckRegexCustomerName function
func CheckRegexCustomerName(customerName string) (bool, error) {
	// check Customer name (any kind of character with length 4-255)
	regexCustomerName, _ := regexp.Compile(`^[\w\W]{4,255}$`)
	result := regexCustomerName.MatchString(customerName)
	return result, nil
}

// CheckRegexProductDescription function
func CheckRegexProductDescription(productDescription string) (bool, error) {
	// check product description (any kind of character with length 4-512)
	regexProductDescription, _ := regexp.Compile(`^[ -~]{0,512}$`)
	result := regexProductDescription.MatchString(productDescription)
	return result, nil
}

// CheckRegexProductID function
func CheckRegexProductID(productID string) (bool, error) {
	regexProductID, _ := regexp.Compile(`^[a-zA-Z0-9]{1,32}$`)
	result := regexProductID.MatchString(productID)
	return result, nil
}

// CheckRegexProductName function
func CheckRegexProductName(productName string) (bool, error) {
	// Check name alphanumeric or not
	regexProductName, _ := regexp.Compile(`^[a-zA-Z0-9 ]{1,255}$`)
	result := regexProductName.MatchString(productName)
	return result, nil
}

// CheckRegexRoleName function
func CheckRegexRoleName(roleName string) (bool, error) {
	// Check name alphanumeric or not
	regexRoleName, _ := regexp.Compile(`^[a-zA-Z0-9 ]{4,255}$`)
	result := regexRoleName.MatchString(roleName)
	return result, nil
}

// CheckRegexCustomerType function
func CheckRegexCustomerType(customerType string) (bool, error) {
	regexCustomerType, _ := regexp.Compile(`^[\w\W]{4,128}$`)
	result := regexCustomerType.MatchString(customerType)
	return result, nil
}

// CheckRegexProductType function
func CheckRegexProductType(ProductType string) (bool, error) {
	regexProductType, _ := regexp.Compile(`^([a-zA-Z]{1})([a-zA-Z0-9]{3,20})$`)
	result := regexProductType.MatchString(ProductType)
	return result, nil
}

// CheckRegexDataLabelOutputKey function
func CheckRegexDataLabelOutputKey(ProductType string) (bool, error) {
	regexProductType, _ := regexp.Compile(`^([a-zA-Z]{1})([a-zA-Z0-9_]{3,20})$`)
	result := regexProductType.MatchString(ProductType)
	return result, nil
}
