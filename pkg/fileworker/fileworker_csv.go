package fileworker

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/vladjong/user_grade_api/internal/entity"
)

type WorkerCsv struct{}

func (f *WorkerCsv) Record(records []entity.UserGrade, header []string) (string, error) {
	filename := "data/backup.csv.gz"

	outputFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	zipWriter := gzip.NewWriter(outputFile)
	csvWriter := csv.NewWriter(zipWriter)

	if err := csvWriter.Write(header); err != nil {
		return "", err
	}
	for _, record := range records {
		var csvRow []string
		csvRow = append(csvRow, record.UserId, fmt.Sprint(record.PostpaidLimit), fmt.Sprint(record.Spp), fmt.Sprint(record.ShippingFee), fmt.Sprint(record.ReturnFee))
		if err := csvWriter.Write(csvRow); err != nil {
			return "", err
		}
	}
	csvWriter.Flush()
	zipWriter.Flush()
	zipWriter.Close()
	return filename, nil
}
