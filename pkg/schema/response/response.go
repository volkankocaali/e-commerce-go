package response

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Count      int         `json:"count"`
	Data       interface{} `json:"data"`
	Error      interface{} `json:"error"`
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}, count int) Response {

	return Response{
		StatusCode: statusCode,
		Message:    message,
		Count:      count,
		Data:       data,
		Error:      err,
	}

}
