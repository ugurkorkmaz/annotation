package main

import (
	"github.com/alexflint/go-arg"
	"github.com/ugurkorkmaz/annotation/annotation"
)

func main() {
	var args struct {
		Directory string // Handler or Controller directory path
		Package   string // Output package name
		Mode      string // Output mode (gofiber, gin, echo, etc.)
		Output    string // Output file path; routes.go
	}
	arg.MustParse(&args)

	app := new(annotation.App)
	app.SetConfig(annotation.Config{
		Directory: args.Directory,
		Package:   args.Package,
		Mode:      args.Mode,
		Output:    args.Output,
	})
	app.Parse()
	app.Generate()
}
