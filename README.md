# Simple Go File Server
A simple program written in go for hosting static files.

## How to compile
To compile this program, you will need a Go installation on your PC, preferably the newest one. Then, you need to clone the repository:
```console
git clone https://github.com/Moritisimor/SimpleGoFileServer
```
Then, cd into the newly cloned repository's directory, where the source code resides:
```console
cd SimpleGoFileServer/SimpleGoFileServer
```
Since the program serves files in the 'files' folder by default, you should create such a folder like this:
```console
mkdir files
```
And finally compile it with the Go compiler:
```console
go build .
```
Depending on your platform, you can then launch the Server by just double-clicking the compiled binary file 
or typing the following in your terminal:
```console
./SimpleGoFileServer
```

## How to launch
To start the server, you either need the go file of the server along with go installed or the compiled binary file.

By default, without any arguments given, the Server listens on all interfaces on port 8000 and serves
files from the "files" directory.

It is possible to change the port, host and folder the server serves files from. The syntax is as follows:

### Basic syntax:
```
SimpleGoFileServer -p <Port Number> -h <Host> -d <Directory>
```

(Note: Boolean means true or false.)
## Examples
If you want to run the server on port 8000 and only on localhost, it would look like this:
```console
SimpleGoFileServer -p 8000 -h localhost
```

If you want to run the server on port 7000 and listen on all interfaces, it would look like this:
```console
SimpleGoFileServer -p 7000 -h 0.0.0.0
```

If you want to specify the directory where files are served from, as well, it would look like this:
```console
SimpleGoFileServer -p 7000 -h 0.0.0.0 -d exampleDirectory
```
File paths are relative to the directory where the server is launched from. Using absolute paths is not recommended.

## How to use
To share files, you put them in the folder named "files" and they will be available there.
Alternatively, you can also specify a different folder to serve files from. This option is the third argument.

To access the files in the “files” directory, assuming the server runs on 192.168.1.10 on all interfaces, 
you can either download them from your terminal with the cURL command like this:

```console
curl http://192.168.1.10:8000/Text.txt
```
Or visit it in your browser:
```console
http://192.168.1.10:8000/
```

If a server listens on all interfaces, that means it is available network-wide and maybe even over the internet if port-forwarding is active.

If a server listens only on localhost, that means only your local machine can access it.
