# BachelorGo

## Requirement

- install [golang](https://golang.org/) >= 1.9
- install [go dep](https://github.com/golang/dep) >= 0.5

## Quick start
- [Set $GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)
- `cd $GOPATH/src/ && mkdir github.com`

- `cd github.com && git clone https://github.com/TLNBS/BachelorGo`

- `cd BachelorGo && dep ensure`

- `go run app.go`

- talk to the BachelorBot on `https://nguyenbachelorthesis.slack.com/`. Use `%new`to start a conversation 

## Further development

- If an another purpose than ArtConsultant is needed, follow these steps:
  
1. Create a struct, which implements the methods of responder/responder.go
2. Create an instance of that struct and use it as parameter in the method `newSlackBot` in app.go 
  