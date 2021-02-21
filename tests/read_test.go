package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/tuuturu/trissect-goal-service/pkg/core/models"
	"gotest.tools/assert"
)

func TestReadGoal(t *testing.T) {
	testCases := []struct {
		name string

		with models.Goal
	}{
		{
			name: "Should act as expected when creating and retrieving a goal",

			with: models.Goal{
				Title:     "Add a goal",
				Reasoning: "because goals are awesome",
			},
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

			response, err := env.DoRequest(fmt.Sprintf("/goals/%s", id), http.MethodGet, goalAsJSONBytes(tc.with))
			assert.NilError(t, err)

			// Ensure correct status code
			assert.Equal(t, http.StatusOK, response.Code)

			responseGoal := JSONBytesAsGoal(response.Body.Bytes())

			// Ensure correct ID
			assert.Equal(t, id, responseGoal.Id)

			responseGoal.Id = ""
			responseGoal.Author = ""

			// Ensure correct content
			assert.Equal(t, true, bytes.Equal(goalAsJSONBytes(tc.with), goalAsJSONBytes(responseGoal)))
		})
	}
}
