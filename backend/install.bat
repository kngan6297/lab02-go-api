@echo off
echo Installing backend dependencies...
go mod tidy
echo Backend dependencies installed successfully!
echo.
echo Don't forget to update the database password in config.env
echo To start the backend, run: go run main.go
pause 