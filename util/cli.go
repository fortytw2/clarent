package util

const (
	MainHelpTemplate = `NAME:
{{.Name}} - {{.Usage}}
USAGE:
{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

VERSION:
{{.Version}}

COMMANDS:
{{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
{{end}}{{if .Flags}}
GLOBAL OPTIONS:
{{range .Flags}}{{.}}
{{end}}{{end}}`

	AppletHelpTemplate = `NAME:
{{.Name}} - {{.Usage}}
USAGE:
{{.Name}} {{if .Flags}}[global options] {{end}}command{{if .Flags}} [command options]{{end}} [arguments...]

COMMANDS:
{{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
{{end}}{{if .Flags}}
GLOBAL OPTIONS:
{{range .Flags}}{{.}}
{{end}}{{end}}`
)
