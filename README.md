# Datatable
Datatable is a simple tool to create a wide variety of binary datatypes, which then can be integrated seamlessly into various platforms.
The main target is GO, but by compiling to JSON, you can easliy create a translation to your language of choice

## Get Started
[Get Started Guide](https://github.com/lucakrueger/Datatable/wiki/Get-Started)

## Goals
- Create new data formats or port exisiting ones in a robust and easy manner
- Create integration in notime

## Install
1. Download latest build
2. Add to PATH (optional)
3. Test Installation by running: `dtm version`

## How to use
1. Create a datatable file (*.dt)
2. Run: `dtm input output target (optional) std (optional)`

## Targets
Currently only GO and JSON are supported

## Building
1. Download source code
2. Extract files
3. Run: `go build -o bin/ ./dtm/dtm.go`
