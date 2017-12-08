package api

import (
	"net/url"
	"strconv"

	"github.com/BSick7/okta-sdk-go/structs/app"
)

type Apps struct {
	executor *ClientExecutor
}

func NewApps(executor *ClientExecutor) *Apps {
	return &Apps{
		executor: executor,
	}
}

type App struct {
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Label         string                 `json:"label"`
	Status        string                 `json:"status"`
	LastUpdated   string                 `json:"lastUpdated"`
	Created       string                 `json:"created"`
	Accessibility *app.Accessibility     `json:"accessibility"`
	Visibility    *app.Visibility        `json:"visibility"`
	Features      []*string              `json:"features"`
	SignOnMode    string                 `json:"signOnMode"`
	Credentials   *app.Credentials       `json:"credentials"`
	Settings      map[string]interface{} `json:"settings"`
}

func (a *App) IsActive() bool {
	return a.Status == "ACTIVE"
}

func (a *Apps) List() ([]*App, error) {
	req, err := a.executor.NewRequest()
	if err != nil {
		return nil, err
	}
	req.SetEndpoint("/apps")
	req.SetQuery(url.Values{
		"limit": []string{strconv.Itoa(20)},
		//"filter": []string{},
		//"after": []string{},
		//"expand": []string{},
	})

	res, err := req.Get()
	if err != nil {
		return nil, err
	}

	list := make([]*App, 0)
	if err := res.BodyJSON(&list); err != nil {
		return nil, err
	}
	return list, nil
}
