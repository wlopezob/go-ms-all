## Create Project for folder
go mod init github.com/wlopezob/go-firstendpoint

## Create module
go work init firstendpoint 

## run
go run .

## install gin
go get -u github.com/gin-gonic/gin

## install air
go install github.com/cosmtrek/air@latest
air -v
air init #add automatic .air.toml