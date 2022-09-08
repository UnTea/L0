package service

import (
	"context"
	"fmt"
	memory "github.com/Saimunyz/L0/pkg/cache"
	"github.com/UnTea/L0/internal/db"
	"github.com/UnTea/L0/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Storage interface {
	Get(ctx context.Context) ([]model.Data, error)
	Set(ctx context.Context, data model.Data) (string, error)
}

type Service struct {
	dbInstance  Storage
	memoryCache *memory.Cache
	ch          chan model.Data
}

func NewService(dbInstance *pgxpool.Pool, ch chan model.Data) *Service {
	s := &Service{
		dbInstance:  db.NewDatabaseInstance(dbInstance),
		memoryCache: memory.NewCache(),
		ch:          ch,
	}
	go s.Start()
	return s
}

func (s *Service) Start() {
	ctx := context.Background()
	items, err := s.dbInstance.Get(ctx)
	if err != nil {
		log.Printf("Error occurred while getting data from database: %v", err)
	}

	for _, item := range items {
		s.memoryCache.Set(item.OrderUID, item)
	}

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

func (s *Service) GetAllIDs() ([]string, error) {
	ids := s.memoryCache.GetAllIDs()
	if len(ids) < 1 {
		return nil, fmt.Errorf("there is no data yet")
	}

	return ids, nil
}
