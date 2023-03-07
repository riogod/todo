package api_common

type ResponseErrorDTO struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type ResponseOK_DTO struct {
	Success bool        `json:"success"`
	Body    interface{} `json:"body,omitempty"`
}
type ResponseERROR_DTO struct {
	Success bool             `json:"success"`
	Error   ResponseErrorDTO `json:"error"`
}
