package constant

import (
	"fmt"
)

type ParamError struct {
	Param string
}

func (e ParamError) Error() string {
	return fmt.Sprintf("Missing or invalid %s parameters", e.Param)
}

func NewParamError(p string) ParamError {
	return ParamError{Param: p}
}
