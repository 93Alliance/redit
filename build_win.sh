#! /bin/bash
cat > redit.rc << EOL
id ICON "./res/app_win.ico"
GLFW_ICON ICON "./res/app_win.ico"
EOL

x86_64-w64-mingw32-windres redit.rc -O coff -o redit.syso
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 go build -ldflags "-s -w -H=windowsgui -extldflags=-static" -p 4 -v -o redit.exe
rm redit.syso
rm redit.rc