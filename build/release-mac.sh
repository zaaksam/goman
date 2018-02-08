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
    DISTNAME=$NAME.$VERSION.$TYPE-mac
    TARNAME=$DISTNAME.tar.gz

    echo
    echo $TARNAME building...
    
    APPDIR=$DISTDIR/$NAME.app/Contents

    mkdir -p $APPDIR/{MacOS,Resources}

    cat > $APPDIR/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>goman</string>
    <key>CFBundleGetInfoString</key>
    <string>goman</string>
    <key>CFBundleIconFile</key>
    <string>goman</string>
    <key>CFBundleIdentifier</key>
    <string>com.naivesound.goman</string>
    <key>CFBundleName</key>
    <string>goman</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
</dict>
</plist>
EOF
    cp $DIR/$NAME.icns $APPDIR/Resources/$NAME.icns

    go build -o $APPDIR/MacOS/$NAME $SRC
    upx $APPDIR/MacOS/$NAME

    cd $DISTDIR
    tar czvf $TARNAME $NAME.app
    rm -r $NAME.app
    
    echo
}

function USAGE()
{
    echo
    echo argument error.
    echo
    echo e.g. \"./release-mac.sh v0.1.0 web\" or \"./release-mac.sh v0.1.0 app\"
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