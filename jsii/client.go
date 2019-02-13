package jsii

import "strconv"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Create(fqn string, args []interface{}, overrides interface{}) (string, error) {
	// TODO register callbacks for overrides with JSII kernel.
	// https://github.com/awslabs/jsii/issues/263

	return fqn + "@" + strconv.Itoa(uID()), nil
}

func (c *Client) Invoke(id, method string, args []interface{}) (Result, error) {
	return Result{}, nil
}

func (c *Client) Get(id, property string) (Result, error) {
	return Result{}, nil
}

func (c *Client) Set(id, property string, value interface{}) (Result, error) {
	return Result{}, nil
}

//... more API methods

var runningID int

func uID() int {
	runningID++
	return runningID
}

type Result struct{}
