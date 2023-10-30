package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sajari/regression"
)

// load dataset
// split train test data
// create and fit regression model with training data
// evaluate model use testing data coming out R square
func main() {
	data, err := loadDataset("boston.csv")
	if err != nil {
		log.Fatal(err)
	}

	trainingData, testingData := splitDataset(data, 0.8)

	startTime := time.Now()
	model := trainModel(trainingData)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Training Time: %s\n", elapsedTime)

	rSquared := evaluateModel(model, testingData)
	fmt.Printf("R-squared: %0.2f\n", rSquared)
}

// load csv and skip the header
func loadDataset(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	//get rid of the header
	records = records[1:]
	numFields := len(records[0]) - 1 //make slice
	data := make([][]float64, len(records))
	for i, record := range records {
		data[i] = make([]float64, numFields)
		for j := 1; j < len(record); j++ { // Start from the second column
			field := strings.TrimSpace(record[j])
			// covert field to float and set to zero when it couldnt convert
			val, err := strconv.ParseFloat(field, 64)
			if err != nil {
				val = 0
			}
			data[i][j-1] = val // Adjust the index
		}
	}
	return data, nil
}

// ratio = training/all data
func splitDataset(data [][]float64, ratio float64) ([][]float64, [][]float64) {
	splitIndex := int(float64(len(data)) * ratio) //calcuate the split index as training & testing data
	return data[:splitIndex], data[splitIndex:]
}

func trainModel(trainingData [][]float64) *regression.Regression {
	r := new(regression.Regression)               // pointer type or var r regression.Regression
	r.SetObserved("Price")                        //as mv the dependent variable
	for i := 0; i < len(trainingData[0])-1; i++ { //independent vars
		r.SetVar(i, fmt.Sprintf("Var%d", i))
	}
	//train model
	for _, record := range trainingData {
		y := record[len(record)-1] //mv
		x := record[:len(record)-1]
		r.Train(regression.DataPoint(y, x))
	}

	r.Run()                                         //fit model
	fmt.Println("Regression Formula:\n", r.Formula) //print parameters
	return r
}

// function creates predicted data using testing data
// using predicted data create r2 score
func evaluateModel(r *regression.Regression, testingData [][]float64) float64 {
	actualV := make([]float64, len(testingData))
	predicted := make([]float64, len(testingData))
	var sumActual float64
	for i, record := range testingData {
		y := record[len(record)-1]
		x := record[:len(record)-1]
		actualV[i] = y
		predictedValue, err := r.Predict(x)
		if err != nil {
			log.Fatal("Prediction error:", err)
		}
		predicted[i] = predictedValue
		sumActual += y
	}

	meanObserved := sumActual / float64(len(actualV))
	var ssTot, ssRes float64
	for i := 0; i < len(actualV); i++ {
		ssTot += math.Pow(actualV[i]-meanObserved, 2)
		ssRes += math.Pow(actualV[i]-predicted[i], 2)
	}
	rSquared := 1 - (ssRes / ssTot)
	return rSquared
}
