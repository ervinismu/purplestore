package handler

type ResponeBody struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// NewRespBodyData new response body with data
func NewRespBodyData(status string, message string, data interface{}) *ResponeBody {
	return &ResponeBody{
		Status:   status,
		Message: message,
		Data:   data,
	}
}

func NewRespMessage(status string, message string) *ResponeBody {
	return &ResponeBody{
		Status:   status,
		Message: message,
	}
}
