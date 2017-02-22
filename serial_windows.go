package main

/*
	Code specific to Windows.
*/

import (
	"fmt"

	"github.com/jacobsa/go-serial/serial"
)

func findSerialLoopback(magicString string) (string, bool) {

	// The initial and final port number we will be checking.
	startPort := 1
	finishPort := 20

	bytesToWrite := []byte(magicString)

	for i := startPort; i <= finishPort; i++ {

		portName := fmt.Sprintf("COM%d", i)

		options := serial.OpenOptions{
			PortName:        portName,
			BaudRate:        19200,
			DataBits:        8,
			StopBits:        1,
			MinimumReadSize: 4,
		}

		port, err := serial.Open(options)

		// If we were unable to open the port, skip this port.
		if err != nil {
			continue
		}

		n, err := port.Write(bytesToWrite)

		// If we were unable to write the magic string, skip this port.
		if err != nil || n != len(bytesToWrite) {
			port.Close()
			continue
		}

		bytesRead := make([]byte, len(magicString))
		n, err = port.Read(bytesRead)

		// If we were unable to read some bytes from the port, skip this port.
		if err != nil {
			port.Close()
			continue
		}

		// If what we read is the same as what we wrote, it means we found our serial loopback.
		if string(bytesRead[:n]) == magicString {
			port.Close()
			return portName, true
		}

		port.Close()
	}

	// If we got here, it means we were unable to find the serial loopback.
	return "", false

}

func main() {

	if port, found := findSerialLoopback("glzbcrt"); found {
		fmt.Println("loopback found on: " + port)
	} else {
		fmt.Println("loopback not found.")
	}

}
