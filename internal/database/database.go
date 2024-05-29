package database

// Service represents a service that interacts with a database.
type Service interface {
}

type service struct {
	// Persistence Layer
}

var (
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	dbInstance = &service{}
	return dbInstance
}
