package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ubozov/wallet-core/samples/go-webapi/crypto"
	"github.com/ubozov/wallet-core/samples/go-webapi/dto"
)

// RegisterSignTransactionRoutes ...
func RegisterSignTransactionRoutes(router *gin.RouterGroup) {
	router.POST("/", signTransaction)
}

func signTransaction(c *gin.Context) {
	var json dto.SignTransactionRequestDto
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, dto.CreateBadRequestErrorDto(err))
		return
	}

	fn, err := crypto.GetSigner(json.Gate)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, dto.CreateBadRequestErrorDto(err))
		return
	}

	seed, ok := c.Keys["Seed"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.CreateBadRequestErrorDto(fmt.Errorf("occured error when read seed")))
		return
	}

	result, err := fn(seed, json.Tx)

	if err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, dto.CreateBadRequestErrorDto(err))
		return
	}

	c.JSON(http.StatusOK, dto.CreateSuccessWithDtoAndMessageDto(result, []string{"TX was signed successfully"}))
}
