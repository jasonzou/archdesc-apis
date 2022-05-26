# Wails - golang Desktop 

## Installation on Fedora

- prerequisite packages
```
$ sudo dnf install gcc-c++ make pkgconf-pkg-config webkit2gtk3-devel gtk3-devel
```
- install wails - go version 1.17+
```
$ go install github.com/wailsapp/wails/cmd/wails@latest
```

## First project

- Create a project
```
wails init
```
- Build the project - node v17 with errors and downgrade to v16, then everything is fine.
```
wails build
```