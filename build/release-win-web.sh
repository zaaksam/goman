#!/bin/sh

NAME=goman
VERSION=v0.2.0

DISTNAME=$NAME.$VERSION.web-win64
DISTDIR=$DIR/build

xgo -out $DISTNAME --targets=windows-6.0/amd64 ../go
mv $DISTNAME-windows-6.0-amd64.exe $DISTNAME.exe
# upx $DISTNAME.exe