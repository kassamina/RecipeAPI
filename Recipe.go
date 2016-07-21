//By: Zoe Toth
package main

type Catagory struct {
	ID 			int `json:"id`
	Name 		string `json:"name"`
	Recipes 	[]*Recipe
}
type Ingredient struct {
	Quantity 	float32  `json:"quantity`
	Unit 		string `json:"unit"`
	Name 		string `json:"name"`
}

//only stories the id of the catagory, name must be looked up in catagory table
type Recipe struct {
	ID 				int `json:"id`
    Name 			string `json:"name"`
    CatagoryIDs		[]int
    Ingredients 	[]Ingredient
    Instructions 	string `json:"instructions"`
}