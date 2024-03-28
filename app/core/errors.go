package core

type Error struct {
    Message string `json:"message"`
    Code int `json:"code"`
}

func NewError(code int, message string) *Error {
    return &Error {
        Message: message,
        Code: code,
    }
}
