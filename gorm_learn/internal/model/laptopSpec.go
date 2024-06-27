package model

type LaptopSpec struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Cpu     string `json:"cpu"`
	Ram     string `json:"ram"`
	Storage string `json:"storage"`
}
