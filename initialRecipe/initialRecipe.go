package initialrecipe

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("InitialRecipe", InitialRecipe)
}

func InitialRecipe(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

  var startData coffeeStartData
  var firstRecipe recipe

  jsonErr := json.Unmarshal(body, &startData)
  if jsonErr != nil {
    log.Fatal(err)
  }


  firstRecipe.GrindSize = calculateGrindSize(startData.CoffeeOrigin, startData.RoastLevel)
  firstRecipe.Dose = startData.BasketSize
  firstRecipe.BrewDuration = 26
  firstRecipe.WeightLiquidCoffee = startData.BasketSize * 2
  firstRecipe.Ratio = "1:2"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(firstRecipe)
}

type recipe struct {
  GrindSize int `json:"grindSize"`
  Dose int `json:"dose"`
  BrewDuration int `json:"brewDuration"`
  WeightLiquidCoffee int `json:"weightLiquidCoffee"`
  Ratio string `json:"ratio"`
}

type roastLevels string
const (
  Light roastLevels = "light"
  Medium roastLevels = "medium"
  Dark roastLevels = "dark"
)

type coffeeStartData struct {
  BasketSize int `json:"basketSize"`
  CoffeeOrigin string `json:"coffeeOrigin"`
  RoastLevel roastLevels  `json:"roastLevel"`
}

func calculateGrindSize(coffeeOrigin string, roastLevel roastLevels) int {
  switch roastLevel {
  case Light: 
    if coffeeOrigin == "Ethiopia" || coffeeOrigin == "Kenya" {
      return 2
    } else {
      return 3
    }
  case Medium:
    if coffeeOrigin == "Ethiopia" || coffeeOrigin == "Kenya" {
      return 3
    } else {
      return 4
    }
  case Dark:
    return 5
  default: 
    return 4
}

}
