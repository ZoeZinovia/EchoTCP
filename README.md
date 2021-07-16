# EchoTCP

Simple client and server side applications for echo protocol over TCP

## Description

This project implements simple client and server side applications for the echo protocol. The server application listens for incoming requests on port 7 (as recommended by RFC 862). The server responds by sending back any data received in the request. The client application reads input from the console, connects to port 7 and sends an echo request to the server. It then outputs the response to the console.

See [here](https://datatracker.ietf.org/doc/html/rfc862) for more information on the Echo protocol.

## Getting Started

### Dependencies

There are only a few requirements for this project.

* Go must be [installed](https://golang.org/dl/)
* Go ROUTE and Go PATH must be configured according to your preferences. See [here](https://www.geeksforgeeks.org/golang-gopath-and-goroot/) for more information.
* There are a few other Go library imports, e.g. io, that do not require additional installations. 

### Installing and using

* Simply clone the code from this repository. It is recommended to run the code in your terminal and two separate terminal windows are needed: one for the client side and one for the server side.
* Once the terminal windows are in the correct directories, you can compile the files by running ```go build echoServer.go``` and ```go build echoClient.go``` respectively.
* Then the code can be run. Start by running the server side application: ```./echoServer```
* Next, the client side application can be run: ```./echoClient```
* The program will then wait for an input. You can type the message that you would like echoed. After typing the full message, signify the end of your input with the end of transmission signal: Ctrl-D. The echoed response will then appear. Below is an example:
```
>> ./echoClient
>> Hello world!
>> Hello world!
```

