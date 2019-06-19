package menu

import (
	"encoding/json"
	"fmt"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
	"lgtm/util"
)

type Menu struct {
	window *astilectron.Window
	appName *string
	builtAt *string
}

func New() *Menu {
	return &Menu{}
}

func (m *Menu)SetWindow(w *astilectron.Window) *Menu {
	m.window = w
	return m
}

func (m *Menu)SetAppName(appName *string) *Menu {
	m.appName = appName
	return m
}

func (m *Menu)SetbuiltAt(builtAt *string) *Menu {
	m.builtAt = builtAt
	return m
}

func (m *Menu)OnAbout(e astilectron.Event) (deleteListener bool) {
	htmlAbout := fmt.Sprintf(`
Welcome on <b>%s</b>!<br>
This is using the go-astilectron as backend and vue.js as frontend.<br>
BuiltAt: %s`,
		*m.appName, *m.builtAt)
	if err := util.SendMessage("menu.about", m.window, htmlAbout, func(m *bootstrap.MessageIn) {
		// Unmarshal payload
		var s string
		if err := json.Unmarshal(m.Payload, &s); err != nil {
			astilog.Error(errors.Wrap(err, "unmarshaling payload failed"))
			return
		}
		astilog.Infof("About modal has been displayed and payload is %s!", s)
	}); err != nil {
		astilog.Error(errors.Wrap(err, "sending about event failed"))
	}
	return
}