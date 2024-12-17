package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestEntity struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&TestEntity{})
	db.Exec("DELETE FROM test_entities")
	return db
}

func TestBaseRepository_Create(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	entity := &TestEntity{Name: "Test Name"}
	err := repo.Create(entity)

	assert.NoError(t, err)
	assert.NotZero(t, entity.ID)
}

func TestBaseRepository_FindByID(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	result, err := repo.FindByID(entity.ID)

	assert.NoError(t, err)
	assert.Equal(t, entity.Name, result.Name)
}

func TestBaseRepository_FindAll(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})

	results, err := repo.FindAll()

	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestBaseRepository_Update(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	entity := &TestEntity{Name: "Old Name"}
	repo.Create(entity)

	entity.Name = "New Name"
	err := repo.Update(entity)

	assert.NoError(t, err)

	result, _ := repo.FindByID(entity.ID)
	assert.Equal(t, "New Name", result.Name)
}

func TestBaseRepository_Delete(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	err := repo.Delete(entity.ID)

	assert.NoError(t, err)

	result, err := repo.FindByID(entity.ID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestBaseRepository_Count(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})

	count, err := repo.Count()

	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}

func TestBaseRepository_Paginate(t *testing.T) {
	db := setupTestDB()
	repo := NewBaseRepository[TestEntity](db)

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})
	repo.Create(&TestEntity{Name: "Test 3"})
	repo.Create(&TestEntity{Name: "Test 4"})

	results, err := repo.Paginate(2, 2)

	assert.NoError(t, err)
	assert.Len(t, results, 2)
}
