package task

import (
	"context"
	"errors"
	"strings"
	"sync"
)

var ErrNotFound = errors.New("task not found")

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type Repository interface {
	List(context.Context) []Task
	Create(context.Context, string) (Task, error)
}

type MemoryRepository struct {
	mu     sync.RWMutex
	nextID int
	tasks  []Task
}

func NewMemoryRepository() *MemoryRepository { return &MemoryRepository{nextID: 1} }
func (r *MemoryRepository) List(context.Context) []Task {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Task, len(r.tasks))
	copy(out, r.tasks)
	return out
}
func (r *MemoryRepository) Create(_ context.Context, title string) (Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Task{}, errors.New("title is required")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	t := Task{ID: r.nextID, Title: title}
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t, nil
}

type Service struct{ repo Repository }

func NewService(repo Repository) Service          { return Service{repo: repo} }
func (s Service) List(ctx context.Context) []Task { return s.repo.List(ctx) }
func (s Service) Create(ctx context.Context, title string) (Task, error) {
	return s.repo.Create(ctx, title)
}
