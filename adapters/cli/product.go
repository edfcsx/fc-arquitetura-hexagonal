package cli

import (
	"fmt"
	"github.com/edfcsx/fc-arquitetura-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface,
	action string,
	id string,
	name string,
	price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(name, price)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s.",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable", "disable":
		product, err := service.Get(id)
		if err != nil {
			return result, err
		}

		var res application.ProductInterface

		if action == "enable" {
			res, err = service.Enable(product)
			if err != nil {
				return result, err
			}
		} else {
			res, err = service.Disable(product)
			if err != nil {
				return result, err
			}
		}

		if action == "enable" {
			result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
		} else {
			result = fmt.Sprintf("Product %s has been disabled.", res.GetName())
		}
	default:
		res, err := service.Get(id)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product Id: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}

	return result, nil
}
