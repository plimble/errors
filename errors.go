package errors

const (
	Internal     = 500
	BadReq       = 400
	NotFound     = 404
	Unauthorized = 401
	Forbidden    = 403
)

type Errors struct {
	HttpStatus int    `json:"-"`
	ErrMessage string `json:"message,omitempty"`
	ErrType    string `json:"error,omitempty"`
	ErrCode    string `json:"code,omitempty"`
	ErrDevMsg  string `json:"dev_message,omitempty"`
}

func (err *Errors) Error() string {
	return err.ErrMessage
}

func New(msg string) *Errors {
	return &Errors{
		ErrMessage: msg,
	}
}

func (err *Errors) Http(i int) *Errors {
	err.HttpStatus = i

	return err
}

func (err *Errors) Type(s string) *Errors {
	err.ErrType = s

	return err
}

func (err *Errors) Code(s string) *Errors {
	err.ErrCode = s

	return err
}

func (err *Errors) DevMsg(s string) *Errors {
	err.ErrDevMsg = s

	return err
}

func IsStatus(status int, err error) bool {
	if errs, ok := err.(*Errors); ok {
		if status == errs.HttpStatus {
			return true
		}
	}

	return false
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
