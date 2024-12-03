package types

const (
	// noValueForTypeErrMsg is the error message for when a node
	// does not have a value for the given type.
	noValueForTypeErrMsg = "no value for type %v"

	// multipleProvidersErrMsg is the error message for when multiple
	// providers are registered for the same type.
	multipleProvidersErrMsg = "multiple providers registered for type %v"

	// noProvidersErrMsg is the error message for when no providers
	// are registered for the given type.
	noProvidersErrMsg = "no providers registered for type %v"
)
