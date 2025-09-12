# SimpleGoFileServer
A very simple program written in go for hosting static files.

## Usage
To start the server you either need the go file of the server along with go installed or the compiled binary file.


### Basic syntax:
```console
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
To access the files in the “files” directory, assuming the server runs on 192.168.1.10 on all interfaces, you can either download them from your terminal with the cURL command like this:

```console
curl http://192.168.1.10:8000/Text.txt
```
Or visit it in your browser:
```console
http://192.168.1.10:8000/
```


If a server listens on all interfaces that means it is available network-wide and maybe even over the internet if port-forwarding is active.
If a server listens only on localhost, that means only your local machine can access it.
