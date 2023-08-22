package api

import (
	"github.com/bagus-k/simple-bank/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		return utils.IsSupportCurrency(currency)
	}
	return false
}