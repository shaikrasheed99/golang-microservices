package requests

type AuthLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
