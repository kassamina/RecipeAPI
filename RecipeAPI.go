package main
/*
By: Zoe Toth

todo
	-port is fixed, should be parameter?

	improvements
		-
		-catagories, breakfast, etc.
			implemented as in relational DB
			breakfast would be a table with pointers to related recipes, ie. pointer to fried egg, poached egg, etc. 


*/
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "strconv"
)


var recipeDB []Recipe
var nextID int

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
    fmt.Println("homepage hit?")
}




/*
	Implemented because indicies in recipesDB are garunteed to be in sorted order,
	but some in the middle may be deleted. 

	O(logN), I don't believe go implements tail recursion though, so this will still not be efficent, and should be implemented in place. 
*/
func binarySearch(low int, high int, target int) int {
	if low > high {
		return -1
	}

	m := (low + high)/2
	if recipeDB[m].ID < target {
    	low = m + 1
    } else if recipeDB[m].ID > target {
    	high = m -1
    } else {
    	return m
    }
    return binarySearch(low, high, target)
}//binarySearch




/*	returnRecipe
	Input: 	Expects a http post body parameter, named "ID" containing the recipeID you wish to be returned.
			request should be sent to http://localhost:8081/return
			
	Return:	json encoded Recipe
	
	Error:	Returns error message in body of response
			should likely be implemented as a GET, with a url parameter */
func returnRecipe(w http.ResponseWriter, r *http.Request) {
	ID,err := strconv.Atoi(r.FormValue("ID"))
	if err != nil {
		panic(err)
	}

    fmt.Println("recieved request to return recipeID: %d", ID)
    
    i := -1
    if ID < nextID && ID >= 0 {
    	i = binarySearch(0, nextID-1, ID)
    	if i == -1 {
    		fmt.Println("I'm sorry, that's no longer a valid recipeID.")
    		fmt.Fprintf(w,"I'm sorry, that's no longer a valid recipeID.")
    	} else {
    		json.NewEncoder(w).Encode(recipeDB[i])
    	}
    } else {
    	fmt.Println("Invalid ID! return request denied!")
    	fmt.Fprintf(w,"Invalid ID! return request denied!")
    }
}//returnRecipe




/*	delRecipe
	Input: 	Expects a http post body parameter, named "ID" containing the recipeID you wish to delete
			request should be sent to http://localhost:8081/delete
	Return:	Sends back the deleted recipe's ID on success

	Error: 	Returns error message in body of response*/
func delRecipe(w http.ResponseWriter, r *http.Request){
	ID,err := strconv.Atoi(r.FormValue("ID"))
	if err != nil {
		panic(err)
	}

    fmt.Println("recieved request to delete recipeID: %d", ID)

    i := -1
    if ID < nextID && ID >= 0 {
    	i = binarySearch(0, nextID-1, ID)
    	if i == -1 {
    		fmt.Println("I'm sorry, that's no longer a valid recipeID.")
    		fmt.Fprintf(w,"I'm sorry, that's no longer a valid recipeID.")
    	} else {
    		//appends the list of all recipes with id's <i, to the list of all recipes with id's >i
			recipeDB = append(recipeDB[:i], recipeDB[i+1:]...)
			fmt.Fprintf(w, "%d", i)
			fmt.Println("Successfully deleted recipe with id == %d", i)
    	}
    } else {
    	fmt.Println("Invalid ID! deletion request denied!")
    	fmt.Fprintf(w,"Invalid ID! deletion request denied!")
    }

    fmt.Println("Successfully deleted recipe with id == %d", i)
}//delRecipe




/*	returnAllRecipes
	Input:	Any request sent to http://localhost:8081/all

	Return: All recipes that are currently being stored, in order of their ID, in json

	Error: N/A */
func returnAllRecipes(w http.ResponseWriter, r *http.Request){  
    fmt.Println("Recieved request to return all recipes")
    json.NewEncoder(w).Encode(recipeDB)
}//returnAllRecipes




/*	addRecipe
	Input: json formated Recipe in body of request sent to http://localhost:8081/add

	Return:	The recipes new ID 

	Error: Throws error if the recipe cannot be decoded*/
func addRecipe(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var newRecipe Recipe
	err := decoder.Decode(&newRecipe)
	if err != nil {
		fmt.Fprintf(w, "Error: Inproperly formated Recipe")
		panic(err)
	}

	newRecipe.ID = nextID
	recipeDB = append(recipeDB, newRecipe)
	//send back the recipe's new ID as confirmation of success
	fmt.Fprintf(w, "%d",nextID)
	nextID++
}//addRecipe


/*	alterRecipe
	Input: 	Expected json formated Recipe in body of request sent to http://localhost:8081/alter
			the recipe's id should match an existing recipe
			
	Return:	The updated json encoded Recipe in body of response
	
	Error:	Returns error message in body of response
			should likely be implemented as a GET, with a url parameter */
func alterRecipe(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var alteredRecipe Recipe
	err := decoder.Decode(&alteredRecipe)
	if err != nil {
		fmt.Fprintf(w, "Error: Inproperly formated Recipe")
		panic(err)
	}

	ID := alteredRecipe.ID

	i := -1
    if ID < nextID && ID >= 0 {
    	i = binarySearch(0, nextID-1, ID)
    	if i == -1 {
    		fmt.Println("I'm sorry, that's no longer a valid recipeID.")
    		fmt.Fprintf(w,"I'm sorry, that's no longer a valid recipeID.")
    	} else {
    		//appends the list of all recipes with id's <i, to the list of all recipes with id's >i
			recipeDB[i] = alteredRecipe
			json.NewEncoder(w).Encode(recipeDB[i])
			fmt.Println("Successfully altered recipe with id == %d", ID)
    	}
    } else {
    	fmt.Println("Invalid ID! alteration request denied!")
    	fmt.Fprintf(w,"Invalid ID! alteration request denied!")
    }
}//alterRecipe


func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/all", returnAllRecipes)
    http.HandleFunc("/return", returnRecipe)
    http.HandleFunc("/delete", delRecipe)
    http.HandleFunc("/add", addRecipe)
    http.HandleFunc("/alter", alterRecipe)
    log.Fatal(http.ListenAndServe(":8081", nil))
}//handRequests

func main() {
	//create the slice which will store all of the recipes; 
	recipeDB = make([]Recipe, 0)
	recipeDB = append(recipeDB, Recipe{ID: 0, Name: "Fried Egg", Ingredients: []Ingredient{{1.0,"", "egg"}, {2.0,"teaspoons", "butter"}}, Instructions: "Fry the egg(s) in the butter."})
	recipeDB = append(recipeDB, Recipe{ID: 1, Name: "Poached Egg", Ingredients: []Ingredient{{1.0,"", "egg"}}, Instructions: "Poach the egg(s)."})
	nextID = 2

	//fmt.Println(recipeDB[0])
    handleRequests()
}//main



