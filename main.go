package main

import (
	"crossair/utils"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

func main() {
	hInstance := win.GetModuleHandle(nil)
	className := syscall.StringToUTF16Ptr("OverlayWindow")

	wc := win.WNDCLASSEX{
		CbSize: uint32(unsafe.Sizeof(win.WNDCLASSEX{})),
		LpfnWndProc: syscall.NewCallback(func(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
			switch msg {
			case win.WM_PAINT:
				utils.DrawCrosshair(hwnd)
				utils.ValidateRect(hwnd, nil)
				return 0
			case win.WM_DESTROY:
				win.PostQuitMessage(0)
				return 0
			default:
				return win.DefWindowProc(hwnd, msg, wParam, lParam)
			}
		}),
		HInstance:     hInstance,
		LpszClassName: className,
	}

	win.RegisterClassEx(&wc)

	hwnd := win.CreateWindowEx(
		win.WS_EX_TOPMOST|utils.WS_EX_LAYERED|utils.WS_EX_TRANSPARENT,
		className,
		syscall.StringToUTF16Ptr("Crosshair"),
		win.WS_POPUP,
		0, 0, 1920, 1080,
		0, 0, hInstance, nil,
	)

	utils.SetTransparent(hwnd)
	win.ShowWindow(hwnd, win.SW_SHOW)
	win.UpdateWindow(hwnd)

	var msg win.MSG
	for {
		if win.GetMessage(&msg, 0, 0, 0) > 0 {
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		} else {
			break
		}
	}
}
