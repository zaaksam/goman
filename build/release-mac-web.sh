#!/bin/sh

NAME=goman
VERSION=v0.2.0

DIR=$(cd $(dirname $0)/..; pwd)
DISTNAME=goman.$VERSION.web-darwin64
DISTDIR=$DIR/build

cd $DIR
go build -o $DISTDIR/$DISTNAME github.com/zaaksam/goman/go
# upx $DISTDIR/$DISTNAME