package status

type Status struct {
	Components map[string]bool `json:"components,omitempty"`
}
