#!/bin/sh

NAME=goman
VERSION=v0.2.0

DIR=$(cd $(dirname $0)/..; pwd)
DISTNAME=$NAME.$VERSION.app-darwin64
DISTDIR=$DIR/build

cd $DIR
mkdir -p $DISTDIR/$NAME.app/Contents/{MacOS,Resources}
go build -o $DISTDIR/$NAME.app/Contents/MacOS/$NAME github.com/zaaksam/goman/go
# upx $DISTDIR/$NAME.app/Contents/MacOS/$NAME

cat > $DISTDIR/$NAME.app/Contents/Info.plist << EOF
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
cp $DISTDIR/$NAME.icns $DISTDIR/$NAME.app/Contents/Resources/$NAME.icns