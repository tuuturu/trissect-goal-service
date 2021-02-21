package core

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
)

/*
 * Config
 */
type Config struct {
	Port string

	DSN DSN
}

/*
 * Storage
 */
var (
	StorageErrorNotFound = errors.New("not found")
)

type StorageFilter struct {
	Author   *string
	Complete *bool
}

type StorageClient interface {
	Open() (err error)
	Close() (err error)

	Add(goal models.Goal) (id string, err error)
	Get(id string) (result models.Goal, err error)
	GetAll(filter StorageFilter) (result []models.Goal, err error)
	Update(goal models.Goal) (updatedGoal models.Goal, err error)
	Delete(id string) (err error)
}

type DSN struct {
	Scheme       string
	Username     string
	Password     string
	URI          string
	Port         string
	DatabaseName string
}

/*
 * Router
 */
type HandlerFuncGenerator func(storage StorageClient) gin.HandlerFunc

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFuncGenerator HandlerFuncGenerator
}

// Routes is the list of the generated Route.
type Routes []Route
