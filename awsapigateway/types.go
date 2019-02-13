package awsapigateway

type ResourceOptionsIface interface {
	resourceOptionsIfacePrivate()
}
type ResourceOptions struct{}

func (ResourceOptions) resourceOptionsIfacePrivate() {}

type ApiKeySourceTypeIface interface {
	apiKeySourceTypeIfacePrivate()
}
type ApiKeySourceType struct{}

func (ApiKeySourceType) apiKeySourceTypeIfacePrivate() {}
