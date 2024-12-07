package utils

import (
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

var (
	user32                         = syscall.NewLazyDLL("user32.dll")
	procSetLayeredWindowAttributes = user32.NewProc("SetLayeredWindowAttributes")
	procValidateRect               = user32.NewProc("ValidateRect")
)

const (
	LWA_COLORKEY      = 0x01
	LWA_ALPHA         = 0x02
	WS_EX_LAYERED     = 0x80000
	WS_EX_TRANSPARENT = 0x20
)

func SetTransparent(hwnd win.HWND) {
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|WS_EX_LAYERED)
	procSetLayeredWindowAttributes.Call(uintptr(hwnd), 0, 0xFF, LWA_COLORKEY|LWA_ALPHA)
}

func DrawCrosshair(hwnd win.HWND, color win.COLORREF) {
	// Get device context
	hdc := win.GetDC(hwnd)
	defer win.ReleaseDC(hwnd, hdc)

	screenWidth := win.GetSystemMetrics(win.SM_CXSCREEN)
	screenHeight := win.GetSystemMetrics(win.SM_CYSCREEN)

	centerX := (screenWidth / 2) + 1
	centerY := (screenHeight / 2) - 1

	logBrush := win.LOGBRUSH{
		LbStyle: win.BS_SOLID,
		LbColor: color,
	}

	pen := win.ExtCreatePen(win.PS_GEOMETRIC|win.PS_SOLID, 2, &logBrush, 0, nil)
	defer win.DeleteObject(win.HGDIOBJ(pen))

	// Select the pen into the device context
	oldPen := win.SelectObject(hdc, win.HGDIOBJ(pen))
	defer win.SelectObject(hdc, oldPen)

	// Draw horizontal and vertical lines for the crosshair
	win.MoveToEx(hdc, int(centerX-10), int(centerY), nil)
	win.LineTo(hdc, int32(centerX+10), int32(centerY))
	win.MoveToEx(hdc, int(centerX), int(centerY-10), nil)
	win.LineTo(hdc, int32(centerX), int32(centerY+10))
}

func ValidateRect(hwnd win.HWND, rect *win.RECT) bool {
	ret, _, _ := procValidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
	)
	return ret != 0
}
