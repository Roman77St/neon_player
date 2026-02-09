package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// NeonTheme — кастомная тема в стиле Neon Dark
type NeonTheme struct{}

func (m *NeonTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return ColorBackground
	case theme.ColorNamePrimary:
		return ColorNeonMain
	case theme.ColorNameButton:
		return ColorButtonBg
	case theme.ColorNameForeground:
		return ColorNeonMain
	}
	return theme.DefaultTheme().Color(name, theme.VariantDark)
}

func (m *NeonTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m *NeonTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m *NeonTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
