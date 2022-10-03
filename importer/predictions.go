package importer

import (
	"context"
	"github.com/sajari/regression"
	"metar.gg/ent"
	"metar.gg/ent/metar"
	"metar.gg/ent/weatherstation"
	"time"
)

func (i *NoaaWeatherImporter) MakeNextImportPrediction(ctx context.Context, stationID string, currentImportTime *time.Time) (*time.Time, error) {
	r := new(regression.Regression)

	r.SetObserved("Delta since last observation time")
	r.SetVar(0, "Sequence number")

	// Add some data points.
	result, err := i.db.Metar.Query().Select(metar.FieldImportTime).Where(metar.HasStationWith(weatherstation.StationID(stationID)), metar.ImportTimeGT(time.Now().Add(time.Hour*-4))).Order(ent.Asc(metar.FieldImportTime)).All(ctx)
	if err != nil {
		return nil, err
	}

	if len(result) < 2 {
		return nil, nil
	}

	// Add the current time as the last observation time, because the newest one isn't persisted yet
	result = append(result, &ent.Metar{ImportTime: *currentImportTime})

	sequenceNumber := 0

	for i, m := range result {
		if i == 0 {
			continue
		}

		delta := m.ImportTime.Sub(result[i-1].ImportTime)

		r.Train(regression.DataPoint(float64(delta), []float64{float64(sequenceNumber)}))
		sequenceNumber++
	}

	err = r.Run()
	if err != nil {
		return nil, err
	}

	// Make a prediction for the next import.
	prediction, err := r.Predict([]float64{float64(sequenceNumber + 1)})
	if err != nil {
		return nil, err
	}

	// Add the prediction to the last import time.
	nextImportTime := result[len(result)-1].ImportTime.Add(time.Duration(prediction))

	return &nextImportTime, nil
}
