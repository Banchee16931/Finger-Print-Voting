package cerr

type CommonError struct {
	Message  string `json:"message"`
	Metadata any    `json:"metadata,omitempty"`
}
