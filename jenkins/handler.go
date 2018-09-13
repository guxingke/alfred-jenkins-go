package jenkins

import (
	"encoding/json"
	"github.com/guxingke/alfred-jenkins-go/alfred"
	"net/http"
	"io/ioutil"
	"strconv"
	"regexp"
)

type Handler struct {
	Url string
}

func (Handler *Handler) InitData(url string) ([]Job, error) {

	getUrl := url + "/api/json?tree=jobs[name,url,color,healthReport[description,score,iconUrl],lastBuild[number,result]]"

	resp, err := http.Get(getUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		return nil, ioErr
	}

	res := Jobs{}

	jsonErr := json.Unmarshal(respBytes, &res)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return res.Jobs, nil
}

func (Handler *Handler) Filter(raw alfred.Feedback, query string) alfred.Feedback {
	ret := alfred.Feedback{}
	for _, item := range raw.Body {
		matched, matchError := regexp.Match("(?i)("+query+")", []byte(item.Uid))
		if matchError != nil {
			continue
		}
		if matched {
			ret.AddItem(item)
		}
	}
	return ret
}

func (Handler *Handler) ToFeedBack(jobs []Job) alfred.Feedback {
	fb := alfred.Feedback{}
	for _, job := range jobs {
		fb.AddItem(toItem(job))
	}
	return fb
}

func toItem(job Job) alfred.Item {
	var description string
	var iconUrl = ""

	if len(job.HealthReport) > 0 {
		description = job.HealthReport[0].Description
		iconUrl = "static/images/" + job.HealthReport[0].IconUrl
	} else {
		description = "No Description"
	}

	buildNumber := strconv.Itoa(job.LastBuild.Number)

	title := job.Name + " [#" + buildNumber + " - " + job.LastBuild.Result + "]"
	return alfred.Item{job.Name, title, description, alfred.Icon{"path", iconUrl }, job.Url}
}
