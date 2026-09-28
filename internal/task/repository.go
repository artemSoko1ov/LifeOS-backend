package task

type Repository struct {
	storage *Storage
}

func NewRepository(storage *Storage) *Repository {
	return &Repository{
		storage: storage,
	}
}

func (r *Repository) Get() []Task {
	return r.storage.Tasks
}

func (r *Repository) Create(task Task) {
	r.storage.Tasks = append(r.storage.Tasks, task)
}
