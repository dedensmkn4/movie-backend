package domain

type OmdbOutboundDto struct {
	Response	string	`json:"response,omitempty"`
	Error		string	`json:"error,omitempty"`
}
