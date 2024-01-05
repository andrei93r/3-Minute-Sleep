package winInteractions

import (
	"syscall"
	"unsafe"
)

const (
	SLEEP_IDLE           = "0"
	SLEEP_SLEEP          = "1"
	SLEEP_HIBERNATE      = "2"
	WAKE_EVENTS_ENABLED  = "0"
	WAKE_EVENTS_DISABLED = "1"
	NOT_CRITICAL         = "0"
	CRITICAL             = "1"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
	powrprof         = syscall.NewLazyDLL("powrprof.dll")
	setSuspendState  = powrprof.NewProc("SetSuspendState")
)

func SetSuspendState(forceSleep, critical string) error {
	disableWakeEvent := WAKE_EVENTS_DISABLED
	ret, _, err := setSuspendState.Call(
		uintptr(len(forceSleep)),
		uintptr(unsafe.Pointer(&forceSleep)),
		uintptr(len(disableWakeEvent)),
		uintptr(unsafe.Pointer(&disableWakeEvent)),
		uintptr(len(critical)),
		uintptr(unsafe.Pointer(&critical)),
	)
	if ret == 0 {
		return err
	}
	return nil
}

func IsDisplaying() (bool, error) {
	nrDisplays, err := getNumberOfDisplays()
	if err != nil {
		return false, err
	}

	if nrDisplays < 1 {
		return false, nil
	}

	return true, nil

}

func getNumberOfDisplays() (int, error) {
	numDisplays, _, _ := getSystemMetrics.Call(80)
	return int(numDisplays), nil
}
