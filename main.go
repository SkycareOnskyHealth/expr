package expr

import (
	"bytes"

	"text/template"
)

// ParseText parse template
func ParseText(name string, body string, data interface{}) (string, error) {
	_template := template.New(name)
	_templateParse, err := _template.Parse(body)
	if err != nil {
		return "", err
	}
	var t = template.Must(_templateParse, nil)
	buf := bytes.Buffer{}
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
