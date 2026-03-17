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
		username := f.TextField(0).Value()
		password := f.TextField(1).Value()

		if username == "" && password == "" {
			return func() tea.Msg { return InstallRegistrySkipMsg{ImageRef: imageRef} }
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
