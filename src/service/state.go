package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/VivekSheregar/address_API/models"
	"gorm.io/gorm"
)

type StateService interface {
	Create(models.Mst_State) (*models.Mst_State, error)
	List(string, string) ([]models.Mst_State, error)
	Get(string) (*models.Mst_State, error)
	Update(string, models.Mst_State) (*models.Mst_State, error)
	Delete(string) (string, error)
	DeleteAll() (string, error)
}
type stateService struct {
	conn *gorm.DB
}

type errorString struct {
	errorstring string
}

func (e *errorString) Error() string {
	return e.errorstring
}

func errorData(text string) error {
	return &errorString{
		errorstring: text,
	}
}
func NewStateService(conn *gorm.DB) StateService {
	return &stateService{
		conn: conn,
	}
}

func (s *stateService) Create(req models.Mst_State) (*models.Mst_State, error) {
	req.Name = strings.Trim(req.Name, " ")
	if req.Name == "" {
		return nil, errorData("State name cannot be empty")
	}
	result := s.conn.Create(&req)
	if result.Error != nil {
		return nil, result.Error
	}
	return &req, nil
}

func (s *stateService) List(page string, size string) ([]models.Mst_State, error) {
	stateList := []models.Mst_State{}
	if page == "" || size == "" {
		return stateList, errorData("Page and Size cannot be empty")
	}

	totalCount, err := strconv.ParseInt(size, 6, 12)
	if err != nil {
		return nil, errorData("Invalid Datatype for size")
	}
	toStart, err := strconv.ParseInt(page, 6, 12)
	if err != nil {
		return nil, errorData("Invalid Datatype for page")
	}
	offSet := (toStart - 1) * totalCount
	limit := toStart * totalCount
	fmt.Println(limit, offSet)
	result := s.conn.Limit(int(limit)).Offset(int(offSet)).Find(&stateList)
	if result.Error != nil {
		return nil, result.Error
	}
	return stateList, nil
}

func (s *stateService) Get(ID string) (*models.Mst_State, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return nil, errorData("Invalid datatype for ID")
	}
	stateData := models.Mst_State{}
	result := s.conn.First(&stateData, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stateData, nil
}

func (s *stateService) Update(ID string, req models.Mst_State) (*models.Mst_State, error) {
	Id, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return nil, errorData("Invalid datatype for ID")
	}
	req.ID = Id
	req.Name = strings.Trim(req.Name, " ")
	if req.Name == "" {
		return nil, errorData("State Name cannot be empty")
	}
	// result := s.conn.Save((&req))
	result := s.conn.Model(&req).Where("id = ?", ID).Update("name", &req.Name)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errorData("ID Not Found")
	}
	return &req, nil
}

func (s *stateService) Delete(ID string) (string, error) {
	_, err := strconv.ParseInt(ID, 0, 8)
	if err != nil {
		return "", errorData("Invalid datatype for ID")
	}
	result := s.conn.Delete(&models.Mst_State{}, ID)
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", errorData("ID Not Found")
	}
	return "deleted", nil
}

func (s *stateService) DeleteAll() (string, error) {
	result := s.conn.Exec("delete from mst_states")
	if result.Error != nil {
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", errorData("No Data Found")
	}
	return "deleted", nil
}
