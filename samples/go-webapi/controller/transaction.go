package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-webapi/dto"
)

// RegisterSignTransactionRoutes ...
func RegisterSignTransactionRoutes(router *gin.RouterGroup) {
	router.POST("/", signTransaction)
}

func signTransaction(c *gin.Context) {
	var json dto.SignTransactionRequestDto
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, dto.CreateBadRequestErrorDto(err))
		return
	}

	// TODO

	fmt.Println("sign tx")

	c.JSON(http.StatusOK, dto.CreateSuccessWithDtoAndMessageDto(json, []string{"TX was signed successfully"}))
}
