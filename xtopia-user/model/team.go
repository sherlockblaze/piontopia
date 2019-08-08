package model

// Team Basic team info
type Team struct {
	ID       string
	ParentID string
	TeamAttr
	Labels
	Power
}

// TeamAttr composition
type TeamAttr struct {
	MemNumber uint8
	Type      []string
}
