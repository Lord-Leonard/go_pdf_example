package main

import (
	"github.com/Lord-Leonard/go_pdf_example/Ausgabesteuerung"
	"github.com/Lord-Leonard/go_pdf_example/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/rechnung", func(c *gin.Context) {
		rechnungsdaten := model.Rechnung{}

		if err := c.ShouldBindJSON(&rechnungsdaten); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		pdfLocation := Ausgabesteuerung.RechnungErstellen(rechnungsdaten, c)

		c.JSON(http.StatusOK, gin.H{
			pdfLocation: pdfLocation,
		})
	})
	err := router.Run()

	if err != nil {
		return
	}
}
