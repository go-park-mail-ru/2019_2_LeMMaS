package http

type Response struct {
	Status string      `json:"status"`
	Body   interface{} `json:"body"`
}

func Ok() Response {
	return Response{
		Status: "ok",
	}
}

func OkWithBody(body interface{}) Response {
	return Response{
		Status: "ok",
		Body:   body,
	}
}

type errorResponseBody struct {
	Message string `json:"message"`
}

func Error(err error) Response {
	return Response{
		Status: "error",
		Body:   errorResponseBody{Message: err.Error()},
	}
}
