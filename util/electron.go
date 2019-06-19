package util

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/pkg/errors"

)

// SendMessage sends a message
func SendMessage(name string, window *astilectron.Window, payload interface{}, cs ...bootstrap.CallbackMessage) error {
	if window == nil {
		return errors.New("global window not init")
	}
	return bootstrap.SendMessage(window, name, payload, cs...)
}