package cxapi

type MissingContextIface interface {
	GetProps() map[string]interface{}
	GetProvider() string
}

type MissingContext struct {
	Props    map[string]interface{}
	Provider string
}

func (m *MissingContext) GetProps() map[string]interface{} { return m.Props }
func (m *MissingContext) GetProvider() string              { return m.Provider }
