package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// user errors
	ErrUserNameEmpty = &Errno{Code: 20101, Message: "User name is empty."}
	ErrPasswordEmpty = &Errno{Code: 20102, Message: "Password is empty."}
	ErrUserNotFound  = &Errno{Code: 20103, Message: "The user was not found."}
)
