package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstallRegistryForm_SubmitBothEmpty_Skips(t *testing.T) {
	form := NewInstallRegistryForm("ghcr.io/org/app:latest")

	registryFormPressTab(&form)
	registryFormPressTab(&form)
	form, cmd := form.Update(keyPressMsg("enter"))
	require.NotNil(t, cmd)

	msg := cmd()
	skip, ok := msg.(InstallRegistrySkipMsg)
	require.True(t, ok, "expected InstallRegistrySkipMsg, got %T", msg)
	assert.Equal(t, "ghcr.io/org/app:latest", skip.ImageRef)
}

func TestInstallRegistryForm_SubmitWithCredentials(t *testing.T) {
	form := NewInstallRegistryForm("ghcr.io/org/app:latest")

	registryFormTypeText(&form, "myuser")
	registryFormPressTab(&form)
	registryFormTypeText(&form, "mytoken")
	registryFormPressTab(&form)
	form, cmd := form.Update(keyPressMsg("enter"))
	require.NotNil(t, cmd)

	msg := cmd()
	submit, ok := msg.(InstallRegistrySubmitMsg)
	require.True(t, ok, "expected InstallRegistrySubmitMsg, got %T", msg)
	assert.Equal(t, "ghcr.io/org/app:latest", submit.ImageRef)
	assert.Equal(t, "myuser", submit.Registry.Username)
	assert.Equal(t, "mytoken", submit.Registry.Password)
}

func TestInstallRegistryForm_UsernameOnly_ShowsError(t *testing.T) {
	form := NewInstallRegistryForm("ghcr.io/org/app:latest")

	registryFormTypeText(&form, "myuser")
	registryFormPressTab(&form)
	registryFormPressTab(&form)
	form, cmd := form.Update(keyPressMsg("enter"))
	assert.Nil(t, cmd)
	assert.True(t, form.form.HasError())
	assert.Equal(t, "Password is required when username is set", form.form.Error())
}

func TestInstallRegistryForm_PasswordOnly_ShowsError(t *testing.T) {
	form := NewInstallRegistryForm("ghcr.io/org/app:latest")

	registryFormPressTab(&form)
	registryFormTypeText(&form, "mytoken")
	registryFormPressTab(&form)
	form, cmd := form.Update(keyPressMsg("enter"))
	assert.Nil(t, cmd)
	assert.True(t, form.form.HasError())
	assert.Equal(t, "Username is required when password is set", form.form.Error())
}

func TestInstallRegistryForm_Cancel(t *testing.T) {
	form := NewInstallRegistryForm("ghcr.io/org/app:latest")

	registryFormPressTab(&form)
	registryFormPressTab(&form)
	registryFormPressTab(&form)
	form, cmd := form.Update(keyPressMsg("enter"))
	require.NotNil(t, cmd)

	msg := cmd()
	_, ok := msg.(InstallRegistryBackMsg)
	assert.True(t, ok, "expected InstallRegistryBackMsg, got %T", msg)
}

// Helpers

func registryFormTypeText(form *InstallRegistryForm, text string) {
	for _, r := range text {
		*form, _ = form.Update(keyPressMsg(string(r)))
	}
}

func registryFormPressTab(form *InstallRegistryForm) {
	*form, _ = form.Update(keyPressMsg("tab"))
}
