# Monte Carlo Simulation for Processor Benchmarks

This Go program performs Monte Carlo simulations to evaluate the performance of different processors under two benchmarking tests: JetStream2 and OpenSSL 3.3 AES-256-GCM. The results are written to CSV files for further analysis.

## Processors and Benchmarks

The following processors and benchmarks are included in the simulation:

1. **AMD Ryzen 7 8700G @ 5.18GHz (JetStream2)**
2. **AMD Ryzen 7 7700X @ 5.57GHz (JetStream2)**
3. **AMD Ryzen 7 8700G @ 5.18GHz (OpenSSL 3.3 AES-256-GCM)**
4. **AMD Ryzen 7 7700X @ 5.57GHz (OpenSSL 3.3 AES-256-GCM)**

## Simulation Results

### JetStream2

**AMD Ryzen 7 8700G @ 5.18GHz (JetStream2)**
- Simulated Mean: 370.30 (Target Mean: 370.28)
- Minimum Value: 367.10
- Maximum Value: 373.73
- Sample Results: 370.22, 370.08, 371.32, 368.79, 370.92

**AMD Ryzen 7 7700X @ 5.57GHz (JetStream2)**
- Simulated Mean: 405.30 (Target Mean: 405.31)
- Minimum Value: 403.40
- Maximum Value: 407.64
- Sample Results: 404.99, 405.48, 405.50, 404.73, 405.17

### OpenSSL 3.3 AES-256-GCM

**AMD Ryzen 7 8700G @ 5.18GHz (OpenSSL 3.3 AES-256-GCM)**
- Simulated Mean: 103718694489.95 (Target Mean: 103716166770.00)
- Minimum Value: 103500669160.24
- Maximum Value: 103943521900.81
- Sample Results: 103737298017.52, 103634981077.05, 103695822241.70, 103757291309.08, 103808331896.21

**AMD Ryzen 7 7700X @ 5.57GHz (OpenSSL 3.3 AES-256-GCM)**
- Simulated Mean: 116859024015.93 (Target Mean: 116859009887.00)
- Minimum Value: 116744030442.73
- Maximum Value: 116988933673.84
- Sample Results: 116847983137.84, 116872780673.05, 116880914665.77, 116887924060.32, 116839872311.03

## Generated Files

The simulation results are written to the following CSV files:

- `ryzen8700G_jetstream2_results.csv`
- `ryzen7700X_jetstream2_results.csv`
- `ryzen8700G_openssl_results.csv`
- `ryzen7700X_openssl_results.csv`

## Running the Simulation

To run the simulation, ensure appropriate mean and SE values are included and execute the following command:

```sh
go run main.go