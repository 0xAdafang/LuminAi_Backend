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

func SplitText(text string, chunkSize int, overlap int) []string {
	var chunks []string
	for i := 0; i < len(text); i += chunkSize - overlap {
		end := i + chunkSize
		if end > len(text) {
			end = len(text)
		}
		chunks = append(chunks, text[i:end])
		if end == len(text) {
			break
		}
	}
	return chunks
}
