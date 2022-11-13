package config

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInvalidConfig   = errors.New("invalid config")
	ErrInvalidSettings = errors.New("invalid settings")
)

type ErrInvalidAction struct {
	action interface{}
}

func NewErrInvalidAction(action interface{}) ErrInvalidAction {
	return ErrInvalidAction{action}
}

func (err ErrInvalidAction) Error() string {
	j, _ := json.Marshal(err.action)
	return fmt.Sprintf("invalid action %s", string(j))
}
