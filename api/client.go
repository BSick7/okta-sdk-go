package api

import (
	"net/http"
)

type Client struct {
	session  Session
	executor *ClientExecutor
}

func NewClient(session Session) *Client {
	return &Client{
		session: session,
		executor: NewClientExecutor(session, &http.Client{
			Transport: session.CreateTransport(),
		}),
	}
}

func (c *Client) Apps() *Apps {
	return NewApps(c.executor)
}
