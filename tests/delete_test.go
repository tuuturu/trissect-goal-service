package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"gotest.tools/assert"
)

func TestDeleteGoal(t *testing.T) {
	testCases := []struct {
		name string

		with models.Goal

		expectStatus int
	}{
		{
			name: "Should return correct status code when deleting a goal",

			with: models.Goal{
				Title:     "create a goal",
				Reasoning: "goals are awesome",
			},

			expectStatus: http.StatusNoContent,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			env, err := newTestEnvironment()
			assert.NilError(t, err)

			defer func() {
				_ = env.Teardown()
			}()

			id := createGoal(t, env, tc.with)
			targetURL := fmt.Sprintf("/goals/%s", id)

			response, err := env.DoRequest(targetURL, http.MethodDelete, nil)
			assert.NilError(t, err)

			// Ensure correct status code
			assert.Equal(t, tc.expectStatus, response.Code)

			response, err = env.DoRequest(targetURL, http.MethodGet, nil)
			assert.NilError(t, err)

			// Ensure goal being gone
			assert.Equal(t, http.StatusNotFound, response.Code)
		})
	}
}
