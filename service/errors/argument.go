package errors

type ArgumentsError struct {
	message string
}

const invalidArgumentsMessage = "Invalid arguments"

func NewArgumentsError() ArgumentsError {
	return ArgumentsError{
		message: invalidArgumentsMessage,
	}
}

func (e ArgumentsError) Error() string {
	return e.message
}
