package adb

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/funnywwh/go-shell"
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
	AdbName string
}

type Shell struct {
	ExitStatus uint32
	Stdout     *bytes.Buffer
	Stderr     *bytes.Buffer
	cmd        *exec.Cmd
	err        error
}

func Run(cmd string, args ...string) *Shell {
	sh := new(Shell)
	sh.Stderr = bytes.NewBuffer(nil)
	sh.Stdout = bytes.NewBuffer(nil)
	sh.cmd = exec.Command(cmd, args...)
	sh.cmd.Stderr = sh.Stderr
	sh.cmd.Stdout = sh.Stdout
	err := sh.cmd.Run()
	if err != nil {
		sh.err = err
	}
	fmt.Printf("%s err:%v\n", sh.cmd.Args, err)
	return sh
}
func (this *Shell) Error() (err error) {
	err = this.err
	return
}
func (this *Adb) Shell(cmd string, args ...string) (stdout string, err error) {
	var _args []string
	_args = append(_args, Adb_Shell, cmd)
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	stdout = p.Stdout.String()
	return
}

func (this *Adb) Install(apk string, args ...string) (err error) {
	var _args []string
	_args = append(_args, Adb_Install)
	_args = append(_args, args...)
	_args = append(_args, apk)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Broadcast(action string, args ...string) (err error) {
	var _args []string
	_args = append(_args, Adb_Shell, "am", "broadcast", "-a", action)
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Start(activity string, args ...string) (err error) {
	var _args []string
	_args = append(_args, Adb_Shell, "am", "start", activity)
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	if p.Stderr.Len() > 0 {
		err = p.Error()
	}
	return
}
func (this *Adb) Forward(args ...string) (err error) {
	var _args []string
	_args = append(_args, "forward")
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Push(localPath, remotePath string) (err error) {
	var _args []string
	_args = append(_args, Adb_Push, localPath, remotePath)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}
