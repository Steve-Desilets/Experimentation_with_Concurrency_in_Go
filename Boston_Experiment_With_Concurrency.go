package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Read CSV file
	file, err := os.Open("boston.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Parse data
	var data [][]float64
	var mv []float64
	for _, record := range records[1:] { // Skip the first row (headers)
		var row []float64
		for _, value := range record[1:] { // Exclude the first column (neighborhood)
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		data = append(data, row)
		target, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		mv = append(mv, target)
	}

	// Set up variables
	var runtimeSlice []int64
	loopCounter := 0

	// Create output file to which we will write the outputs
	file, err = os.Create("bostonOutputWithConcurrency.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	// Perform regression for each combination of four or more dependent variables
	for i := 0; i < 100; i++ {
		startTime := time.Now()

		numVars := len(data[0])
		var wg sync.WaitGroup
		for i := 4; i <= numVars; i++ {
			fmt.Printf("Regression with %d variables\n", i)
			combinations := getCombinations(numVars, i)
			for _, combo := range combinations {
				wg.Add(1)
				go func(combo []int) {
					defer wg.Done()
					coefficients, predictions := linearRegression(data, mv, combo)

					// Calculate mean square error
					mse := calculateMSE(mv, predictions)

					// Calculate AIC
					aic := calculateAIC(mse, i, len(mv))

					// Print coefficients
					fmt.Println("Coefficients:", coefficients)

					// Print results
					fmt.Printf("Variables: %v, MSE: %.4f, AIC: %.4f\n", combo, mse, aic)
				}(combo)
			}
		}
		wg.Wait()

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		runtimeSlice = append(runtimeSlice, executionTime.Microseconds())

		loopCounter += i
	}
	printSlice(runtimeSlice)

	// Calculate the total and average runtime across each of the experimental trials. Print these statistics to the output .txt file
	var runtimeSum int64 = 0

	for i := 0; i < 100; i++ {
		runtimeSum += runtimeSlice[i]
	}
	runtimeSumString := strconv.FormatInt(runtimeSum, 10)

	avgRuntime := (float64(runtimeSum)) / (float64(100))
	avgRuntimeString := fmt.Sprintf("%f", avgRuntime)

	fmt.Fprintf(file, "Summary Statistics For Experimental Trial Runtimes \n")

	fmt.Fprintf(file, "Runtime Sum in Microseconds \n")
	fmt.Fprintf(file, runtimeSumString)

	fmt.Fprintf(file, "\nAverage Trial Runtime in Microseconds \n")
	fmt.Fprintf(file, avgRuntimeString)
}

func linearRegression(data [][]float64, mv []float64, combo []int) ([]float64, []float64) {
	// Create the design matrix
	var designMatrix [][]float64
	for _, row := range data {
		var newRow []float64
		for _, idx := range combo {
			newRow = append(newRow, row[idx])
		}
		newRow = append(newRow, 1) // Add intercept term
		designMatrix = append(designMatrix, newRow)
	}

	// Calculate coefficients
	coefficients := calculateCoefficients(designMatrix, mv)

	// Calculate predicted values
	var predictions []float64
	for _, row := range designMatrix {
		predictions = append(predictions, predict(row, coefficients))
	}

	return coefficients, predictions
}

func calculateCoefficients(X [][]float64, y []float64) []float64 {
	var XtX [][]float64
	var Xty []float64

	// Calculate (X^T * X)
	for j := 0; j < len(X[0]); j++ {
		var row []float64
		for k := 0; k < len(X[0]); k++ {
			var val float64
			for i := 0; i < len(X); i++ {
				val += X[i][j] * X[i][k]
			}
			row = append(row, val)
		}
		XtX = append(XtX, row)
	}

	// Calculate (X^T * y)
	for j := 0; j < len(X[0]); j++ {
		var val float64
		for i := 0; i < len(X); i++ {
			val += X[i][j] * y[i]
		}
		Xty = append(Xty, val)
	}

	// Solve XtX * coefficients = Xty
	coefficients := gaussElimination(XtX, Xty)

	return coefficients
}

func gaussElimination(A [][]float64, b []float64) []float64 {
	n := len(b)
	coefficients := make([]float64, n)

	// Forward elimination
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			factor := A[j][i] / A[i][i]
			for k := i; k < n; k++ {
				A[j][k] -= factor * A[i][k]
			}
			b[j] -= factor * b[i]
		}
	}

	// Back substitution
	for i := n - 1; i >= 0; i-- {
		coefficients[i] = b[i]
		for j := i + 1; j < n; j++ {
			coefficients[i] -= A[i][j] * coefficients[j]
		}
		coefficients[i] /= A[i][i]
	}

	return coefficients
}

func predict(features []float64, coefficients []float64) float64 {
	var prediction float64
	for i := range features {
		prediction += features[i] * coefficients[i]
	}
	return prediction
}

func getCombinations(n, k int) [][]int {
	var result [][]int
	var helper func(int, []int)

	helper = func(start int, combination []int) {
		if len(combination) == k {
			temp := make([]int, len(combination))
			copy(temp, combination)
			result = append(result, temp)
			return
		}

		for i := start; i < n; i++ {
			combination = append(combination, i)
			helper(i+1, combination)
			combination = combination[:len(combination)-1]
		}
	}

	helper(0, []int{})
	return result
}

func sliceSubtract(a, b []float64) []float64 {
	result := make([]float64, len(a))
	for i := range result {
		result[i] = a[i] - b[i]
	}
	return result
}

func calculateMSE(actual, predicted []float64) float64 {
	var sumError float64
	for i := range actual {
		error := actual[i] - predicted[i]
		sumError += error * error
	}
	return sumError / float64(len(actual))
}

func calculateAIC(mse float64, numVars, numObs int) float64 {
	return float64(numObs) * (1 + 2*float64(numVars)) * (mse / float64(numObs))
}

// Define a function that will print out the slice of runtimes from each experimental trial
func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
