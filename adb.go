package adb

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"syscall"

	"github.com/funnywwh/go-shell"
)

var (
	Adb_Shell   = "shell"
	Adb_Install = "install"
	Adb_Push    = "push"
	Adb_Devices = "devices"
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
	AdbName      string
	DeviceSerial string
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
	sh.cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
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

var regDevices = regexp.MustCompile(`(\S+)\s+device\b`)

func (this *Adb) Devices() (devices []string) {
	p := Run(this.AdbName, Adb_Devices)
	err := p.Error()
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	stdout := p.Stdout.String()
	m := regDevices.FindAllStringSubmatch(stdout, -1)
	if len(m) == 0 {
		fmt.Printf("find none")
		return
	}
	for _, line := range m {
		devices = append(devices, line[1])
	}
	fmt.Printf("devices:%#v\n", devices)
	return
}
func (this *Adb) Shell(cmd string, args ...string) (stdout string, err error) {
	var _args []string
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}
	_args = append(_args, Adb_Shell, cmd)
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	stdout = p.Stdout.String()
	return
}

func (this *Adb) Install(apk string, args ...string) (err error) {
	var _args []string
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}

	_args = append(_args, Adb_Install)
	_args = append(_args, args...)
	_args = append(_args, apk)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Broadcast(action string, args ...string) (err error) {
	var _args []string
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}

	_args = append(_args, Adb_Shell, "am", "broadcast", "-a", action)
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Start(activity string, args ...string) (err error) {
	var _args []string
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}

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
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}

	_args = append(_args, "forward")
	_args = append(_args, args...)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}

func (this *Adb) Push(localPath, remotePath string) (err error) {
	var _args []string
	if len(this.DeviceSerial) > 0 {
		_args = append(_args, "-s", this.DeviceSerial)
	}

	_args = append(_args, Adb_Push, localPath, remotePath)
	p := Run(this.AdbName, _args...)
	err = p.Error()
	return
}
