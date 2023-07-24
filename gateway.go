package goway

type Gateway struct {
	Register *Register

	Addr string
	Port int

	PingInterval int
	PingData     []byte
}
