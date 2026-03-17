package ui

import (
	tea "charm.land/bubbletea/v2"

	"github.com/basecamp/once/internal/docker"
)

type InstallRegistrySubmitMsg struct {
	ImageRef string
	Registry docker.RegistrySettings
}
type InstallRegistryBackMsg struct{}
type InstallRegistrySkipMsg struct{ ImageRef string }

const (
	registryFormUsernameField = iota
	registryFormPasswordField
)

type InstallRegistryForm struct {
	form     Form
	imageRef string
}

func NewInstallRegistryForm(imageRef string) InstallRegistryForm {
	passwordField := NewTextField("password or access token")
	passwordField.SetEchoPassword()

	m := InstallRegistryForm{
		form: NewForm("Next",
			FormItem{Label: "Username", Field: NewTextField("username")},
			FormItem{Label: "Password", Field: passwordField},
		),
		imageRef: imageRef,
	}

	m.form.OnSubmit(func(f *Form) tea.Cmd {
		username := f.TextField(registryFormUsernameField).Value()
		password := f.TextField(registryFormPasswordField).Value()

		if username == "" && password == "" {
			return func() tea.Msg { return InstallRegistrySkipMsg{ImageRef: imageRef} }
		}

		if username == "" {
			f.errorField = registryFormUsernameField
			f.error = "Username is required when password is set"
			f.focused = registryFormUsernameField
			return nil
		}
		if password == "" {
			f.errorField = registryFormPasswordField
			f.error = "Password is required when username is set"
			f.focused = registryFormPasswordField
			return nil
		}

		return func() tea.Msg {
			return InstallRegistrySubmitMsg{
				ImageRef: imageRef,
				Registry: docker.RegistrySettings{
					Username: username,
					Password: password,
				},
			}
		}
	})
	m.form.OnCancel(func(f *Form) tea.Cmd {
		return func() tea.Msg { return InstallRegistryBackMsg{} }
	})

	return m
}

func (m InstallRegistryForm) Init() tea.Cmd {
	return m.form.Init()
}

func (m InstallRegistryForm) Update(msg tea.Msg) (InstallRegistryForm, tea.Cmd) {
	var cmd tea.Cmd
	m.form, cmd = m.form.Update(msg)
	return m, cmd
}

func (m InstallRegistryForm) View() string {
	return m.form.View()
}
