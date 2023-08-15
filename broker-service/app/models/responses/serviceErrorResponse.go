package responses

type ServiceErrorResponse struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
