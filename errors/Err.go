package errors

type StudentError struct {
	Message string
}

func (s StudentError) Error() string {
	return s.Message
}

type Taskerr struct {
	Message string
}

func (t Taskerr) Error() string {
	return t.Message
}
