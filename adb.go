package adb

import (
	"runtime"

	"github.com/progrium/go-shell"
)

var (
	Adb_Shell   = "shell"
	Adb_Install = "install"
	Adb_Push    = "push"
)

func init() {
	switch runtime.GOOS {
	case "linux", "android":
	case "windows":
		shell.Shell = []string{"cmd.exe", "/c"}
	}
	shell.Panic = false
}

type Adb struct {
}

func (this *Adb) Shell(cmd string, args ...interface{}) (stdout string, err error) {
	var _args []interface{}
	_args = append(_args, "adb", Adb_Shell)
	_args = append(_args, args...)
	p := shell.Run(_args...)
	err = p.Error()
	stdout = p.Stdout.String()
	return
}

func (this *Adb) Install(apk string, args ...interface{}) (err error) {
	var _args []interface{}
	_args = append(_args, "adb", Adb_Install)
	_args = append(_args, args...)
	_args = append(_args, apk)
	p := shell.Run(_args...)
	if p.ExitStatus != 0 {
		err = p.Error()
	}
	return
}

func (this *Adb) Broadcast(action string, args ...interface{}) (err error) {
	var _args []interface{}
	_args = append(_args, "adb", "am", "broadcast")
	_args = append(_args, args...)
	p := shell.Run(_args...)
	err = p.Error()
	return
}

func (this *Adb) Start(activity string, args ...interface{}) (err error) {
	var _args []interface{}
	_args = append(_args, "adb", Adb_Shell, "am", "start", activity)
	_args = append(_args, args...)
	p := shell.Run(_args...)
	err = p.Error()
	return
}
func (this *Adb) Forward(args ...interface{}) (err error) {
	var _args []interface{}
	_args = append(_args, "adb", "forward")
	_args = append(_args, args...)
	p := shell.Run(_args...)
	err = p.Error()
	return
}

func (this *Adb) Push(localPath, remotePath string) (err error) {
	var _args []interface{}
	_args = append(_args, "adb", Adb_Push, localPath, remotePath)
	p := shell.Run(_args...)
	err = p.Error()
	return
}
