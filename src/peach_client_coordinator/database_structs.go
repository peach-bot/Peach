package main

type dbOption struct {
	OptionValue  string `json:"option_value"`
	Type         string `json:"type"`
	Experimental bool   `json:"experimental"`
	Beta         bool   `json:"beta"`
	Hidden       bool   `json:"hidden"`
}

type dbExtension struct {
	Options map[string]dbOption `json:"options"`
}

type dbSettings struct {
	Extensions map[string]dbExtension `json:"extensions"`
}
