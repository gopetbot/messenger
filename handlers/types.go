package handlers


type Response struct {
	Message string `json:"message"`
}

type DialogFlowRequest struct {
	OriginalRequest OriginalRequest `json:"originalRequest"`
}

type OriginalRequest struct {
	Source string `json:"source"`
}