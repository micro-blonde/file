package client

import "github.com/ginger-core/errors"

var notEnabledError = errors.Forbidden().
	WithId("FileClientNotEnabledError").
	WithMessage("client is not enabled to perform your request")
