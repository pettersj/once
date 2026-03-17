package ui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/basecamp/once/internal/docker"
)

const (
	registryUsernameField = iota
	registryPasswordField
)

type SettingsFormRegistry struct {
	settingsFormBase
}

func NewSettingsFormRegistry(settings docker.ApplicationSettings) SettingsFormRegistry {
	usernameField := NewTextField("username")
	usernameField.SetValue(settings.Registry.Username)

	passwordField := NewTextField("password or access token")
	passwordField.SetEchoPassword()
	passwordField.SetValue(settings.Registry.Password)

	m := SettingsFormRegistry{
		settingsFormBase: settingsFormBase{
			title: "Registry",
			form: NewForm("Done",
				FormItem{Label: "Username", Field: usernameField},
				FormItem{Label: "Password", Field: passwordField},
			),
		},
	}

	m.form.OnSubmit(func(f *Form) tea.Cmd {
		s := settings
		s.Registry.Username = f.TextField(registryUsernameField).Value()
		s.Registry.Password = f.TextField(registryPasswordField).Value()
		return func() tea.Msg { return SettingsSectionSubmitMsg{Settings: s} }
	})
	m.form.OnCancel(func(f *Form) tea.Cmd {
		return func() tea.Msg { return SettingsSectionCancelMsg{} }
	})

	return m
}

func (m SettingsFormRegistry) Update(msg tea.Msg) (SettingsSection, tea.Cmd) {
	var cmd tea.Cmd
	m.settingsFormBase, cmd = m.update(msg)
	return m, cmd
}
