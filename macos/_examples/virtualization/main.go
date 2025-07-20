package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/pkg/term/termios"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/virtualization"
	"golang.org/x/sys/unix"
)

func setRawMode(f *os.File) {
	var attr unix.Termios

	// Get settings for terminal
	termios.Tcgetattr(f.Fd(), &attr)

	// Put stdin into raw mode, disabling local echo, input canonicalization,
	// and CR-NL mapping.
	attr.Iflag &^= syscall.ICRNL
	attr.Lflag &^= syscall.ICANON | syscall.ECHO

	// Set minimum characters when reading = 1 char
	attr.Cc[syscall.VMIN] = 1

	// set timeout when reading as non-canonical mode
	attr.Cc[syscall.VTIME] = 0

	// reflects the changed settings
	termios.Tcsetattr(f.Fd(), termios.TCSANOW, &attr)
}

func serialConsole() virtualization.VirtioConsoleDeviceSerialPortConfiguration {
	console := virtualization.NewVirtioConsoleDeviceSerialPortConfiguration()

	stdin := os.Stdin
	stdout := os.Stdout

	// // Put stdin into raw mode, disabling local echo, input canonicalization,
	// // and CR-NL mapping.
	setRawMode(stdin)

	stdioAttachment := virtualization.NewFileHandleSerialPortAttachment()
	stdioAttachment.InitWithFileHandleForReadingFileHandleForWriting(
		foundation.NewFileHandleWithFileDescriptor(int(stdin.Fd())),
		foundation.NewFileHandleWithFileDescriptor(int(stdout.Fd())),
	)

	console.SetAttachment(stdioAttachment)
	return console
}

func bootLoader(kernelURL, initrdURL foundation.URL) virtualization.IBootLoader {
	bootLoader := virtualization.NewLinuxBootLoaderWithKernelURL(kernelURL)
	bootLoader.SetInitialRamdiskURL(initrdURL)
	kernelCommandLineArguments := []string{
		"console=hvc0",
	}

	bootLoader.SetCommandLine(strings.Join(kernelCommandLineArguments, " "))
	return bootLoader
}

func main() {
	configuration := virtualization.NewVirtualMachineConfiguration()
	configuration.SetCPUCount(1)
	configuration.SetMemorySize(2 * 1024 * 1024 * 1024)
	configuration.SetSerialPorts([]virtualization.ISerialPortConfiguration{serialConsole()})
	configuration.SetBootLoader(bootLoader(
		foundation.NewURLFileURLWithPath("./macos/_examples/virtualmachine/vmlinuz"),
		foundation.NewURLFileURLWithPath("./macos/_examples/virtualmachine/initrd.img"),
	))

	var nserr foundation.Error
	if !configuration.ValidateWithError(unsafe.Pointer(&nserr)) {
		fmt.Printf("Failed to validate the virtual machine configuration. %v\n", foundation.ToGoError(nserr))
		os.Exit(1)
	}

	virtualMachine := virtualization.NewVirtualMachineWithConfiguration(configuration)
	delegate := new(virtualization.VirtualMachineDelegate)
	virtualMachine.SetDelegate(delegate)

	virtualMachine.StartWithCompletionHandler(func(nserr foundation.Error) {
		if !nserr.IsNil() {
			fmt.Printf("Failed to start the virtual machine. %v", foundation.ToGoError(nserr))
		}
	})

	foundation.RunLoop_MainRunLoop().RunUntilDate(foundation.Date_DistantFuture())
}
