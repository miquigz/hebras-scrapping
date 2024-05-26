package models

type HebrasTea struct {
	Name     string `json:"name,omitempty"`
	Price    string `json:"price,omitempty"`
	RawPrice int    `json:"rawPrice,omitempty"`
	From     string `json:"from,omitempty"`
	Img      string `json:"img,omitempty"`
	Link     string `json:"link,omitempty"`
	Grams    string `json:"grams,omitempty"`
}
