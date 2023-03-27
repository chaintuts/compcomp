## General
____________

### Author
* Josh McIntyre

### Website
* jmcintyre.net

### Overview
* CompComp is a utility for comparing the size of a compressed and uncompressed file

## Development
________________

### Git Workflow
* master for releases (merge development)
* development for bugfixes and new features

### Building
* make build
Build the application - wraps `go build`
* make clean
Clean the build directory

### Features
* Fetch the size in bytes of two files (compressed and uncompressed)
* Compare the percentage difference in file sizes

### Requirements
* Requires Go language build tools

### Platforms
* Windows
* MacOSX
* Linux

## Usage
____________

### Command Line Usage
* Run `./compcomp <filename1> <filename2>`