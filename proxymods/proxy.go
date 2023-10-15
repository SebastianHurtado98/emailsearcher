package main


import (
    "fmt"
    "io"
    "log"
    "net/http"
    "strings"
	"encoding/json"
)

type Item struct {
	ID          int    `json:"id"`
	From        string `json:"from"`
	To          []string `json:"to"`
	Subject     string `json:"subject"`
	Message     string `json:"message"`
}

func proxyToZincsearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    message := r.URL.Query().Get("message")
    if message == "" {
        http.Error(w, "El parámetro 'message' es obligatorio.", http.StatusBadRequest)
        return
    }
	data := zincsearchQuery(message)
    jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error generando la respuesta JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func zincsearchQuery(message string) []Item {
	query := `{
        "search_type": "match",
        "query":
        {
            "term": "MESSAGE"
        },
        "from": 0,
        "max_results": 20,
        "_source": ["message", "subject", "to", "from"]
    }`
	query = strings.Replace(query, "MESSAGE", message, 1)
    req, err := http.NewRequest("POST", "http://localhost:4080/api/enron/_search", strings.NewReader(query))
    if err != nil {
        log.Fatal(err)
    }
    req.SetBasicAuth("admin", "Complexpass#123")
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    log.Println(resp.StatusCode)
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    var response map[string]interface{}
    fmt.Println(string(body))
    err = json.Unmarshal(body, &response)
    if err != nil {
        log.Fatal(err)
    }
    if errorMsg, found := response["error"].(string); found && errorMsg != "" {
        log.Println("Error received:", errorMsg)
    }

    items := []Item{}

    results := response["hits"].(map[string]interface{})["hits"]
    for id, hit := range results.([]interface{}) {
        source := hit.(map[string]interface{})["_source"]
        log.Println(source)
        items = append(items, Item{
            ID:          id,
            From:        source.(map[string]interface{})["from"].(string),
            To:          convertInterfaceToStrings(source.(map[string]interface{})["to"]),
            Subject:     source.(map[string]interface{})["subject"].(string),
            Message:     source.(map[string]interface{})["message"].(string),
        })
    }
    return items
}

func convertInterfaceToStrings(interf interface{}) []string {
    interfArray, ok := interf.([]interface{})
    if !ok {
        log.Println("Error: cannot convert 'to' to array")
        return nil
    }
    
    strArray := make([]string, len(interfArray))
    for i, v := range interfArray {
        str, ok := v.(string)
        if !ok {
            log.Println("Error: cannot convert array element to string")
            return nil
        }
        strArray[i] = str
    }

    return strArray
}
