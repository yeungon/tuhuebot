package xata

// import (
// 	"context"
// 	"encoding/json"
// 	"log"

// 	"github.com/xataio/xata-go/xata"
// )

// type Response struct {
// 	Meta    Meta     `json:"meta"`
// 	Records []Record `json:"records"`
// }

// type Meta struct {
// 	Page Page `json:"page"`
// }

// type Page struct {
// 	Cursor string `json:"cursor"`
// 	More   bool   `json:"more"`
// 	Size   int    `json:"size"`
// }

// type Record struct {
// 	ID       string `json:"id"`
// 	Question string `json:"question"`
// 	Answer   string `json:"answer"`
// 	Xata     Xata   `json:"xata"`
// }

// type Xata struct {
// 	CreatedAt string `json:"createdAt"`
// 	UpdatedAt string `json:"updatedAt"`
// 	Version   int    `json:"version"`
// }

// func Query() Response {
// 	// Query the "qa" table (database and branch are passed here)
// 	qa, err := searchClient.Query(context.TODO(), xata.QueryTableRequest{
// 		BranchRequestOptional: xata.BranchRequestOptional{
// 			DatabaseName: xata.String("tuhuebot"),
// 			BranchName:   xata.String("main"),
// 		},
// 		TableName: "qa",
// 	})
// 	if err != nil {
// 		log.Fatalf("Error querying the qa table: %v", err)
// 	}

// 	// // Convert data to a readable JSON format
// 	qa_JSON, err := json.MarshalIndent(qa, "", "  ")
// 	if err != nil {
// 		log.Fatalf("Failed to marshal qa response: %v", err)
// 	}

// 	var response Response
// 	err = json.Unmarshal([]byte(qa_JSON), &response)
// 	if err != nil {
// 		log.Fatalf("Error unmarshaling JSON: %v", err)
// 	}

// 	return response
// 	// // Iterate through records and print the id, question, and answer
// 	// for _, record := range response.Records {
// 	// 	fmt.Printf("ID: %s\n", record.ID)
// 	// 	fmt.Printf("Question: %s\n", record.Question)
// 	// 	fmt.Printf("Answer: %s\n\n", record.Answer)
// 	// }
// }
