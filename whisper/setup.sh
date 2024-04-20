#!/usr/bin/bash
SCRIPT_DIR=$(cd $(dirname $BASH_SOURCE[0]) > /dev/null; pwd)
cd $SCRIPT_DIR

if [ ! -d whisper.cpp ]; then
  echo "git clone https://github.com/ggerganov/whisper.cpp.git"
  git clone https://github.com/ggerganov/whisper.cpp.git
  if [ $? -ne 0 ]; then
    echo "git clone failed..."
    exit 1
  fi
fi

cd whisper.cpp

if [ ! -f models/ggml-base.en.bin ]; then
  echo "Download a whisper model"
  bash ./models/download-ggml-model.sh base.en
  if [ $? -eq 0 ]; then
    echo "model download failed"
    exit 1
  fi
fi

make

./main -m models/ggml-base.en.bin -f samples/jfk.wav


exit

  make base.en

####
exit

make

./main -f samples/jfk.wav

make base.en
