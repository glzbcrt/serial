# Serial Loopback Detector

This is a simple hardware and software solution I created to detect where a serial device is connected. 

We had this issue in our company to detect in what serial port on Windows a device was connected. Usually the computers we found in the field had several serial devices
and to find the correct serial port was painful. People usually selected the incorrect serial port.

We already tried to detect where the device is by sending some bytes and checking the return value, but as we do not know what each device might return it doesn't work sometimes.
So the solution was a physical one, disconnect the device we want to detect the serial port, plug our serial loopback and let our software detect our device instead of the original one.

If you have a problem like this follow this simple three steps.

## Step 1 - Create the serial loopback device
This is the hardware part.

What we need is a device that sends back what it receives. It was easy because the serial interface is so old and, in this case, old means simple.
To create this device we need:

    1. Serial connector. It can be a male or female, 9 or 25 pins. Adapt it to your requirements
    2. An welding soldering iron
    3. Some wire for soldering


![DB-9 Diagram](doc/db9-diagram.png)

Weld the pin 2 (TX) to the pin 3 (RX). This way any data that leaves the serial port from TX will return to the computer in RX.
We also need to weld the ground pin 5 (GND) to the computer. Simply weld it to the connector, and that's it.


## Step 2 - Write and read from each serial port
This is the software part.

In this part we need a piece of code to iterate over the serial ports writing some bytes. After we write we read it back.
If what we read is the same as what we wrote, it means we found our serial loopback.

I wrote a sample application in Go. You can find the source code in this repository file serial_windows.go.


## Step 3 - Run the solution
Now the final part.

Connect your serial loopback device in the serial port we want to detect and run your program.
You should see the following output: 

```
D:\projects\serial>serial.exe
loopback found on: COM5
```

It means the serial loopback was found on COM5 in this case.

Hope it helps you!