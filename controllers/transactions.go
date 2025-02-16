package controllers

import (
	"open-banking-expenses/config"
	"open-banking-expenses/models"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTransactions(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, amount, category, date, description FROM transactions")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar transações"})
		return
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.Amount, &t.Category, &t.Date, &t.Description)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar transações"})
			return
		}
		transactions = append(transactions, t)
	}

	c.JSON(http.StatusOK, transactions)
}
