Install Go:
```bash
sudo tar -C $HOME -xzf <filename>.tar.gz
```

Edit the .bashrc file (nano ~/.bashrc), add the following lines:

```bash
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```
