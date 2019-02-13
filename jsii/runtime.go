package jsii

type Runtime struct {
	client *Client
}

var GlobalRuntime = &Runtime{
	client: NewClient(),
}

func (r *Runtime) Client() *Client {
	return r.client
}
