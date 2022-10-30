# gin-server

repo for learning golang and gin

this repo build a server for simple mock API

| stack | repo                                                                 | desc             |
| ----- | -------------------------------------------------------------------- | ---------------- |
| go    | [https://golang.google.cn/](https://golang.google.cn/)               |                  |
| gin   | [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) | go web framework |
| air   | [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)   | go hmr           |


request in vscode     
plugin: [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)     

## develop      
go version 1.19 

```shell
# first: install air
go install github.com/cosmtrek/air@latest
# second: add air to env e.g. alias in .zshrc

# .zshrc
...
alias air='$(go env GOPATH)/bin/air'
...

```

## debug      
in the root directory, run:      
```shell
air
```

