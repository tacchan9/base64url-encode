# Overview

指定されたファイルをbase64urlエンコードします

# Installation

go version 1.2 以上 をサポートしています。

インストールしていない場合 [go install](https://golang.org/doc/install)からインストールしてください

インストールした後は、以下のコマンドから、インストールできます。
   
```bash
$ go get github.com/tacchan9/base64url-encode
```

##  CLIとしてのInstallation
pathを通して installコマンドを利用すると、cliとして利用できます。
```bash
export PATH=$PATH:$GOPATH/bin
```

* install
```bash
$ go install
```

# Examples

```bash
$ base64url-encode -f test.jpg
```

## Global Options
```
--file value, -f value
       を指定してください。
```

```
--version, -v                
        print the version
```
