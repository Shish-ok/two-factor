package redis_save

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"two-factor-auth/internal/models/redis_dump"
)

const saveFile = "redis_records.json"

type Storage interface {
	PutRecords(context.Context, []redis_dump.RedisRecord)
	GetAllRecords(context.Context) ([]redis_dump.RedisRecord, error)
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

type Service struct {
	storage Storage
}

func (s *Service) RolloutData(ctx context.Context) {
	file, err := os.Open(saveFile)
	if err != nil {
		log.Printf("RolloutData: cant open file: %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("RolloutData: cant read data: %v", err)
		return
	}

	var records []redis_dump.RedisRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Printf("RolloutData: cant unmarshal data: %v", err)
		return
	}

	s.storage.PutRecords(ctx, records)
}

func (s *Service) SaveData(ctx context.Context) {
	records, err := s.storage.GetAllRecords(ctx)
	if err != nil {
		log.Printf("SaveData: cant get all records: %v", err)
		return
	}

	data, err := json.Marshal(records)
	if err != nil {
		log.Printf("SaveData: cant marshal to data: %v", err)
		return
	}

	file, err := os.Create(saveFile)
	if err != nil {
		log.Printf("SaveData: cant create file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		log.Printf("SaveData: write data: %v", err)
	}
}
