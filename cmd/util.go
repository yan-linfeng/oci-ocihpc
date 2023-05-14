// This software is licensed under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/oracle-quickstart/oci-ocihpc/stacks"
	"github.com/oracle/oci-go-sdk/example/helpers"
)

var filename = ".stackinfo.json"

func addStackInfo(s Stack) {
	var stackInfo map[string]Stack
	if _, err := os.Stat(filename); err != nil {
		stackInfo = map[string]Stack{clusterName: s}
	} else {
		content, err := ioutil.ReadFile(filename)
		helpers.FatalIfError(err)
		json.Unmarshal([]byte(content), &stackInfo)
		stackInfo[clusterName] = s
	}

	file, _ := json.MarshalIndent(stackInfo, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func getSourceStackName() string {

	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	return info[clusterName].SourceStackName
}

func getDeployedStackName() string {

	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	return info[clusterName].DeployedStackName
}

func getStackID() string {

	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	return info[clusterName].StackID
}

func getStackIP() string {

	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	return info[clusterName].StackIP
}

func getJobID() string {

	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	return info[clusterName].JobID
}

func getWd() string {
	dir, err := os.Getwd()
	helpers.FatalIfError(err)

	return dir
}

func downloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	helpers.FatalIfError(err)

	defer resp.Body.Close()

	out, err := os.Create(filepath)
	helpers.FatalIfError(err)

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func copyFile(src fs.File, dest string) error {

	out, err := os.Create(dest)
	helpers.FatalIfError(err)

	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func getRandomNumber(n int) string {
	numbers := []rune("0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func getOutputQuery(data string, query string) string {

	var p map[string]interface{}
	json.Unmarshal([]byte(data), &p)
	q := p["outputs"].(map[string]interface{})[query].(map[string]interface{})["value"]
	str := fmt.Sprint(q)
	return str
}

func getConfirmation(prompt string) bool {
	var response string

	fmt.Printf("\n%s (y/n): ", prompt)
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println(prompt)
		return getConfirmation(prompt)
	}
}

func getStackQuery(stack string, value string) string {

	stackQuery, err := stacks.ConfigFS.ReadFile("stackQuery.json")
	helpers.FatalIfError(err)

	m := make(map[string]interface{})
	err = json.Unmarshal(stackQuery, &m)
	if err != nil {
		panic("Failed to decoding JSON data")
	}

	if localStackConfigPath != "" {
		localStackConfigContent, err := os.ReadFile(localStackConfigPath)
		helpers.FatalIfError(err)

		localStackConfig := make(map[string]interface{})
		err = json.Unmarshal(localStackConfigContent, &localStackConfig)
		if err != nil {
			panic("Failed to decoding JSON data")
		}

		for key, value := range localStackConfig {
			m[key] = value
		}
	}

	r := m[stack].(map[string]interface{})[value]
	str := fmt.Sprint(r)
	return (str)
}

func removeStackInfo() {
	content, err := ioutil.ReadFile(filename)
	helpers.FatalIfError(err)

	var info map[string]Stack
	json.Unmarshal([]byte(content), &info)

	delete(info, clusterName)

	stackJson, err := json.Marshal(info)
	helpers.FatalIfError(err)

	err = os.WriteFile(filename, stackJson, 0644)
	helpers.FatalIfError(err)
}
