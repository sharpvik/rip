package proto

// Response status constants.
const (
	StatusOK = iota

	StatusConnectionError
	StatusBadRequest
	StatusServiceMalfunction
	StatusBadResponse
)
