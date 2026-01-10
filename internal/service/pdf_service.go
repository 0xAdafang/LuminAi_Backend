package service

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(data []byte) (string, error) {
	reader, err := pdf.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	b, err := reader.GetPlainText()
	if err != nil {
		return "", err
	}

	buf.ReadFrom(b)
	return buf.String(), nil
}
