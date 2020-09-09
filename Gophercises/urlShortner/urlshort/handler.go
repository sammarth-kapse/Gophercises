package urlshort

import (
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
)

type Config struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusPermanentRedirect)
		} else{
			fallback.ServeHTTP(w,r)
		}
	}
}


func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]Config, error) {
	configArr := make([]Config, 0)
	err := yaml.Unmarshal(yml, &configArr)
	checkError(err)
	return configArr, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func buildMap(configArr []Config) map[string]string {
	m := make(map[string]string)
	for _, config := range configArr {
		//fmt.Println(config.Path)
		m[config.Path] = config.Url
	}
	return m
}