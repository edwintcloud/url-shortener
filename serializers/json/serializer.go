package json

import (
	"encoding/json"

	"github.com/edwintcloud/url-shortener/shortener"
	errs "github.com/pkg/errors"
)

// Redirect represents a redirect
type Redirect struct{}

// Decode decodes bytes to a Redirect
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode encodes a redirect to bytes
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
