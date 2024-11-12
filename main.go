package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Processor struct holds the details for each processor in the simulation
type Processor struct {
	Name        string
	Mean        float64
	StdDev      float64 // Updated field name for clarity
	Simulations int
}

// MonteCarloSimulation function runs the simulation for a given processor
func MonteCarloSimulation(processor Processor) []float64 {
	rand.Seed(time.Now().UnixNano())
	results := make([]float64, processor.Simulations)

	for i := 0; i < processor.Simulations; i++ {
		// Generate a random value from a normal distribution
		result := rand.NormFloat64()*processor.StdDev + processor.Mean
		results[i] = result
	}

	return results
}

// WriteResultsToCSV writes the simulation results to a CSV file
func WriteResultsToCSV(filename string, processorName string, results []float64) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Processor", "Simulation_Number", "Benchmark_Score"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write results
	for i, result := range results {
		row := []string{processorName, strconv.Itoa(i + 1), fmt.Sprintf("%.2f", result)}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// AnalyzeResults provides summary statistics for the simulation
func AnalyzeResults(processor Processor, results []float64) {
	sum := 0.0
	min := results[0]
	max := results[0]
	for _, value := range results {
		sum += value
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	mean := sum / float64(len(results))

	fmt.Printf("Results for %s:\n", processor.Name)
	fmt.Printf("Simulated Mean: %.2f (Target Mean: %.2f)\n", mean, processor.Mean)
	fmt.Printf("Minimum Value: %.2f\n", min)
	fmt.Printf("Maximum Value: %.2f\n", max)
	fmt.Printf("Sample Results: %.2f, %.2f, %.2f, %.2f, %.2f\n\n",
		results[0], results[1], results[2], results[3], results[4])
}

func main() {
	// Define processors and their parameters for JetStream2
	ryzen8700G := Processor{
		Name:        "AMD Ryzen 7 8700G @ 5.18GHz (JetStream2)",
		Mean:        370.28,
		StdDev:      1.11, // Use StdDev directly
		Simulations: 1000, // Generate 1000 simulated results
	}

	ryzen7700X := Processor{
		Name:        "AMD Ryzen 7 7700X @ 5.57GHz (JetStream2)",
		Mean:        405.31,
		StdDev:      0.68, // Use StdDev directly
		Simulations: 1000, // Generate 1000 simulated results
	}

	// Define processors and their parameters for OpenSSL 3.3 AES-256-GCM
	ryzen8700GOpenSSL := Processor{
		Name:        "AMD Ryzen 7 8700G @ 5.18GHz (OpenSSL 3.3 AES-256-GCM)",
		Mean:        103716166770,
		StdDev:      69039683.61, // Use StdDev directly
		Simulations: 1000,        // Generate 1000 simulated results
	}

	ryzen7700XOpenSSL := Processor{
		Name:        "AMD Ryzen 7 7700X @ 5.57GHz (OpenSSL 3.3 AES-256-GCM)",
		Mean:        116859009887,
		StdDev:      36966056.95, // Use StdDev directly
		Simulations: 1000,        // Generate 1000 simulated results
	}

	// Run Monte Carlo simulation for each processor and benchmark
	results8700G := MonteCarloSimulation(ryzen8700G)
	results7700X := MonteCarloSimulation(ryzen7700X)
	results8700GOpenSSL := MonteCarloSimulation(ryzen8700GOpenSSL)
	results7700XOpenSSL := MonteCarloSimulation(ryzen7700XOpenSSL)

	// Analyze and print summary statistics
	AnalyzeResults(ryzen8700G, results8700G)
	AnalyzeResults(ryzen7700X, results7700X)
	AnalyzeResults(ryzen8700GOpenSSL, results8700GOpenSSL)
	AnalyzeResults(ryzen7700XOpenSSL, results7700XOpenSSL)

	// Write results to CSV files
	if err := WriteResultsToCSV("ryzen8700G_jetstream2_results.csv", ryzen8700G.Name, results8700G); err != nil {
		fmt.Println("Error writing Ryzen 7 8700G JetStream2 results:", err)
	} else {
		fmt.Println("Ryzen 7 8700G JetStream2 results written to ryzen8700G_jetstream2_results.csv")
	}

	if err := WriteResultsToCSV("ryzen7700X_jetstream2_results.csv", ryzen7700X.Name, results7700X); err != nil {
		fmt.Println("Error writing Ryzen 7 7700X JetStream2 results:", err)
	} else {
		fmt.Println("Ryzen 7 7700X JetStream2 results written to ryzen7700X_jetstream2_results.csv")
	}

	if err := WriteResultsToCSV("ryzen8700G_openssl_results.csv", ryzen8700GOpenSSL.Name, results8700GOpenSSL); err != nil {
		fmt.Println("Error writing Ryzen 7 8700G OpenSSL results:", err)
	} else {
		fmt.Println("Ryzen 7 8700G OpenSSL results written to ryzen8700G_openssl_results.csv")
	}

	if err := WriteResultsToCSV("ryzen7700X_openssl_results.csv", ryzen7700XOpenSSL.Name, results7700XOpenSSL); err != nil {
		fmt.Println("Error writing Ryzen 7 7700X OpenSSL results:", err)
	} else {
		fmt.Println("Ryzen 7 7700X OpenSSL results written to ryzen7700X_openssl_results.csv")
	}
}
