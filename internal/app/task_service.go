package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	GetForUser(uId uint64) ([]domain.Task, error)
	//1
	FindById(id uint64) (domain.Task, error)
	//2
	Update(t domain.Task) (domain.Task, error)
	//3
	Delete(id uint64) error
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

// 1
func (s taskService) FindById(id uint64) (domain.Task, error) {
	task, err := s.taskRepo.FindById(id)
	if err != nil {
		log.Printf("TaskService -> FindById: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

// 2
func (s taskService) Update(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Update(t)
	if err != nil {
		log.Printf("TaskService -> Update: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

// 3
func (s taskService) Delete(id uint64) error {
	err := s.taskRepo.Delete(id)
	if err != nil {
		log.Printf("TaskService -> Delete: %s", err)
		return err
	}
	return nil
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService -> Save: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) GetForUser(uId uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.GetByUserId(uId)
	if err != nil {
		log.Printf("TaskService -> GetForUser: %s", err)
		return nil, err
	}
	return tasks, nil
}
