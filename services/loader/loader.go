package loader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	log "github.com/mdeheij/logwrap"
	"github.com/mdeheij/monitoring/configuration"
	"github.com/mdeheij/monitoring/services/model"
)

func readFile(path string) []byte {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic("Cannot read file: ", path, err.Error())
	}
	return raw
}

func readServiceFiles(searchDir string) []model.Service {
	findingList := []string{}
	var allServices []model.Service
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		findingList = append(findingList, path)
		if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
			log.Notice("Loaded ", path)
			allServices = append(allServices, servicesBuilder(path)...)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return allServices
}

func servicesBuilder(path string) []model.Service {
	var fileServices []model.Service
	err := yaml.Unmarshal(readFile(path), &fileServices)
	if err != nil {
		panic(err)
	}
	return fileServices
}

func FindServices() []model.Service {
	services := readServiceFiles(configuration.C.Paths.Services)
	if len(services) < 1 {
		log.Panic("No services found! This makes me useless! Panic!")
	}
	return services
}
