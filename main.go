package main

import (
	"fmt"
	"encoding/json"
	"github.com/guxingke/alfred-jenkins-go/jenkins"
	"os"
)

func main() {

	// 1. parse args
	if len(os.Args) < 1 {
		fmt.Println(emptyMsg())
		return
	}

	pArgs := os.Args[1:]

	url := pArgs[0]
	var query = ""
	if len(pArgs) > 1 {
		query = pArgs[1]
	}

	handler := jenkins.Handler{}

	// 2. query jobs from jenkins
	jobs, err := handler.InitData(url)
	if err != nil {
		fmt.Println(errorMsg(err))
		return
	}

	// 3. transform to feedback
	rawFeedback := handler.ToFeedBack(jobs)
	//log.Println(rawFeedback)

	// 4. filter by args
	feedback := handler.Filter(rawFeedback, query)
	//log.Println(feedback)

	if len(feedback.Body) < 1 {
		fmt.Println(emptyMsg())
		return
	}

	// 5. show it
	ret, jsonErr := json.Marshal(feedback)
	if jsonErr != nil {
		fmt.Println(errorMsg(err))
		return
	}

	fmt.Println(string(ret))
}

func errorMsg(err error) string {
	return `
	{
		"items": [
			"title": "ERROR",
			"subtitle": ` + err.Error() + `
		]
	}
`
}

func emptyMsg() string {
	return `
{
	"items": [
		"title": "Empty Result",
		"subtitle": "try with another query"
	]
}
`
}
