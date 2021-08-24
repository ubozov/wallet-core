package dto

type SignTransactionRequestDto struct {
	Gate string `form:"gate" json:"gate" xml:"gate" binding:"required"`
	Tx   string `form:"tx" json:"tx" xml:"tx" binding:"required"`
}
