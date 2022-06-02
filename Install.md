# Instalación en Linux
1. Ir a la pag. oficial: [link](https://go.dev/doc/install)
2. Agregar variables de ambiente, en el archivo $HOME/.bashrc
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export GOROOT=/usr/local/go

export PATH=$PATH:$GOBIN:$GOROOT/bin
```

Luego para actualizar las vaibales de ambiente: ``` . .bashrc```
3. Crear entorno local: ```mkdir -p $HOME/go/{bin,pkg,src}```