@echo off

set NAME=goman
set VERSION=%1
set TYPE=%2
set SRC=github.com/zaaksam/goman/go/main/%TYPE%

if "%VERSION%" == "" GOTO USAGE
if "%TYPE%" == "web" GOTO BUILD
if "%TYPE%" == "app" GOTO BUILD
GOTO USAGE

:BUILD
set DIR=%cd%
set DISTDIR=%DIR%\dist
set DISTNAME=%NAME%.%VERSION%.%TYPE%-win64.exe

set DISTSYSO=%GOPATH%\src\%SRC:/=\%\%NAME%.syso

if not exist %DISTDIR% (
	mkdir %DISTDIR%
)

echo.
echo %DISTNAME% building...

windres -i %DIR%\%NAME%.rc -O coff -o %DISTSYSO%

if "%TYPE%" == "app" (
	go build -ldflags="-H windowsgui" -o %DISTDIR%\%DISTNAME% %SRC%
) else (
	go build -o %DISTDIR%\%DISTNAME% %SRC%
)

del %DISTSYSO%

:: upx %DISTNAME%

GOTO END

:USAGE
echo.
echo argument error.
echo.
echo e.g. "release-win v0.1.0 web" or "release-win v0.1.0 app"
echo.

:END