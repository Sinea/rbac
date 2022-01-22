package http

type CreatePermissionRequest struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
