package manager

// ConnectionState represents the state of the SimConnect connection.
type ConnectionState int

const (
	Offline ConnectionState = iota
	Connecting
	Online
)
