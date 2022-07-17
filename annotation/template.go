package annotation

const fiber string = `
package {{.Package}}

import (
	"github.com/gofiber/fiber/v2"
    "{{.Module}}"
)

func routeLoad(http *fiber.App) {
	{{range .Routes}}
    http.Add("{{.Method}}", "{{.Path}}", example.{{.Handler}})
    {{end}}
}
`
