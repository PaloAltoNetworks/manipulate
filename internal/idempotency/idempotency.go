package idempotency

// Keyer is the interface of an object
// that can set an Idempotency Key.
type Keyer interface {
	SetIdempotencyKey(string)
	IdempotencyKey() string
}
