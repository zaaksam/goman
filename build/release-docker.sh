#!/bin/sh

# chmod 777 release-docker.sh

NAME=zaaksam/goman
VERSION=$1

function BUILD()
{
    echo
    echo $NAME:latest and $NAME:$VERSION building...
    echo

    cd ..
    
    docker build -t $NAME:latest -t $NAME:$VERSION -f ./build/Dockerfile .

    docker rmi $(docker images -f "dangling=true" -q)
    
    echo
}

function USAGE()
{
    echo
    echo argument error.
    echo
    echo e.g. \"./release-docker.sh 0.1.0\"
    echo
}

if [ "$VERSION" == "" ];then
    USAGE;
else
    BUILD;
fi