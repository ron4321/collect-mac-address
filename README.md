# collect-mac-address

`collect-mac-address` is simple MAC address collection tool for IT administrators.
This tool provides an executable file by golang cross-platform build.

## Quick start
If you installed Go, you can use `go get`. 
```
go get github.com/ron4321/collect-mac-address
``` 
and run `collect-mac-address`.
```
/Users/hoge% collect-mac-address
------------------------------
- Hostname
xxxMacBook-Pro.local

- MAC Addresses
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
ab:cd:ef:gh:00:11
------------------------------
Please Enter key to exit.
```

## Build executable file
If you want to distribute this tool, you have two build method.

### use Docker
`./build-with-docker.sh`

### use local Go
`./build.sh`

If your build is successful, you can get the following output.
```
bin/
├── darwin386
│   └── collect-mac-address
├── darwin64
│   └── collect-mac-address
├── linux386
│   └── collect-mac-address
├── linux64
│   └── collect-mac-address
├── windows386
│   └── collect-mac-address.exe
└── windows64
    └── collect-mac-address.exe
```