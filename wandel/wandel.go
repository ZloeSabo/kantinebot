package wandel

import (
	"log"
	"net/http"
	"os"
)

const (
	// APIURL of the wandel api.
	APIURL = "https://api.qnips.com/cons/api/"
)

// httpClient defines the minimal interface needed for an http.Client to be implemented.
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type logger interface {
	Printf(format string, v ...interface{})
}

type debuglogger interface {
	DebugEnabled() bool
	DebugPrintf(format string, v ...interface{})
}

type Client struct {
	authorization string
	endpoint      string
	app           string
	restaraunt    string
	debug         bool
	log           logger
	httpclient    httpClient
}

type Option func(*Client)

func OptionAuthorization(authorization string) func(*Client) {
	return func(c *Client) {
		c.authorization = authorization
	}
}

func OptionDebug(debug bool) func(*Client) {
	return func(c *Client) {
		c.debug = debug
	}
}

//NewClient returns new Wandel API client
func NewClient(options ...Option) *Client {
	w := &Client{
		authorization: "",
		endpoint:      APIURL,
		app:           "Wandel",
		restaraunt:    "10191", //TODO make it configurable
		debug:         false,
		httpclient:    &http.Client{},
		log:           log.New(os.Stderr, "zloesabo/wandel", log.LstdFlags|log.Lshortfile),
	}

	for _, opt := range options {
		opt(w)
	}

	return w
}

func (client *Client) DebugEnabled() bool {
	return client.debug
}

func (client *Client) DebugPrintf(format string, v ...interface{}) {
	if client.debug {
		client.log.Printf(format, v)
	}
}
