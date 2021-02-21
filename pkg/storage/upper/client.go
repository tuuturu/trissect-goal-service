package upper

import (
	"fmt"

	"github.com/imdario/mergo"
	"github.com/tuuturu/trissect-goal-service/pkg/core"
	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

func (c *client) Add(goal models.Goal) (err error) {
	collection := c.Session.Collection(goalTable)

	_, err = collection.Insert(&goal)
	if err != nil {
		return fmt.Errorf("error inserting goal: %w", err)
	}

	return nil
}

func (c *client) Get(id string) (result models.Goal, err error) {
	collection := c.Session.Collection(goalTable)

	condition := db.Cond{"id": id}

	results := collection.Find(condition)

	exists, err := results.Exists()
	if err != nil {
		return result, fmt.Errorf("fetching goal: %w", err)
	}

	if !exists {
		return result, core.StorageErrorNotFound
	}

	err = results.One(&result)
	if err != nil {
		return result, fmt.Errorf("finding goal: %w", err)
	}

	return result, nil
}

func (c *client) GetAll(filter core.StorageFilter) (goals []models.Goal, err error) {
	condition := filterToDBCond(filter)

	collection := c.Session.Collection(goalTable)

	err = collection.Find(condition).All(&goals)
	if err != nil {
		return nil, fmt.Errorf("fetching goals: %w", err)
	}

	return goals, nil
}

func (c *client) Update(update models.Goal) (updatedGoal models.Goal, err error) {
	collection := c.Session.Collection(goalTable)

	var originalGoal models.Goal

	condition := db.Cond{"id": update.Id}

	result := collection.Find(condition)

	exists, err := result.Exists()
	if err != nil {
		return updatedGoal, fmt.Errorf("fetching goal: %w", err)
	}

	if !exists {
		return updatedGoal, core.StorageErrorNotFound
	}

	err = result.One(&originalGoal)
	if err != nil {
		return updatedGoal, fmt.Errorf("fetching original goal: %w", err)
	}

	err = mergo.Merge(&originalGoal, update, mergo.WithOverride)
	if err != nil {
		return updatedGoal, fmt.Errorf("merging updated with original goal: %w", err)
	}

	err = collection.UpdateReturning(&originalGoal)
	if err != nil {
		return updatedGoal, fmt.Errorf("updating goal: %w", err)
	}

	return originalGoal, nil
}

func (c *client) Delete(id string) (err error) {
	collection := c.Session.Collection(goalTable)

	condition := db.Cond{"id": id}

	result := collection.Find(condition)

	exists, err := result.Exists()
	if err != nil {
		return fmt.Errorf("fetching goal: %w", err)
	}

	if !exists {
		return core.StorageErrorNotFound
	}

	err = result.Delete()
	if err != nil {
		return fmt.Errorf("deleting goal with ID %s: %w", id, err)
	}

	return nil
}

func (c *client) Open() error {
	sess, err := postgresql.Open(c.connectionURL)
	if err != nil {
		return fmt.Errorf("error connecting to Postgres: %w", err)
	}

	c.Session = sess

	return c.setup()
}

func (c *client) Close() error {
	return c.Session.Close()
}

func (c *client) setup() error {
	sql := c.Session.SQL()

	_, err := sql.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id text primary key,
		author text not null,
		parent text,
		title text not null,
		reasoning text not null,
		complete bool not null
	)`, goalTable))
	if err != nil {
		return fmt.Errorf("error creating tables: %w", err)
	}

	return nil
}

func filterToDBCond(filter core.StorageFilter) (condition db.Cond) {
	condition = db.Cond{}

	if filter.Author != nil {
		condition["author"] = *filter.Author
	}

	if filter.Complete != nil {
		condition["complete"] = *filter.Complete
	}

	return condition
}
