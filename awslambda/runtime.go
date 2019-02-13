package awslambda

type RuntimeIface interface {
	runtimePrivate()
}
type Runtime struct{}

func (Runtime) runtimePrivate() {}

func Runtime_NodeJS810() string { return "" }

type CodeIface interface {
	codePrivate()
}
type Code struct{}

func (Code) codePrivate() {}

type AssetCodeIface interface {
	assetCodePrivate()
}
type AssetCode struct{}

func (AssetCode) assetCodePrivate() {}

func Code_Asset(path string) AssetCodeIface { return nil }
