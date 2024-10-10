package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/rpraveenkumar/Golang/db/utils"
)

var ValidCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {

	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}
	return false

}
