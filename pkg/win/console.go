package win

import (
	"syscall"
	"unsafe"

	"github.com/rs/zerolog/log"
)

// SetConsoleTitle sets windows console title
func SetConsoleTitle(title string) (int, error) {
	handle, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := syscall.FreeLibrary(handle); err != nil {
			log.Error().Err(err).Msg("Cannot free library")
		}
	}()

	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}

	rTitle, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return 0, err
	}

	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(rTitle)), 0, 0)
	return int(r), err
}
