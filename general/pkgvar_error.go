package general

import "errors"

var (
	ErrUnknown = errors.New("error unknown")
	ErrDbUnkown = errors.New("error in db")
	DbNotFound = errors.New("info of data not found")
	DbNotUpdata = errors.New("updata is ok, and don't updata any datas")
)