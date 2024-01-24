package algorithm

import (
	"fmt"
	"log"

	"github.com/sajari/regression"
)

func NewRegression(x, y []float64) (*regression.Regression, error) {
	// Create a new regression instance
	r := new(regression.Regression)
	r.SetObserved("Sales")
	r.SetVar(0, "Time")

	// Add data points
	for i, xi := range x {
		log.Printf("%+v", []float64{xi})
		r.Train(regression.DataPoint(y[i], []float64{xi}))
	}

	// Run the regression
	r.Run()

	return r, nil
}

func PredictedSales(r *regression.Regression, period int) (float64, error) {
	// Make predictions``
	predictedSales, err := r.Predict([]float64{float64(period)})
	if err != nil {
		fmt.Println("Error predicting:", err)
		return 0, err
	}

	return predictedSales, nil
}
