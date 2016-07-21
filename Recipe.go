package main
type Ingredient struct {
	Quantity 	float32  `json:"quantity`
	Unit		string `json:"unit"`
    Name 		string `json:"name"`
}

type Recipe struct {
	ID 				int `json:"id`
    Name 			string `json:"name"`
    Ingredients 	[]Ingredient
    Instructions 	string `json:"instructions"`
}