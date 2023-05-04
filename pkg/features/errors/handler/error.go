package handler

type HandlerError interface {
	Error() string
	IsBadInput() bool
}

type Error struct {
	msg        string
	isBadInput bool
}

func (err *Error) Error() string {
	return err.msg
}

func (err *Error) IsBadInput() bool {
	return err.isBadInput
}

func New(msg string, isBadInput bool) *Error {
	return &Error{msg: msg, isBadInput: isBadInput}
}
