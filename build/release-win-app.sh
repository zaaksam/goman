#!/bin/sh

NAME=goman
VERSION=v0.2.0

DISTNAME=$NAME.$VERSION.app-win64

xgo -ldflags="-H windowsgui" -out $DISTNAME --targets=windows-6.1/amd64 ../go
mv $DISTNAME-windows-6.1-amd64.exe $DISTNAME.exe
# upx $DISTNAME.exe