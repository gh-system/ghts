@echo off

IF EXIST C:\Go\bin\go.exe (
    SET GOROOT=C:\Go
) ELSE IF EXIST %ProgramFiles%\Go\bin\go.exe (
    SET GOROOT=%ProgramFiles%\Go
)

IF NOT DEFINED GOPATH (
    SET GOPATH=%USERPROFILE%\Go
)

SET CGO_ENABLED=1
SET GOARCH=386
SET PATH=%GOROOT%\bin;C:\msys64\mingw32\bin;C:\msys64\usr\bin

cd %GOPATH%\src\github.com\ghts\ghts\experimental\xing\c32
go run xing_C32.go
