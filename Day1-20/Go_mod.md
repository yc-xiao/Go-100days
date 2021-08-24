- go 项目由包组成
    - go build 编译包生成库文件，不可执行
    - 特殊包，package main 为可执行包，生成可执行文件
    - 同一个包内(包名一样)，多个文件下变量只能有一个，且包内变量名唯一。不同包直接，变量名可重复，仅允许调用大写变量

- go build / install
    - go build 编译文件，若是main包会生成可执行文件，默认在路径下(-o 可指定路径)
    - go install 编译并将文件移动到指定目录(GOROOT/GOPATH)下的bin(main包)/pkg(普通包)目录下
    - go build/install -a -x，执行并查看执行过程

- GOPATH/GOROOT/MOD
    - GOAPTH是GO默认的工作区(可多个)，目录下有src/bin/pkg三个目录。src是源代码，bin是可执行文件，pkg是二进制包。
        - 一般项目路径 src/gitlab.com/author/project
            - 在project 执行go install 会在bin目录下生成project可执行文件或者在pkg/gitlab.com/author/project生成project.a文件
    - GOROOT是GOROOT目录，类似GOPATH目录，也有src/bin/pkg目录。是GO安装的目录。
    - MOD是Go对包的管理工具
        - 在没有MOD时，项目只能写在$GOPATH/src目录下。
        - go mod init "gitlab.com/author/project"创建go.mod文件，go install 会在GOROOT/GOPATH的bin目录下创建project可执行文件，或者在某个文件下创建.a的库文件