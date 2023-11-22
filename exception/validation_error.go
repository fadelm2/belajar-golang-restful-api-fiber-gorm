package exception

type ValidationError struct {
	Message string
	Field   string
	Tag     string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
