package main

import (
	"encoding/json"
	"fmt"
	"gojenkins"
)

const (
	JENKINS_ADDR  = "http://10.12.4.16"
	JENKINS_AUTH  = "chenling1"
	JENKINS_TOKEN = "113eb03fe3c5017ac4b182c038a00da3bc"
)

func main() {

	jenkins, _ := gojenkins.CreateJenkins(nil, JENKINS_ADDR, JENKINS_AUTH, JENKINS_TOKEN).Init()
	job, _ := jenkins.GetJob("go-pack", "DMD")

	pr, _ := job.GetPipelineRun("38")

	jsonstr, err := json.Marshal(pr)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonstr)

}
