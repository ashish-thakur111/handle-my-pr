package osHelper

import "runtime"

var Windows = "windows"
var Linux = "linux"

func GetOperatingSystem() string {
	return runtime.GOOS
}
