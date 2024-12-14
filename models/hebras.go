package models

type HebrasTea struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	RawPrice int    `json:"rawPrice"`
	From     string `json:"from"`
	Img      string `json:"img,omitempty"`
	Link     string `json:"link,omitempty"`
	Grams    string `json:"grams,omitempty"`
}
