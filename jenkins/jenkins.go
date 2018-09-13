package jenkins

/*
{
  "jobs": [
    {
      "_class": "org.jenkinsci.plugins.workflow.job.WorkflowJob",
      "name": "athena-op",
      "url": "https://ci-new.duitang.net/job/athena-op/",
      "color": "blue",
      "healthReport": [
        {
          "description": "Build stability: No recent builds failed.",
          "iconUrl": "health-80plus.png",
          "score": 100
        }
      ],
      "lastBuild": {
        "_class": "org.jenkinsci.plugins.workflow.job.WorkflowRun",
        "number": 1454,
        "result": "SUCCESS"
      }
    }
  ]
}
*/

type Jobs struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Color string `json:"color"`
	HealthReport []HealthReport `json:"healthReport"`
	LastBuild LastBuild `json:"lastBuild"`
}

type HealthReport struct {
	Description string `json:"description"`
	IconUrl string `json:"iconUrl"`
	Score int `json:"score"`
}

type LastBuild struct {
	Number int `json:"number"`
	Result string `json:"result"`
}
