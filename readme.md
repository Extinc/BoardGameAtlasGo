# Beggining with Go
![Golang](https://img.shields.io/badge/GOLANG-1.19-brightgreen)

Golang workshop using BoardGameAtlas with Event hosted by GeeksHacking   

https://www.boardgameatlas.com/
https://www.boardgameatlas.com/api/docs

## Setup


### Windows 
https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10 

### OSX
https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos 

### Linux

https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04 


## Some useful commands
```
go init mod 
go run . 
go build . 

Cross compile - compile to other platform
go tool dist list
GOOS=darwin GOARCH=arm64 go build -o LearningGoLang-darwin-arm64 .
GOOS=darwin GOARCH=amd64 go build -o LearningGoLang-windows-amd64 .
```

## Additional information

| Do not have unused import / variable | lowercase == private | uppercase == public  |
| ------------------------------------ | -------------------- | -------------------- |


| Library                |
| ---------------------- |
| github.com/fatih/color |
