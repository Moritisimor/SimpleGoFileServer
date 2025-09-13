# Simple Go File Server
A very simple program written in go for hosting static files.

## How to compile
To compile this program you will need a Go installation on your PC, preferably the newest one. Then, you need to clone the repository:
```console
git clone https://github.com/Moritisimor/SimpleGoFileServer
```
Then, cd into the newly cloned repository's directory, where the source code resides:
```console
cd SimpleGoFileServer/SimpleGoFileServer
```
And finally compile it with the Go compiler:
```console
go build . 
```
Depending on your platform, you can then launch the Server by just double clicking the compiled binary file or typing the following in your terminal:
```console
./SimpleGoFileServer
```

## How to launch
To start the server you either need the go file of the server along with go installed or the compiled binary file.

By default, without any arguments given, the Server listens on all interfaces and on port 8000,
so you can theoretically just double click the .exe file (or whatever format your OS uses) and it would run until the terminal window is closed, or the process is terminated in some other way.

### Basic syntax:
```
SimpleGoFileServer <Port (Integer)> <Localhost (Boolean)>
```

(Note: Boolean means true or false.)
## Examples
### If you want to run the server on port 8000 and only on localhost, it would look like this:
```console
SimpleGoFileServer 8000 true
```

### If you want to run the server on port 7000 and listen on all interfaces it would look like this:
```console
SimpleGoFileServer 7000 false
```

## Error codes and their meanings
#### 1: The selected port is not a number.
#### 2: The selected port is not in the valid range of available ports.
#### 3: 2nd Argument got a value which is not true or false.
#### 4: Got more arguments than necessary.

## How to use
To make files available, you simply put them in the folder named "files" and they will be available there.

To access the files in the “files” directory, assuming the server runs on 192.168.1.10 on all interfaces, 
you can either download them from your terminal with the cURL command like this:

```console
curl http://192.168.1.10:8000/Text.txt
```
Or visit it in your browser:
```console
http://192.168.1.10:8000/
```

If a server listens on all interfaces that means it is available network-wide and maybe even over the internet if port-forwarding is active.
If a server listens only on localhost, that means only your local machine can access it.
