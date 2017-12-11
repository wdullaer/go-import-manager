# Go Import Manager
Go Import Manager is a small cli tool which allows you to manipulate the import statements of a go file. Rather than relying on regex text replacements, it will parse and manipulate the AST of a file to do its job.
This is useful if you want to automate the build of an application that uses a [caddy](https://caddyserver.com) like plugin system.

The latest binaries can be downloaded from the Github release pages.

## Usage
The application uses a docker style cli, with subcommands

### List
```bash
go-import-manager list main.go
```

### Add
```bash
go-import-manager add main.go "fmt"
```

### Delete
```bash
go-import-manager delete main.go "fmt"
```

### Replace
```bash
go-import-mananager replace main.go "fmt" "github.com/foo/my-fmt"
```

## Build
```bash
go get "github.com/wdullaer/go-import-manager"
go get "github.com/urfave/cli"
cd $GOPATH/src/github.com/wdullaer/go-import-manager
go build
```

## TODO
* Add tests
* Add automated builds for more platforms
* Add options to specify the output

## License
Apache-2.0 © [Wouter Dullaert](https://wdullaer.com)
