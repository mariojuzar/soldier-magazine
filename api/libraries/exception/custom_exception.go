package exception

import "errors"


func NotExistException() error {
	return errors.New("data not exist")
}
