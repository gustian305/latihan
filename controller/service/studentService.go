package services

import (
	"latihan/database"
	"latihan/dto"
	"latihan/models"
)

// StudentService mendefinisikan kontrak untuk pengelolaan data mahasiswa.
type StudentService interface {
	Create(input dto.RequestStudent) (dto.ResponseStudent, error)
	FindAll() ([]dto.ResponseStudent, error)
}

// studentService adalah implementasi konkret dari StudentService.
type studentService struct{}

// NewStudentService membuat instance baru untuk StudentService.
func NewStudentService() StudentService {
	return &studentService{}
}

func (s *studentService) Create(input dto.RequestStudent) (dto.ResponseStudent, error) {
	student := models.Student{
		Name:  input.Name,
		Class: input.Class,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		return dto.ResponseStudent{}, err
	}

	return dto.ResponseStudent{
		ID:    student.ID,
		Name:  student.Name,
		Class: student.Class,
	}, nil
}

func (s *studentService) FindAll() ([]dto.ResponseStudent, error) {
	var students []models.Student

	if err := database.DB.Find(&students).Error; err != nil {
		return nil, err
	}

	responses := make([]dto.ResponseStudent, 0, len(students))
	for _, student := range students {
		responses = append(responses, dto.ResponseStudent{
			ID:    student.ID,
			Name:  student.Name,
			Class: student.Class,
		})
	}

	return responses, nil
}
