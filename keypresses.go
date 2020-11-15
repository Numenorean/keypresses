package keypresses

import (
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

var (
	mod 					= windows.NewLazyDLL("user32.dll")
	procGetClassNameW 		= mod.NewProc("GetClassNameW")
	procGetAsyncKeyState 		= mod.NewProc("GetAsyncKeyState")
	classtolook = []string{
		"ConsoleWindowClass",
	}
)

type (
	HANDLE uintptr
	HWND HANDLE
)

func IsKeyPressed(keyVirtualCode int) bool {
	// Query key mapped to integer `0x00` to `0xFF` if it's pressed.
	asynch, _, _ := procGetAsyncKeyState.Call(uintptr(keyVirtualCode))
	
	// If the least significant bit is set ignore it.
	//
	// As it's written in the documentation:
	// `if the least significant bit is set, the key was pressed after the previous call to GetAsyncKeyState.`
	// Which we don't care about :)
	if asynch&0x1 == 0 {
		return false
	}
	
	return true
}

func getClassName(hwnd HWND) (name string, err error) {
	n := make([]uint16, 256)
	p := &n[0]
	r0, _, e1 := syscall.Syscall(procGetClassNameW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(p)), uintptr(len(n)))
	if r0 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
		return
	}
	name = syscall.UTF16ToString(n)
	return
}

func getWindow(funcName string) uintptr {
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func checkInArray(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func IsKeyPressedGlobal(keyVirtualCode int, global bool) bool {
	if global {
		return IsKeyPressed(keyVirtualCode)
	}
	if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
			cn , _ := getClassName(HWND(hwnd))
			if checkInArray(cn ,classtolook) {
				return IsKeyPressed(keyVirtualCode)
			return false
		}
	}
	return false
}
