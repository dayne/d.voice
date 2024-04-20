#!/usr/bin/bash
SCRIPT_DIR=$(cd $(dirname $BASH_SOURCE[0]) > /dev/null; pwd)
cd $SCRIPT_DIR

echo "no.. just ignore this for now..."
exit

if [ ! -d willow-application-server ]; then
  git clone https://github.com/toverainc/willow-application-server.git 
  if [ $? -ne 0 ]; then
    echo "Failed: git clone willow-application-server"
    exit 1
  fi
fi

cd willow-application-server

if [ ! -f utils.sh ]; then
  echo "Missing: utils.sh"
  exit 1
fi

./utils.sh build
