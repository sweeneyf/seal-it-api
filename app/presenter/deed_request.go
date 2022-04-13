package presenter

type DeedRequest struct {
	Data     string `json:"data" binding:"required"`
	DataType string `json:"dataType" binding:"required"`
	Action   string `json:"action" binding:"required"`
}
