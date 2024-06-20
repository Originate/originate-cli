package templates

import "text/template"

var GoModTemplate = template.Must(template.New("go.mod.tmpl").Parse(`module github.com/{{.Organization}}/{{.Module}}

go 1.22.0

require (
	github.com/Originate/go-utilities v1.0.1
	github.com/gin-gonic/gin v1.9.1
	github.com/go-playground/validator/v10 v10.19.0
	github.com/swaggo/swag v1.16.3
)
`))
