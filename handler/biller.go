package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"task-golang-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBillerData(c *gin.Context) {
	// External API URL
	externalAPI := "http://147.139.143.164:8082/api/v1/biller" // replace with the actual API URL

	// Make a GET request to the external API
	resp, err := http.Get(externalAPI)
	if err != nil {
		log.Printf("Failed to fetch data from external API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Parse the JSON response into a slice of Biller structs
	var billers []model.Biller
	if err := json.Unmarshal(body, &billers); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	// Return the billers as JSON
	c.JSON(http.StatusOK, billers)
}

func GetBiller(c *gin.Context) {
	// External API URL
	externalAPI := "http://147.139.143.164:8082/api/v1/biller/data"

	// Make a GET request to the external API
	resp, err := http.Get(externalAPI)
	if err != nil {
		log.Printf("Failed to fetch data from external API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	// Parse the JSON response into a map with biller IDs as keys
	var billers map[string]model.Biller
	if err := json.Unmarshal(body, &billers); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		log.Printf("Response body: %s", string(body)) // Log the response for debugging
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	// Return the billers as JSON
	c.JSON(http.StatusOK, billers)
}
func PayBillerAccount(c *gin.Context, db *gorm.DB) {
	// Mendapatkan parameter dari request
	var payRequest model.PayBillerRequest
	if err := c.ShouldBindJSON(&payRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Membuat payload untuk API eksternal
	externalPayload := map[string]interface{}{
		"biller_id":         payRequest.BillerID,
		"biller_account_id": payRequest.BillerAccountID,
		"amount":            payRequest.Amount,
	}

	// Konversi payload ke JSON
	payloadBytes, err := json.Marshal(externalPayload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to prepare request"})
		return
	}

	// Mengirimkan POST request ke API eksternal
	externalAPI := "http://147.139.143.164:8082/api/v1/biller"
	req, err := http.NewRequest("POST", externalAPI, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("Error creating POST request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending POST request:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process payment"})
		return
	}
	defer resp.Body.Close()

	// Log the raw response to inspect its format
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read external response"})
		return
	}
	log.Println("Raw response from external API:", string(body))

	// Check if response is JSON
	var externalResponse map[string]interface{}
	if jsonErr := json.Unmarshal(body, &externalResponse); jsonErr != nil {
		// If not JSON, return the raw response as an error message
		log.Println("External API returned non-JSON response:", string(body))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from external API", "details": string(body)})
		return
	}

	// Check if the payment was successful in the external API response
	if resp.StatusCode != http.StatusOK {
		// log.Println("Payment unsuccessful or external API error:", externalResponse/)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to complete payment"})
		return
	}

	// Menyimpan transaksi ke database
	transaction := model.Transaction{
		BillerAccountID: payRequest.BillerAccountID,
		Amount:          payRequest.Amount,
		TransactionDate: time.Now(),
	}
	if err := db.Create(&transaction).Error; err != nil {
		log.Println("Error saving transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	// Mengirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}

func GetBillerAccount(c *gin.Context) {
	// Retrieve parameters from URL
	billerID := c.Param("biller_id")
	accountID := c.Param("account_id")

	// Construct external API endpoint
	externalAPI := "http://147.139.143.164:8082/api/v1/biller/" + billerID + "/" + accountID
	log.Printf("Requesting external API at: %s with billerID: %s and accountID: %s", externalAPI, billerID, accountID)

	// Call external API
	resp, err := http.Get(externalAPI)
	if err != nil {
		log.Printf("Failed to call external API: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call external API"})
		return
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("External API returned status code %d. Response body: %s", resp.StatusCode, string(body))
		c.JSON(http.StatusBadGateway, gin.H{"error": "Failed to retrieve valid data from external API", "details": string(body)})
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Log the response body
	log.Printf("Response body from external API: %s", string(body))

	// Parse JSON into struct BillerAccountResponse
	var accountResponse model.BillerAccountResponse
	if err := json.Unmarshal(body, &accountResponse); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
		return
	}

	// Return the nested data as JSON
	c.JSON(http.StatusOK, accountResponse.Data)
}
