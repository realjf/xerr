// #############################################################################
// # File: xerr.go                                                             #
// # Project: xerr                                                             #
// # Created Date: 2024/11/21 18:10:27                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2024/11/21 18:56:57                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package xerr

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type XErr struct {
	Code     int    `json:"code"`
	Msg      string `json:"message"`
	Cause    error  `json:"cause"`
	Occurred string `json:"occurred"`
}

func New() *XErr {
	return &XErr{
		Occurred: generateOccurred(2),
	}
}

func NewCode(code int) *XErr {
	return &XErr{
		Occurred: generateOccurred(2),
		Code:     code,
	}
}

func WrapError(err error) *XErr {
	return &XErr{
		Occurred: generateOccurred(2),
		Code:     0,
		Cause:    err,
	}
}

func NewError(msg string) *XErr {
	return &XErr{
		Occurred: generateOccurred(2),
		Code:     0,
		Cause:    fmt.Errorf(msg),
	}
}

func generateOccurred(skip int) string {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("hh", ok)
		return ""
	}
	return fmt.Sprintf("%s, %s, line: %d", file, runtime.FuncForPC(pc).Name(), line)
}

func (x *XErr) Error() string {
	if x == nil {
		return ""
	}

	formattedErr := struct {
		Code     int    `json:"code"`
		Msg      string `json:"message"`
		Cause    string `json:"cause"`
		Occurred string `json:"occurred"`
	}{
		Code:     x.Code,
		Msg:      x.Msg,
		Occurred: x.Occurred,
	}

	if x.Cause != nil {
		formattedErr.Cause = x.Cause.Error()
	}
	errByte, _ := json.Marshal(formattedErr)
	return string(errByte)
}
