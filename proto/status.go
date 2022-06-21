package proto

// Response status constants.
const (
	StatusOK = iota

	StatusConnectionError
	StatusBadRequest
	StatusServiceFailure
	StatusBadResponse
)
