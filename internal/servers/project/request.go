package project

type ProjectRequest struct {
	Title       string `json:"Title"`
	Description string `json:"Description,omitempty"`
	Logo        string `json:"Logo,omitempty"`
}

type ProjectResponse struct {
	Title       string `json:"Title"`
	Description string `json:"Description,omitempty"`
	Logo        string `json:"Logo,omitempty"`
}
