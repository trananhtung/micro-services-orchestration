package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	ginSwagger "github.com/swaggo/gin-swagger"
	files "github.com/swaggo/files"
	_ "shipping-service/docs" // import docs đã sinh ra
)

// Shipment model
type Shipment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   int       `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("shipping.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Shipment{})

	r := gin.Default()

	r.POST("/shipments", createShipment)
	r.GET("/shipments", listShipments)
	r.GET("/shipments/:id", getShipment)
	r.PUT("/shipments/:id", updateShipment)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.Run(":8080")
}

// createShipment godoc
// @Summary Create a new shipment
// @Accept json
// @Produce json
// @Param shipment body Shipment true "Shipment info"
// @Success 201 {object} Shipment
// @Router /shipments [post]
func createShipment(c *gin.Context) {
	var input Shipment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Status = defaultStatus(input.Status)
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// listShipments godoc
// @Summary List all shipments
// @Produce json
// @Success 200 {array} Shipment
// @Router /shipments [get]
func listShipments(c *gin.Context) {
	var shipments []Shipment
	db.Find(&shipments)
	c.JSON(http.StatusOK, shipments)
}

// getShipment godoc
// @Summary Get shipment by ID
// @Produce json
// @Param id path int true "Shipment ID"
// @Success 200 {object} Shipment
// @Failure 404 {object} map[string]interface{}
// @Router /shipments/{id} [get]
func getShipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var shipment Shipment
	if err := db.First(&shipment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}
	c.JSON(http.StatusOK, shipment)
}

// updateShipment godoc
// @Summary Update shipment status
// @Accept json
// @Produce json
// @Param id path int true "Shipment ID"
// @Param shipment body Shipment true "Shipment info"
// @Success 200 {object} Shipment
// @Failure 404 {object} map[string]interface{}
// @Router /shipments/{id} [put]
func updateShipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var shipment Shipment
	if err := db.First(&shipment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}
	var input Shipment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Status != "" {
		shipment.Status = input.Status
	}
	shipment.UpdatedAt = time.Now()
	db.Save(&shipment)
	c.JSON(http.StatusOK, shipment)
}

func defaultStatus(status string) string {
	if status == "" {
		return "pending"
	}
	return status
} 