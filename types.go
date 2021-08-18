package main

type rawSpecField struct {
	SectionID int
	NameID    int
	ValueID   int
}

type rawProduct struct {
	ArkID int
	ID    int
	Title string
}

type populatedProduct struct {
	ArkID int        `json:"id"`
	Title string     `json:"name"`
	Specs [][]string `json:"specs"`
}
