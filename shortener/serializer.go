package shortener

// RedirectSerializer is the redirect serializer interface
type RedirectSerializer interface {
	Decode(input []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
