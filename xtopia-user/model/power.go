package model

// Power unit's power
// ID: User/Team + ID = 0xxxx/1xxxx
// Kind: 0 -- user  1 -- team
// Labels: type -- label contents 1: TecPower 2: SoftPower 3: OtherPower
type Power struct {
	ID     string
	Kind   byte
	Powers map[uint8]string
}
