# Go-API-Template

## 내부 패키지 구성하는 법

```sh
    ## 내부 패키지 등록
    cmd or internal path로 등록

    ex) go mod init cmd/hello
    ex) go mod init internal/net

    ## work
	go work init
	go work use [path]

    ex) go work use ./cmd/hello
    ex) go work use ./internal/net
```

## swagger 등록하는 법

- cmd/examples 참고
- Docekrfile.example.swagger 참고

```sh
    ## in repository
    go install github.com/swaggo/swag/cmd/swag@latest

    ## swag error
    zsh: command not found: swag

    ## swag error solution
    export PATH=$(go env GOPATH)/bin:$PATH

    ## fmt (formatting)
    swag fmt

    ## init (init)
    swag init

    ## import 등록 
    
    import (
        ...
        _ "cmd/{application}/docs
    )
```