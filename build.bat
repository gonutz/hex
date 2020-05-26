set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w"
if errorlevel 1 pause
