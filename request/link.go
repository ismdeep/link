package request

// LinkSave link save req struct
type LinkSave struct {
	ID          string `json:"id"`          // target id
	Target      string `json:"target"`      // target url
	Description string `json:"description"` // description
}
