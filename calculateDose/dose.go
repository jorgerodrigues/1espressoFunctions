package dose

import (
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"io"
	"log"
	"net/http"
	"strconv"
)

func init() {
	functions.HTTP("Dose", Dose)
}

type taste string

const (
	Bitter taste = "bitter"
	Sour   taste = "sour"
)

// helloHTTP is an HTTP Cloud Function with a request parameter.
func Dose(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var doseResult dose
	jsonErr := json.Unmarshal(body, &doseResult)
	if jsonErr != nil {
		log.Fatal(err)
	}
	grindSize, _ := strconv.Atoi(doseResult.GrindSize)
	duration, _ := strconv.Atoi(doseResult.BrewDuration)
	basketSize, _ := strconv.Atoi(doseResult.BaseketSize)
	coffeDose, _ := strconv.Atoi(doseResult.Dose)

	if canAdjustGrind(grindSize, doseResult.TasteResult) {
		grindSize = adjustGrind(grindSize, doseResult.TasteResult)
	}

	if canAdjustDose(coffeDose, doseResult.TasteResult, basketSize) {
		coffeDose = adjustDose(coffeDose, basketSize, doseResult.TasteResult)
	}

	if canAdjustBrewDuration(duration, doseResult.TasteResult) {
		duration = adjustBrewDuration(duration, doseResult.TasteResult)
	}

	finalResult := dose{
		GrindSize:          strconv.Itoa(grindSize),
		Dose:               strconv.Itoa(coffeDose),
		BrewDuration:       strconv.Itoa(duration),
		WeightLiquidCoffee: doseResult.WeightLiquidCoffee,
		BaseketSize:        doseResult.BaseketSize,
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fmt.Sprintf("Hello, %s!", finalResult))
}

type dose struct {
	TasteResult        taste  `json:"tasteResult"`
	Dose               string `json:"dose"`
	BrewDuration       string `json:"duration"`
	GrindSize          string `json:"grindSize"`
	WeightLiquidCoffee string `json:"weightLiquidCoffee"`
	BaseketSize        string `json:"basketSize"`
}

func canAdjustGrind(grindSize int, tasteResult taste) bool {
	if tasteResult == Bitter && grindSize < 10 {
		return true
	}

	if tasteResult == Sour && grindSize > 1 {
		return true
	}

	return false
}

func canAdjustDose(dose int, tasteResult taste, baseketSize int) bool {
	if tasteResult == Bitter && dose > baseketSize-1 {
		return true
	}

	if tasteResult == Sour && dose < baseketSize+2 {
		return true
	}

	return false
}

func canAdjustBrewDuration(previousTime int, tasteResult taste) bool {
	if tasteResult == Bitter && previousTime < 30 {
		return true
	}

	if tasteResult == Sour && previousTime > 20 {
		return true
	}
	return false
}

func adjustGrind(grindSize int, tasteResult taste) int {
	if tasteResult == Bitter {
		return grindSize + 1
	}
	return grindSize - 1
}

func adjustDose(dose int, baseketSize int, tasteResult taste) int {
	if tasteResult == Bitter {
		return dose - 1
	}
	return dose + 1
}

func adjustBrewDuration(previousTime int, tasteResult taste) int {
	if tasteResult == Bitter {
		return previousTime + 1
	}
	return previousTime - 1
}
