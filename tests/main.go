package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/oslokommune/go-gin-tools/pkg/v1/servicetesting"
	authtesting "github.com/oslokommune/go-oidc-middleware/pkg/v1/testing"
	"github.com/tuuturu/trissect-goal-service/pkg/core"
	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"github.com/tuuturu/trissect-goal-service/pkg/core/router"
	"gotest.tools/assert"
)

func newTestEnvironment() (*servicetesting.Environment, error) {
	authTestOptions := authtesting.NewTestTokenOptions()

	discoveryServer := authtesting.CreateTestDiscoveryServer(authTestOptions)
	bearerToken := authtesting.CreateTestToken(authTestOptions)

	discoveryURL, _ := url.Parse(discoveryServer.URL)

	dbPassword := "postgres"

	env, err := servicetesting.NewGinTestEnvironment(servicetesting.CreatePostgresDatabaseBackendOptions(dbPassword), bearerToken)
	if err != nil {
		return nil, fmt.Errorf("creating test environment: %w", err)
	}

	parts := strings.Split(env.GetDatabaseBackendURI(), ":")
	dbURI := parts[0]
	dbPort := parts[1]

	cfg := core.Config{
		DiscoveryURL: discoveryURL,
		Port:         "3000",
		DSN: core.DSN{
			Scheme:       "postgres",
			URI:          dbURI,
			Port:         dbPort,
			DatabaseName: "postgres",
			Username:     "postgres",
			Password:     dbPassword,
		},
	}

	err = cfg.Validate()
	if err != nil {
		_ = env.Teardown()

		return nil, fmt.Errorf("validating config: %w", err)
	}

	env.TestServer = router.NewRouter(cfg)

	return env, nil
}

func goalAsJSONBytes(goal models.Goal) (result []byte) {
	result, _ = json.Marshal(goal)

	return result
}

func JSONBytesAsGoal(data []byte) (result models.Goal) {
	_ = json.Unmarshal(data, &result)

	return result
}

func createGoal(t *testing.T, env *servicetesting.Environment, goal models.Goal) string {
	result, err := env.DoRequest("/goals", http.MethodPost, goalAsJSONBytes(goal))
	assert.NilError(t, err)

	resultGoal := JSONBytesAsGoal(result.Body.Bytes())

	return resultGoal.Id
}
