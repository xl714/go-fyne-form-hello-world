# go hello world

```bash
go mod init github.com/xl714/go-fyne-ui

go get fyne.io/fyne/v2@latest
```

Execute the following command to run the code:

```bash
go run main.go
```

Formatting: For better readability, use the go fmt command to automatically format your code:

```bash
go fmt main.go
```

Building a Binary: To create a standalone executable file, use the go build command:

```bash
go build -o hello main.go
```

This will create an executable file named hello in your project directory. You can run it directly without the go run command.

Using Modules: For larger projects, consider using Go modules for dependency management.
