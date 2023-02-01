package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func JSONResponses(code int, message string, data interface{}) Response {
	jsonResponse := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return jsonResponse
}
