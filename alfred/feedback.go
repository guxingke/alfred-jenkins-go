package alfred

type Feedback struct {
	Body []Item `json:"items"`
}

type Items struct {
	Body []Item `json:"item"`
}

type Item struct {
	Uid      string `json:"uid"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Icon     Icon `json:"icon"`
	Arg      string `json:"arg"`
}

type Icon struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

func (fb *Feedback) AddItem(item Item) {
	fb.Body= append(fb.Body, item)
}
