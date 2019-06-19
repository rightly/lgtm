package main

import (
	"github.com/asticode/go-astilectron"
	"lgtm/menu"
)

var (
	m = menu.New().
		SetWindow(gWindow).
		SetAppName(&AppName).
		SetbuiltAt(&BuiltAt)

	gMenuOptions = []*astilectron.MenuItemOptions{
		{
			Label: astilectron.PtrStr("파일"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label:   astilectron.PtrStr("정보"),
					OnClick: m.OnAbout,
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		},
		{
			Label: astilectron.PtrStr("메뉴"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label:   astilectron.PtrStr("정"),
					OnClick: m.OnAbout,
				},
				{Role: astilectron.MenuItemRoleClose},
			},
		},
	}
)