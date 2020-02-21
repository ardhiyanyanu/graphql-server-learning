package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	model "github.com/alterra/graphql-server/graphql/models"
	"github.com/alterra/graphql-server/graphql/resolver/util"
)

type allProductResponse struct {
	Code    int        `json:"code"`
	Message []*Product `json:"message"`
}

type SpecialDate struct {
	time.Time
}

type Product struct {
	ProductID     string       `json:"productId"`
	Description   *string      `json:"description"`
	CustomerID    string       `json:"customerId"`
	Name          string       `json:"name"`
	ProductType   *string      `json:"productType"`
	ProductTypeID *int         `json:"productTypeId"`
	CreatedBy     *string      `json:"createdBy"`
	UpdatedBy     *string      `json:"updatedBy"`
	CreatedAt     *SpecialDate `json:"createdAt"`
	UpdatedAt     *SpecialDate `json:"updatedAt"`
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02 15:04:05", strInput)
	if err != nil {
		return err
	}

	sd.Time = newTime
	return nil
}

func convertProduct(product *Product) *model.Product {
	temp := model.Product{
		ProductID:     product.ProductID,
		Description:   product.Description,
		CustomerID:    product.CustomerID,
		Name:          product.Name,
		ProductType:   product.ProductType,
		ProductTypeID: product.ProductTypeID,
		CreatedBy:     product.CreatedBy,
		UpdatedBy:     product.UpdatedBy,
		CreatedAt:     &product.CreatedAt.Time,
		UpdatedAt:     &product.UpdatedAt.Time,
	}
	return &temp
}

/*
* convert local product to graphql product because of marshalling problem in date time
 */
func convertProducts(products []*Product) []*model.Product {
	result := make([]*model.Product, 0)
	for _, product := range products {
		result = append(result, convertProduct(product))
	}
	return result
}

func GetAllProduct() ([]*model.Product, error) {
	var client = &http.Client{}

	token := util.GetToken()
	fmt.Println("token: " + token)
	req, _ := http.NewRequest("GET", util.BaseURL+"/yanu/product", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var responseObject allProductResponse
	err = json.NewDecoder(response.Body).Decode(&responseObject)
	if err != nil {
		return nil, err
	}

	return convertProducts(responseObject.Message), nil
}
