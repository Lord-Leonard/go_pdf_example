compile:
	echo "Compiling for Linus and Windiws"
	env GOOS=windows GOARCH=amd64 go build -o build/luna-reporting-windows.exe
	env GOOS=linux GOARCH=amd64 go build -o build/luna-reporting-linux