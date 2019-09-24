package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	simulator "github.com/monochromegane/queuing-theory-simulator"
)

var (
	seed   int
	step   int
	server int
	lambda float64
	mu     float64

	outDir     string
	params     string
	simulation string
	history    string
)

func init() {
	flag.IntVar(&seed, "seed", 1, "Seed")
	flag.IntVar(&step, "step", 1, "Number of iterations")
	flag.IntVar(&server, "server", 1, "Number of servers")
	flag.Float64Var(&lambda, "lambda", 0.1, "Lambda (Number of arrival in UnitTime)")
	flag.Float64Var(&mu, "mu", 0.2, "Mu (Number of service in UnitTime)")

	flag.StringVar(&outDir, "dir", "out", "Output directory")
	flag.StringVar(&params, "params", "params.csv", "File name for parameters")
	flag.StringVar(&simulation, "simulation", "simulation.csv", "File name for simulation result")
	flag.StringVar(&history, "history", "history.csv", "File name for request history")
}

func main() {
	flag.Parse()

	fParams, fSimulation, fHistory, err := setup()
	if err != nil {
		panic(err)
	}
	defer func() {
		fParams.Close()
		fSimulation.Close()
		fHistory.Close()
	}()

	served := 0
	avgResponse := 0.0
	lambdas := func(t int) float64 { return lambda }
	mus := func(t int) float64 { return mu }
	model := simulator.NewMMSModel(int64(seed), lambdas, mus)
	for i := 0; i < step; i++ {
		_, processing, waiting, rs := model.Progress(server)
		for _, r := range rs {
			avgResponse = onlineAvg(r, served, avgResponse)
			served++
		}
		fmt.Fprintf(fSimulation, "%d,%d,%f\n", processing, waiting, avgResponse)
	}

	histories := model.History()
	for _, h := range histories {
		fmt.Fprintf(fHistory, "%d,%d,%d,%d\n", h[0], h[1], h[2], h[3])
	}
}

func setup() (*os.File, *os.File, *os.File, error) {
	if outDir != "" {
		err := os.MkdirAll(outDir, 0755)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	fParams, err := os.OpenFile(filepath.Join(outDir, params), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Fprintf(fParams, "seed,step,server,lambda,mu\n")
	fmt.Fprintf(fParams, "%d,%d,%d,%f,%f\n", seed, step, server, lambda, mu)

	fSimulation, err := os.OpenFile(filepath.Join(outDir, simulation), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Fprintf(fSimulation, "processing,waiting,averageResponseTime\n")

	fHistory, err := os.OpenFile(filepath.Join(outDir, history), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Fprintf(fHistory, "arriveAt,serviceTime,beginAt,endAt\n")

	return fParams, fSimulation, fHistory, nil
}

func onlineAvg(x, n int, avg float64) float64 {
	return (float64(n)*avg + float64(x)) / float64(n+1)
}
