package requests

type RequestPayload struct {
	Action string            `json:"action"`
	Signup AuthSignupPayload `json:"signup,omitempty"`
}
