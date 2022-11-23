package fileworker

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/vladjong/user_grade_api/internal/entity"
)

type WorkerCsv struct{}

func (f *WorkerCsv) Record(records []entity.UserGrade, header []string) (string, error) {
	filename := "data/backup.csv"
	outputFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	if err := writer.Write(header); err != nil {
		return "", err
	}
	for _, record := range records {
		var csvRow []string
		csvRow = append(csvRow, record.UserId, fmt.Sprint(record.PostpaidLimit), fmt.Sprint(record.ReturnFee), fmt.Sprint(record.ShippingFee), fmt.Sprint(record.Spp))
		if err := writer.Write(csvRow); err != nil {
			return "", err
		}
	}
	return filename, nil
}
