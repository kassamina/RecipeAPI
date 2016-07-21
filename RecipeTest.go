//By: Zoe Toth
package main

import (
    //"fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io"
    "os"
    "strconv"
    "net/url"
)
/*
    Testing is not in proper format, and ideally should be automated, but I will admit that I am not farmilar with go
    prefering to test as I went, and focus on how go itself worked. 
*/

func testAddRecipe() {
    newRecipe := Recipe{ID: 0, Name: "Boiled Egg", CatagoryIDs: []int{0}, Ingredients: []Ingredient{{1.0,"", "egg"}}, Instructions: "Boil the egg(s)."}
   
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(newRecipe)
    res, _ := http.Post("http://localhost:8081/add", "application/json; charset=utf-8", b)
    io.Copy(os.Stdout, res.Body)
}

func testReturnRecipe(id int) {
    data := url.Values{}
    data.Set("ID", strconv.Itoa(id))

    r, _ := http.Post("http://localhost:8081/return", "application/x-www-form-urlencoded; charset=utf-8", bytes.NewBufferString(data.Encode())) 
    io.Copy(os.Stdout, r.Body)
}

func testDeleteRecipe(id int) {
    data := url.Values{}
    data.Set("ID", strconv.Itoa(id))

    r, _ := http.Post("http://localhost:8081/delete", "application/x-www-form-urlencoded; charset=utf-8", bytes.NewBufferString(data.Encode())) 
    io.Copy(os.Stdout, r.Body)
}

func testAlterRecipe(id int) {
    //will double the amount of eggs in recipe with id
    data := url.Values{}
    data.Set("ID", strconv.Itoa(id))

    //get the old version of the recipe
    r, _ := http.Post("http://localhost:8081/return", "application/x-www-form-urlencoded; charset=utf-8", bytes.NewBufferString(data.Encode())) 

    decoder := json.NewDecoder(r.Body)
    var oldRecipe Recipe
    err := decoder.Decode(&oldRecipe)
    if err != nil {
        panic(err)
    }
    oldRecipe.Ingredients[0].Quantity = 2.0

    //send in the alteration request
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(oldRecipe)
    res, _ := http.Post("http://localhost:8081/alter", "application/json; charset=utf-8", b)
    io.Copy(os.Stdout, res.Body)
}

func testReturnCat(id int) {
    data := url.Values{}
    data.Set("ID", strconv.Itoa(id))

    r, _ := http.Post("http://localhost:8081/cat", "application/x-www-form-urlencoded; charset=utf-8", bytes.NewBufferString(data.Encode())) 
    io.Copy(os.Stdout, r.Body)
}
func main() {
    for i := 0; i < 20; i++ {
        testAddRecipe()
    }
    //testing delete
    testReturnRecipe(8)
    testDeleteRecipe(8)
    testReturnRecipe(8)
    
    //re-deletion?
    //testDeleteRecipe(8)
    
    //the delete will have shifted the indices in the "DB", check to see if return still works
    testReturnRecipe(20)

    //should return 22
    //testAddRecipe()

    testAlterRecipe(21)

    testReturnCat(0)


}