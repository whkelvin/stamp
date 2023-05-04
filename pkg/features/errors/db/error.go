package db

type DbError interface {
	Error() string
}

type Error struct {
	msg string
}

func (err *Error) Error() string {
	return err.msg
}

func New(msg string) *Error {
	return &Error{msg: msg}
}
