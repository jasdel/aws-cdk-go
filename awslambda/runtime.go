package awslambda

type Runtime interface{}
type Runtime_ struct{}

func Runtime_NodeJS810() string { return "" }

type Code interface{}
type Code_ struct{}

func Code_Asset(path string) AssetCode { return nil }

type AssetCode interface{}
type AssetCode_ struct{}
