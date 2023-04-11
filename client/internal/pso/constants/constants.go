package constants

const (
	UnseenServerName  = "unseen"
	EphineaServerName = "ephinea"
	UltimaServerName  = "ultima"
)

type EphineaAccountMode int

const (
	Normal EphineaAccountMode = iota
	Hardcore
	Sandbox
)
