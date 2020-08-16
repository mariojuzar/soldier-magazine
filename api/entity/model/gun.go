package model

type Gun struct {
	Name 		string		`json:"name"`
	Magazines	[]Magazine	`json:"magazines"`
}
