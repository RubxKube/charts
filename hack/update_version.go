package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	dig_yaml "github.com/esakat/dig-yaml"
	"gopkg.in/yaml.v2"
)

func contains(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}

func searchAndReplace(path, wordToFind, wordToReplace string) error {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(fileData), wordToFind, wordToReplace)

	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

// This function is used to update the version of the chart in the Chart.yaml and values.yaml files
func updateChartVersion(chartName string, isGitHubAction bool) error {
	chartPath := fmt.Sprintf("../charts/%s/Chart.yaml", chartName)
	valuePath := fmt.Sprintf("../charts/%s/values.yaml", chartName)

	f, err := os.Open(chartPath)
	defer f.Close()

	if err != nil {
		log.Fatal("Error when trying to read " + "../charts/" + chartName + "/Chart.yaml")
		return err
	}

	dec := yaml.NewDecoder(f)
	var y map[interface{}]interface{}
	dec.Decode(&y)

	chartVersion, err := dig_yaml.DigYaml(y, "appVersion")
	if err != nil {
		log.Fatal("Error when trying to read " + "../charts/" + chartName + "/Chart.yaml")
		log.Fatal(err)
	}
	chartVersion = fmt.Sprintf("%v", chartVersion)

	chartRelease, err := dig_yaml.DigYaml(y, "version")
	if err != nil {
		fmt.Println("Can't read chart release")
		log.Fatal(err)
	}
	chartRelease = fmt.Sprintf("%v", chartRelease)

	f, err = os.Open(valuePath)
	defer f.Close()
	if err != nil {
		fmt.Println("Can't read value file")
		log.Fatal(err)
	}

	dec = yaml.NewDecoder(f)
	dec.Decode(&y)

	appVersion, err := dig_yaml.DigYaml(y, "common", "image", "tag")
	if err != nil {
		fmt.Println("Can't read tag")
		return err
	}
	appVersion = fmt.Sprintf("%v", appVersion)

	if appVersion == chartVersion {
		fmt.Println("✔️ Versions are identicals.")
	} else {
		fmt.Printf("✖️ Versions in 'Chart.yaml'(%s) and 'values.yaml'(%s) are differents.\n", chartVersion, appVersion)

		v := strings.Split(chartRelease.(string), ".")
		x, _ := strconv.ParseInt((v[len(v)-1]), 10, 0)
		x += 1
		v = v[:len(v)-1]
		v = append(v, fmt.Sprint("", x))

		newChartRelease := strings.Join(v, ".")

		if !isGitHubAction {
			var sedHelp = fmt.Sprintf(`
If you want to replace the version in Chart.yaml file, please execute this command:
		  
sed -i 's/%s/%s/' -i %s
sed -i 's/%s/%s/' -i %s
git add %s
git commit -m "%s: update version to match docker image tag"
			`, chartVersion, appVersion, chartPath, chartRelease, newChartRelease, chartPath, chartPath, chartName)

			fmt.Println(sedHelp)

			var replaceFiles string
			fmt.Print("Or do you want me to replace the version in files? (yes/no): ")
			fmt.Scan(&replaceFiles)

			if replaceFiles == "yes" || replaceFiles == "y" || replaceFiles == "" {
				fmt.Println("Let's replace !")
				err3 := searchAndReplace(chartPath, chartVersion.(string), appVersion.(string))
				if err3 != nil {
					log.Fatal(err3)
				}
				err4 := searchAndReplace(chartPath, chartRelease.(string), newChartRelease)
				if err4 != nil {
					log.Fatal(err4)
				}
			} else {
				fmt.Println("Let's not replace")
			}
		} else {
			fmt.Println("Script is running in a GitHub Action")
			err3 := searchAndReplace(chartPath, chartVersion.(string), appVersion.(string))
			if err3 != nil {
				log.Fatal(err3)
			}
			err4 := searchAndReplace(chartPath, chartRelease.(string), newChartRelease)
			if err4 != nil {
				log.Fatal(err4)
			}
		}
	}
	return nil

}
func main() {

	isGitHubAction := os.Getenv("GITHUB_ACTIONS") == "true"
	chartNames := make([]string, 0)

	if isGitHubAction {
		fmt.Println("Script is running in a GitHub Action")
		chartPaths := os.Args[1:]

		for _, path := range chartPaths {
			parts := strings.Split(path, "/")
			chartName := strings.TrimPrefix(parts[1], "charts/")
			// If the chart name is not in the list, add it
			if !contains(chartNames, chartName) {
				chartNames = append(chartNames, chartName)
			}
		}

	} else {
		fmt.Println("Script is not running in a GitHub Action")
		var chartName string
		fmt.Print("Which chart to update ? ")
		fmt.Scan(&chartName)
		chartNames = append(chartNames, chartName)
	}
	fmt.Println(chartNames)

	for _, chartName := range chartNames {
		fmt.Println("Updating " + chartName + "...")
		err := updateChartVersion(chartName, isGitHubAction)
		if err != nil {
			fmt.Sprintf("Error when trying to update %s", chartName)
			continue
		}

	}

}
