package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"gotest.tools/assert"
)

func TestUpdateGoal(t *testing.T) {
	testCases := []struct {
		name string

		withGoal   models.Goal
		withUpdate models.Goal

		expect     models.Goal
		expectCode int
	}{
		{
			name: "Should return updated data correctly",

			withGoal: models.Goal{
				Title:     "add a goal",
				Reasoning: "goals are cool",
			},
			withUpdate: models.Goal{
				Reasoning: "goals are super cool",
				Complete:  true,
			},

			expect: models.Goal{
				Title:     "add a goal",
				Reasoning: "goals are super cool",
				Complete:  true,
			},
			expectCode: http.StatusOK,
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

			id := createGoal(t, env, tc.withGoal)

			goalURL := fmt.Sprintf("/goals/%s", id)

			response, err := env.DoRequest(goalURL, http.MethodPatch, goalAsJSONBytes(tc.withUpdate))
			assert.NilError(t, err)

			// Ensure correct status code
			assert.Equal(t, tc.expectCode, response.Code)

			response, err = env.DoRequest(goalURL, http.MethodGet, nil)
			assert.NilError(t, err)

			responseGoal := JSONBytesAsGoal(response.Body.Bytes())

			responseGoal.Id = ""
			responseGoal.Author = ""

			// Ensure content was updated
			assert.Equal(t, tc.expect, responseGoal)
		})
	}
}
