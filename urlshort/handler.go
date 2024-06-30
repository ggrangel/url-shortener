package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	return MapHandler(parsedYaml, fallback), nil
}

type PathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYAML(yamlData []byte) (map[string]string, error) {
	pathUrl := []PathURL{}
	err := yaml.Unmarshal(yamlData, &pathUrl)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	for _, pu := range pathUrl {
		m[pu.Path] = pu.URL
	}
	return m, nil
}
