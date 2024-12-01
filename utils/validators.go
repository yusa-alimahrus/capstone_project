package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

// Response helpers
func ResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}

func ResponseSuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
