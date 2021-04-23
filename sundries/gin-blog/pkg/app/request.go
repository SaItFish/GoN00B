// Package app
// @file: request.go
// @date: 2020/12/31
package app

import (
	"gin-blog/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}
