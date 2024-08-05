package client

import (
	"net/url"

	"github.com/ginger-core/errors"
)

func (c *client[T]) GetDownloadUrlByKey(key string) (string, errors.Error) {
	if c.properties == nil {
		return "", errors.Internal().
			WithId("ServiceUnavailable").
			WithMessage("Service is unavailable right now")
	}

	path, err := url.JoinPath(c.properties.KeyBaseUrl, key)
	if err != nil {
		return "", errors.New(err).
			WithDetail(
				errors.NewDetail().
					With("keyBaseUrl", c.properties.KeyBaseUrl).
					With("key", key)).
			WithTrace("url.JoinPath")
	}
	return path, nil
}
