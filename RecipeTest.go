package main

import (
    "fmt"
    "bytes"
    "net/http"
    "encoding/json"
    "io"
    "os"
    "strconv"
    "net/url"
)
type Recipe struct {
    ID int `json:"id`
    Name string `json:"name"`
    Ingredients string `json:"ingredients"`
    Instructions string `json:"instructions"`
}

func testAddRecipe() {
    newRecipe := Recipe{ID: 0, Name: "Boiled Egg", Ingredients: "1 Egg", Instructions: "Boil the egg."}
   
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(newRecipe)
    res, _ := http.Post("http://localhost:8081/add", "application/json; charset=utf-8", b)
    io.Copy(os.Stdout, res.Body)
}

func testReturnRecipe(id int) {
    data := url.Values{}
    data.Set("ID", strconv.Itoa(id))

    r, _ := http.Post("http://localhost:8081/return", "application/x-www-form-urlencoded; charset=utf-8", bytes.NewBufferString(data.Encode())) 

    decoder := json.NewDecoder(r.Body)
    var newRecipe Recipe
    err := decoder.Decode(&newRecipe)

    if err != nil {
        //panic(err)
    }
    fmt.Println(newRecipe)

}

func main() {
    testAddRecipe()
    testReturnRecipe(0)
    testReturnRecipe(1)
}