package dto

type SignTransactionRequestDto struct {
	Gate string      `form:"gate" json:"gate" xml:"gate" binding:"required"`
	Tx   interface{} `form:"tx" json:"tx,omitempty" xml:"tx"`
}
