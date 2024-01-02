#!/bin/bash

HOME_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

echo ###################################################################
echo Dbeast toolkit home folder: $HOME_DIR
echo ###################################################################
echo #


if which java > /dev/null 2>&1; then
    java -version
else
    echo "Java not installed"
    exit 1
fi

java -Xmx1g -Xms1g -Dlog4j2.configurationFile=$HOME_DIR/config/log4j2.xml -cp $HOME_DIR/lib/*:$HOME_DIR/bin/* co.dbeast.dbeast_toolkit.runner.DbeastToolkit -c $HOME_DIR