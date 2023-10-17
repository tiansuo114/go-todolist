package handler

import "errors"

func PanicIfUserErr(err error) {
	if err != nil {
		err = errors.New("User Service -- " + err.Error())
	}
}
