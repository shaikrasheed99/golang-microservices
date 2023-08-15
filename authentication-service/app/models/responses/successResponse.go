package responses

type SuccessResponse struct {
	Status  string      `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
