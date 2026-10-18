package gui

import (
	"testing"

	"github.com/jesseduffield/lazygit/pkg/gocui"
	"github.com/stretchr/testify/assert"
)

func TestConfigureDiffSelection(t *testing.T) {
	gui := NewDummyGui()
	gui.Views.Main = &gocui.View{}
	gui.Views.Secondary = &gocui.View{}

	assert.True(t, gui.c.UserConfig().Gui.UseFullLineDiffSelection)
	gui.configureDiffSelection()
	assert.Equal(t, 0, gui.Views.Main.SelectedLineColorWidth)
	assert.Equal(t, 0, gui.Views.Secondary.SelectedLineColorWidth)

	gui.c.UserConfig().Gui.UseFullLineDiffSelection = false
	gui.configureDiffSelection()
	assert.Equal(t, 2, gui.Views.Main.SelectedLineColorWidth)
	assert.Equal(t, 2, gui.Views.Secondary.SelectedLineColorWidth)

	gui.c.UserConfig().Gui.UseFullLineDiffSelection = true
	gui.configureDiffSelection()
	assert.Equal(t, 0, gui.Views.Main.SelectedLineColorWidth)
	assert.Equal(t, 0, gui.Views.Secondary.SelectedLineColorWidth)
}
