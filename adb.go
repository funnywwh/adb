package adb

import (
	"github.com/progrium/go-shell"
)

var (
	Adb_Shell   = "shell_hx"
	Adb_Install = "install_hx"
)

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
	err = p.Error()
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
	_args = append(_args, "adb", "am", "start", activity)
	_args = append(_args, args...)
	p := shell.Run(_args...)
	err = p.Error()
	return
}
