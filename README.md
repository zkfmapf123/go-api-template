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