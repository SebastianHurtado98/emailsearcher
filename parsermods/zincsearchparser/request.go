package zincsearchparser

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

func SendRequestToIndexer(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulk", nil)
	if err != nil {
		return err
	}

	req.Body = io.NopCloser(file)

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Add("Content-Type", "application/x-ndjson")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("No se pudo subir la información a ZincSearch: %s", resp.Status)
	}

	fmt.Println("Se terminó de indexar la información.")
	return nil
}
