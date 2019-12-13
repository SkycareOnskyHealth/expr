package errors

import (
	"fmt"
	"strings"

	"github.com/onskycloud/io"
)

// ParseText parse template
// func New(locale Locale, errorType ErrorType, errorCode Status, key string, messages interface{}) (string, error) {

// }
func (e Errors) detectLocale(locale Locale) ([]ErrorDictionary, error) {
	fileName := "en_us"
	if locale != ConvertLocale(fileName) {
		fileName = strings.ToLower(locale.String())
	}
	path := fmt.Sprintf("%s/%s.yaml", RootPath, fileName)
	file, err := io.ReadAll(path)
	if err != nil || file == "" {
		return err
	}
	return file
}

func LoadErrorList() ([]Errors, error) {
	errorDicts := []Errors{}
	files, err := io.ReadDir(RootPath)
	fmt.Printf("files : %+v\n", files)
	if err != nil {
		return errorDicts, err
	}
	for i := 0; i < len(files); i++ {
		file := &[]ErrorDictionary{}
		path := fmt.Sprintf("%s/%s", RootPath, files[i].Name)
		err := io.ReadYamlFile(path, file)
		if err != nil || file == nil {
			continue
		}
		locale := strings.Split(files[i].Name, ".")[0]
		errorDict := Errors{
			Locale:    ConvertLocale(locale),
			ErrorList: *file,
		}
		errorDicts = append(errorDicts, errorDict)
	}
	return errorDicts, nil
}
