package ports

// AuthProvider defines the contract for authentication mechanisms
// (e.g., JWT, API keys, etc.) used by the Commander to issue and verify tokens.
type AuthProvider interface {
	GenerateToken(soldierID string) (string, error)
	ValidateToken(token string) (bool, string)
}
