package utils

import "github.com/lxn/win"

func GetColor(color string) win.COLORREF {
	switch color {
	case "blue":
		return win.RGB(0, 0, 255)
	case "red":
		return win.RGB(255, 0, 0)
	case "green":
		return win.RGB(0, 255, 0)
	case "yellow":
		return win.RGB(255, 255, 0)
	case "pink":
		return win.RGB(255, 192, 203)
	case "purple":
		return win.RGB(128, 0, 128)
	case "orange":
		return win.RGB(255, 165, 0)
	case "white":
		return win.RGB(255, 255, 255)
	case "black":
		return win.RGB(0, 0, 0)
	case "gray":
		return win.RGB(128, 128, 128)
	case "brown":
		return win.RGB(165, 42, 42)
	case "cyan":
		return win.RGB(0, 255, 255)
	case "magenta":
		return win.RGB(255, 0, 255)
	case "silver":
		return win.RGB(192, 192, 192)
	case "gold":
		return win.RGB(255, 215, 0)
	case "maroon":
		return win.RGB(128, 0, 0)
	case "olive":
		return win.RGB(128, 128, 0)
	case "lime":
		return win.RGB(0, 255, 0)
	case "teal":
		return win.RGB(0, 128, 128)
	case "navy":
		return win.RGB(0, 0, 128)
	case "sakura":
		return win.RGB(255, 212, 245)
	default:
		return win.RGB(0, 0, 0)
	}
}
