package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dig_yaml "github.com/esakat/dig-yaml"
	"gopkg.in/yaml.v2"
)

func main() {

	var chartName = os.Getenv("CHART")

	chartPath := fmt.Sprintf("charts/%s/Chart.yaml", chartName)
	valuePath := fmt.Sprintf("charts/%s/values.yaml", chartName)

	f, err := os.Open(chartPath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	dec := yaml.NewDecoder(f)
	var y map[interface{}]interface{}
	dec.Decode(&y)

	chartVersion, err := dig_yaml.DigYaml(y, "appVersion")
	if err != nil {
		log.Fatal(err)
	}

	chartRelease, err := dig_yaml.DigYaml(y, "version")
	if err != nil {
		log.Fatal(err)
	}

	f, err = os.Open(valuePath)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	dec = yaml.NewDecoder(f)
	dec.Decode(&y)

	appVersion, err2 := dig_yaml.DigYaml(y, "common", "image", "tag")
	if err2 != nil {
		log.Fatal(err2)
	}

	if appVersion != chartVersion {
		fmt.Printf("✖️ Versions in 'Chart.yaml'(%s) and 'values.yaml'(%s) are differents.\n", chartVersion, appVersion)
	} else {
		fmt.Println("✔️ Versions are identicals.")
		os.Exit(0)
	}

	v := strings.Split(chartRelease.(string), ".")
	x, _ := strconv.ParseInt((v[len(v)-1]), 10, 0)
	x += 1
	v = v[:len(v)-1]
	v = append(v, fmt.Sprint("", x))

	newChartRelease := strings.Join(v, ".")

	var sedHelp = fmt.Sprintf(`
If you want to replace the version in Chart.yaml file, please execute this command:
  
  sed -i 's/%s/%s/' -i %s
  sed -i 's/%s/%s/' -i %s
  git add %s
  git commit -m "%s: update version to match docker image tag"
    `, chartVersion, appVersion, chartPath, chartRelease, newChartRelease, chartPath, chartPath, chartName)

	fmt.Println(sedHelp)
}
