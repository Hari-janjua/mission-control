package ports

type AuthProvider interface {
	GetToken() (string, error)
}
