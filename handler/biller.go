package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"task-golang-api/model"

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

	// URL API Eksternal untuk mendapatkan informasi biller berdasarkan BillerAccountID
	billerAPI := "http://147.139.143.164:8082/api/v1/biller/data/" + payRequest.BillerAccountID

	// Mengambil data biller menggunakan BillerAccountID
	resp, err := http.Get(billerAPI)
	if err != nil {
		log.Println("Error fetching biller data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch biller data"})
		return
	}
	defer resp.Body.Close()

	// // Membaca respons body untuk mendapatkan data biller
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println("Error reading response body:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
	// 	return
	// }

	// // Parse data biller
	// var biller model.Biller
	// if err := json.Unmarshal(body, &biller); err != nil {
	// 	log.Println("Error parsing biller data:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse biller data"})
	// 	return
	// }

	// Menyimpan transaksi ke database
	transaction := model.Transaction{
		BillerAccountID: payRequest.BillerAccountID,
		Amount:          payRequest.Amount,
	}

	// Menyimpan transaksi ke database
	if err := db.Create(&transaction).Error; err != nil {
		log.Println("Error saving transaction:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	// Mengirim respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Payment successful"})
}
