package utils

import (
	"github.com/astaxie/beego/logs"
)

func VerifyField(fieldValue interface{}, fieldType ...string) bool {
	logs.Info("validating field %+v with %+v", fieldValue, fieldType)
	return true
}
