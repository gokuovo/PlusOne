package response

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data any) Response {
	return Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}
