package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"gotest.tools/assert"
)

func TestCreateGoal(t *testing.T) {
	testCases := []struct {
		name string

		with models.Goal

		expectStatus int
	}{
		{
			name: "Should return correct status code on successful goal creation",

			with: models.Goal{
				Title:     "Conquer the world",
				Reasoning: "Would be awesome",
			},

			expectStatus: http.StatusCreated,
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

			response, err := env.DoRequest("/goals", http.MethodPost, goalAsJSONBytes(tc.with))
			assert.NilError(t, err)

			// Ensure correct status code
			assert.Equal(t, tc.expectStatus, response.Code)

			resultGoal := JSONBytesAsGoal(response.Body.Bytes())

			// Ensure the goal has been given an ID
			assert.Equal(t, true, resultGoal.Id != "")

			resultGoal.Id = ""
			resultGoal.Author = ""

			// Ensure content integrity
			assert.Equal(t, true, bytes.Equal(goalAsJSONBytes(tc.with), goalAsJSONBytes(resultGoal)))
		})
	}
}
