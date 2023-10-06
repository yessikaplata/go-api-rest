package errors

type UserNotFoundError struct {
	message string
}

const userNotFoundErrorMessage = "User not exists"

func NewUserNotFoundError() UserNotFoundError {
	return UserNotFoundError{
		message: userNotFoundErrorMessage,
	}
}

func (e UserNotFoundError) Error() string {
	return e.message
}
