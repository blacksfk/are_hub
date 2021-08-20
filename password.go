package are_hub

// password implements json.Marshaler.
type password string

// Create a password from a string.
func NewPassword(str string) password {
	return password(str)
}

// Prevent passwords from being marshaled and sent to clients by implementing
// json.Marshaler and returning a empty byte slice. This allows for passwords
// to be unmarshaled without issue.
func (p password) MarshalJSON() ([]byte, error) {
	return []byte(`""`), nil
}

// Get the password as a native string.
func (p password) String() string {
	return string(p)
}

// Set the password to the string.
func (p *password) Set(str string) {
	*p = password(str)
}
