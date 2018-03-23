# Go Import Manager
Go Import Manager is a small cli tool which allows you to manipulate the import statements of a go file. Rather than relying on regex text replacements, it will parse and manipulate the AST of a file to do its job.
This is useful if you want to automate the build of an application that uses a [caddy](https://caddyserver.com) like plugin system.

The latest binaries can be downloaded from the Github release pages.

## Usage
The application uses a docker style cli, with subcommands

### List
**Options**
* *output, o*: The file to which the output will be written (default: stdout)

**Example**
```bash
go-import-manager list main.go
```

### Add
**Options**
* *output, o*: The file to which the output will be written (default: stdout)
* *inplace, i*: Set to edit the file in place

**Example**
```bash
go-import-manager add main.go "fmt"
```

### Delete
**Options**
* *output, o*: The file to which the output will be written (default: stdout)
* *inplace, i*: Set to edit the file in place

**Example**
```bash
go-import-manager delete main.go "fmt"
```

### Replace
**Options**
* *output, o*: The file to which the output will be written (default: stdout)
* *inplace, i*: Set to edit the file in place

**Example**
```bash
go-import-mananager replace main.go "fmt" "github.com/foo/my-fmt"
```

## Build
```bash
go get "github.com/wdullaer/go-import-manager"
```

## TODO
* Add tests
* Add automated builds for more platforms

## License
Apache-2.0 © [Wouter Dullaert](https://wdullaer.com)
