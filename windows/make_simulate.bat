@echo off

set VERSION=v.0.1.0
echo version=%VERSION%

set GOPATH_BAK=%GOPATH%
set GOPATH=%GOPATH%;%~dp0;%~dp0..;

set GOOS=windows
set GOARCH=amd64


echo start install simulate ...
go install -gcflags " -N -l" simulate

echo make ok

set GOPATH=%GOPATH_BAK%