# Mimic3


## Tasks

### Depends
```
sudo apt-get install libespeak-ng1                                                  
```

### Create virtual environment & install
```
ENV_DIR=$HOME/.local/venv/mimic3
mkdir -p $ENV_DIR
python3 -m venv $ENV_DIR
source $ENV_DIR/bin/activate
pip3 install --upgrade pip                                                          
pip3 install mycroft-mimic3-tts[all]                                                
```
### Test-It
```
mimic3 'Hello world.' | aplay  
```
