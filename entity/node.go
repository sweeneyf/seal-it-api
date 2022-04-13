package entity

type Node struct {
	Name          string  `json:"name"`
	SpaceFreeMB   float64 `json:"spaceMB"`
	SpaceFreePerc float64 `json:"spacePerc"`
}
