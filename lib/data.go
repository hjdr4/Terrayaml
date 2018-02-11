package lib

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}

func LoadData(dirs []string) (string, error) {
	data := ""
	for _, dir := range dirs {
		var slash string
		if strings.HasSuffix(dir, "/") {
			slash = ""
		} else {
			slash = "/"
		}

		expanded, err := expand(dir)
		if err != nil {
			return "", err
		}

		stat, err := os.Stat(expanded)
		if err != nil {
			return "", err
		}

		if stat.IsDir() {
			files, _ := filepath.Glob(expanded + slash + "*.yml")
			for _, file := range files {
				bytes, err := ioutil.ReadFile(file)
				if err != nil {
					return "", err
				}
				data += string(bytes) + "\n"
			}
		} else {
			bytes, err := ioutil.ReadFile(expanded)
			if err != nil {
				return "", err
			}
			data += string(bytes) + "\n"
		}
	}
	return data, nil
}

func Convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = Convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = Convert(v)
		}
	}
	return i
}
