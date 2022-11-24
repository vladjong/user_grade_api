package fileworker

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/vladjong/user_grade_api/internal/entity"
)

const (
	userIdIter        = 0
	postpaidLimitIter = 1
	sppIter           = 2
	shippingFeeIter   = 3
	returnFeeIter     = 4
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

func (f *WorkerCsv) GetRecord(filename string) (users []entity.UserGrade, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileGZip, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer fileGZip.Close()

	fileCsv := csv.NewReader(fileGZip)
	rec, err := fileCsv.ReadAll()
	if err != nil {
		return nil, err
	}
	users, err = getUsers(rec)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getUsers(rec [][]string) (users []entity.UserGrade, err error) {
	for i, row := range rec {
		if i == 0 {
			continue
		}
		postPaidLimit, err := strconv.Atoi(row[postpaidLimitIter])
		if err != nil {
			return nil, err
		}
		spp, err := strconv.Atoi(row[sppIter])
		if err != nil {
			return nil, err
		}
		shippingFee, err := strconv.Atoi(row[shippingFeeIter])
		if err != nil {
			return nil, err
		}
		returnFee, err := strconv.Atoi(row[returnFeeIter])
		if err != nil {
			return nil, err
		}
		users = append(users, entity.UserGrade{
			UserId:        row[userIdIter],
			PostpaidLimit: postPaidLimit,
			Spp:           spp,
			ShippingFee:   shippingFee,
			ReturnFee:     returnFee,
		})
	}
	return users, nil
}
