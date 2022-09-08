package service

import (
	"context"
	"fmt"
	memorycache "github.com/Saimunyz/L0/pkg/cache"
	"github.com/UnTea/L0/internal/db"
	"github.com/UnTea/L0/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// Storage - interface representing interaction with the database
type Storage interface {
	Get(ctx context.Context) ([]model.Data, error)
	Set(ctx context.Context, data model.Data) (string, error)
}

// Service - stores database instance, in-memory cache
// and a channel for exchanging data with the nats handler
type Service struct {
	dbInstance  Storage
	memoryCache *memorycache.Cache
	ch          chan model.Data
}

// NewService - create new service instance
// and starts goroutine for handling data
func NewService(dbInstance *pgxpool.Pool, ch chan model.Data) *Service {
	s := &Service{
		dbInstance:  db.NewDatabaseInstance(dbInstance),
		memoryCache: memorycache.NewCache(),
		ch:          ch,
	}
	go s.Start()
	return s
}

// Start - restore in-memory cache from DB and
// adds a new data to DB and cache
func (s *Service) Start() {
	ctx := context.Background()

	// getting al data from DB
	items, err := s.dbInstance.Get(ctx)
	if err != nil {
		log.Printf("cannot get data from database due to err: %v", err)
	}

	// restore cache
	for _, item := range items {
		s.memoryCache.Set(item.OrderUID, item)
	}

	// getting data from nats and storing it in the database and cache
	for {
		data, ok := <-s.ch
		if !ok {
			return
		}
		id, err := s.dbInstance.Set(ctx, data)
		if err != nil {
			log.Printf("cannot insert data into database due to err: %v", err)
		} else {
			s.memoryCache.Set(data.OrderUID, data)
			log.Printf("data successefuly added into in-memory cache and database with id: %s", id)
		}
	}
}

// Get - returns element from cache by ID
func (s *Service) Get(id string) (*model.Data, error) {
	data, ok := s.memoryCache.Get(id)
	if !ok {
		return nil, fmt.Errorf("no elements with such id: %s", id)
	}
	result, ok := data.(model.Data)
	if !ok {
		return nil, fmt.Errorf("convertion err")
	}

	return &result, nil
}

// GetAllIDs - returns slice with all IDs in cache
func (s *Service) GetAllIDs() ([]string, error) {
	ids := s.memoryCache.GetAllIDs()
	if len(ids) < 1 {
		return nil, fmt.Errorf("there is no data yet")
	}
	return ids, nil
}
