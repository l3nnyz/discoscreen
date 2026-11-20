@echo off
go build -ldflags="-s -w" -trimpath -o discoscreen.exe
