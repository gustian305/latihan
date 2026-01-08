package handlers

import (
	"latihan/controller/service"
	"latihan/dto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StudentHandler bertanggung jawab menangani request HTTP terkait data mahasiswa.
type StudentHandler struct {
	service services.StudentService
}

// NewStudentHandler adalah constructor untuk menginisialisasi StudentHandler dengan service yang dibutuhkan.
func NewStudentHandler(s services.StudentService) *StudentHandler {
	return &StudentHandler{service: s}
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var input dto.RequestStudent

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("[ERROR] CreateStudent : input tidak valid:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	result, err := h.service.Create(input)
	if err != nil {
		log.Println("[ERROR] CreateStudent : gagal membuat data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": result})
}
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	result, err := h.service.FindAll()
	if err != nil {
		log.Println("[ERROR] GetAllStudents : gagal mengambil data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
