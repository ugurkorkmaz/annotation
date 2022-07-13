package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
	"text/template"
)

type Config struct {
	Directory string // Handler or Controller directory path
	Package   string // Output package name
	Mode      string // Output mode (gofiber, gin, echo, etc.)
	Output    string // Output file path; routes.go
}

type Route struct {
	Method  string // HTTP method
	Path    string // Path
	Handler string // Handler name
	Name    string // Route name
}
type App struct {
	Config
	Routes []Route
}

func (a *App) SetConfig(c Config) {
	a.Config = c
}

func (a *App) Parse() {
	fset := token.NewFileSet()
	node, err := parser.ParseDir(fset, a.Config.Directory, nil, parser.ParseComments)

	if err != nil {
		log.Fatal(err)
	}
	for _, v := range node {
		for _, f := range v.Files {
			for _, d := range f.Decls {
				switch d.(type) {
				case *ast.FuncDecl:
					pa, err := a.ParseAnnotations(d.(*ast.FuncDecl).Doc.Text())
					if err != nil {
						log.Fatal(err)
					}
					pa.Handler = d.(*ast.FuncDecl).Name.Name
					a.Routes = append(a.Routes, pa)

				}
			}
		}
	}

}
func (a *App) ParseAnnotations(route string) (Route, error) {
	regex := regexp.MustCompile(`@Route\(\"(.*?)\", name=\"(.*?)\", methods=\{\"(.*?)\"\}\)`)
	match := regex.FindStringSubmatch(route)

	return Route{
		Method: match[3],
		Path:   match[1],
		Name:   match[2],
	}, nil
}
func (a *App) ParseHandlerDirName() string {
	// ./example/dir -> example/dir
	return strings.Replace(a.Config.Directory, "./", "", 1)
}
func (a *App) Generate() {

	temp, _ := template.ParseFiles("template/" + a.Config.Mode + ".tmpl")

	f, _ := os.Create(a.Config.Output)
	i, _ := debug.ReadBuildInfo()
	temp.Execute(f, map[string]interface{}{
		"Routes":  a.Routes,
		"Package": a.Config.Package,
		"Module":  i.Main.Path + "/" + a.ParseHandlerDirName(),
	})
}
