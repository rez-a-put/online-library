package utils

func ErrorRetrieveData() string {
	return "Failed to get Books data"
}

func ErrorFailedReadRequest() string {
	return "Failed to read request data"
}

func ErrorWrongTimeRequest() string {
	return "Wrong time input"
}

func ErrorCantBeBeforeNow() string {
	return "Pickup time can't be before current time"
}

func ErrorRequired(param string) string {
	return "Please input " + param
}

func ErrorWrongUserPass() string {
	return "Wrong username or password"
}

func ErrorAlreadyPickedup() string {
	return "Book already has pickup schedule"
}

func ErrorEmptyData() string {
	return "There are no books for that genre"
}
