package shortener

// RedirectRepository is the redirect repo interface
type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
