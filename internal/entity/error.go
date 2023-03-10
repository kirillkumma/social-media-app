package entity

type ErrCode int

const (
	_ = ErrCode(iota)
	ErrCodeBadInput
	ErrCodeInternal
)

type Error struct {
	msg  string
	code ErrCode
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

func NewError(msg string, code ErrCode) *Error {
	return &Error{msg, code}
}
