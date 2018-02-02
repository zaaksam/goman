#!/bin/sh

# chmod 777 release-mac.sh

NAME=goman
VERSION=$1
TYPE=$2
SRC=github.com/zaaksam/goman/go/main/$TYPE

function BUILD()
{
    DIR=$(cd $(dirname $0); pwd)
    DISTDIR=$DIR/dist
    DISTNAME=$NAME.$VERSION.$TYPE-win
    ZIPNAME=$DISTNAME.zip

    echo
    echo $ZIPNAME building...

    cd $DISTDIR

    # See github.com/karalabe/xgo
    if [ "$TYPE" == "app" ];then
        xgo -ldflags="-H windowsgui" -out $DISTNAME --targets=windows-6.1/amd64 $GOPATH/src/$SRC
    else
        xgo -out $DISTNAME --targets=windows-6.1/amd64 $GOPATH/src/$SRC
    fi
    
    mv $DISTNAME-windows-6.1-amd64.exe $DISTNAME.exe
    # upx $DISTNAME.exe
    
    zip -m $ZIPNAME $DISTNAME.exe
    
    echo
}

function USAGE()
{
    echo
    echo argument error.
    echo
    echo e.g. \"./release-win.sh v0.1.0 web\" or \"./release-win.sh v0.1.0 app\"
    echo
}

if [ "$VERSION" == "" ];then
    USAGE;
elif [ "$TYPE" == "web" ];then
    BUILD;
elif [ "$TYPE" == "app" ];then
    BUILD;
else
    USAGE;
fi