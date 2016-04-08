package main

import (
	"fmt"

	"flag"
	"net/http"
	"strings"

	"github.com/coreos/fleet/unit"
	"github.com/kr/pretty"
)

var (
	rootURI = flag.String("rootURI", "", "Base uri to use when constructing service file URI. Only used if service file URI is relative.")
	sdc     httpServiceDefinitionClient
)

type services struct {
	Services []service `yaml:"services"`
}

type service struct {
	Name                 string `yaml:"name"`
	Version              string `yaml:"version"`
	Count                int    `yaml:"count"`
	URI                  string `yaml:"uri"`
	DesiredState         string `yaml:"desiredState"`
	SequentialDeployment bool   `yaml:"sequentialDeployment"`
}

func main() {
	flag.Parse()
	sdc = httpServiceDefinitionClient{httpClient: &http.Client{}, rootURI: *rootURI}
	servicesDefinition, _ := sdc.servicesDefinition()
	fmt.Printf("Services: [%# v]\n", pretty.Formatter(servicesDefinition))

	for _, srv := range servicesDefinition.Services {
		serviceTemplate, _ := sdc.serviceFile(srv)
		fmt.Printf("Service Template: \n[%s]\n-------------\n", string(serviceTemplate))

		vars := make(map[string]interface{})
		vars["version"] = srv.Version
		serviceFile, _ := renderedServiceFile(serviceTemplate, vars)
		fmt.Printf("Renderred Service File: \n[%s]\n-------\n", serviceFile)

		uf, _ := unit.NewUnitFile(serviceFile)
		fmt.Printf("Unit File: \n[%# v]\n------------------\n", pretty.Formatter(uf))
	}
}

func renderedServiceFile(serviceTemplate []byte, context map[string]interface{}) (string, error) {
	if context["version"] == "" {
		return string(serviceTemplate), nil
	}
	versionString := fmt.Sprintf("DOCKER_APP_VERSION=%s", context["version"])
	serviceTemplateString := strings.Replace(string(serviceTemplate), "DOCKER_APP_VERSION=latest", versionString, 1)
	return serviceTemplateString, nil
}
