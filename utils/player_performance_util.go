package utils

import (
	"github.com/gbart0198/bball-tracker-api/db"
	"github.com/gbart0198/bball-tracker-api/storage"
)

func HandlePlayerGoalUpdates(playerPerformance *db.PlayerPerformance, repo storage.Storage) {
	// lets get all of the goals for the player with the associated drill id
	// only want the goals that haven't been completed
	params := db.GetGoalsByPlayerAndDrillParams{
		PlayerID: playerPerformance.PlayerID,
		DrillID:  playerPerformance.DrillID,
	}
	goals := repo.GetGoalsByPlayerAndDrill(params)

	// TODO: Find some way to deduplicate between goals that have the same drill but differnet number of attempts between runs
	// e.g goal is to make 15/25 but the performance had 25/30 - obviously that should pass based on percentages.
	// on the flip side, if the goal is 15/25 and the performance was 15/30, that should fail.

	for _, goal := range goals {
		// get the goal type - if it is a completion goal, increment the current value
		// if it is a performance goal, check if the performance value is higher than the current value (max value recorded)
		// if the goal model was updated, update it in the database
		goalUpdated := false
		if goal.Category == "completion" {
			goal.CurrentValue.Int32++
			if goal.CurrentValue.Int32 >= goal.GoalValue {
				goal.Completed = true
			}
			goalUpdated = true
		} else if goal.Category == "performance" {
			if playerPerformance.Successful.Int32 > goal.CurrentValue.Int32 {
				goal.CurrentValue = playerPerformance.Successful
				if goal.CurrentValue.Int32 >= goal.GoalValue {
					goal.Completed = true
				}
				goalUpdated = true
			}
		}

		// should probably propogate the error up to the caller if the update fails
		if goalUpdated {
			updateParams := db.UpdatePlayerGoalParams{
				PlayerGoalID:    goal.PlayerGoalID,
				PlayerID:        goal.PlayerID,
				DrillID:         goal.DrillID,
				CurrentValue:    goal.CurrentValue,
				GoalValue:       goal.GoalValue,
				GoalCategoryID:  goal.GoalCategoryID,
				GoalName:        goal.GoalName,
				GoalDescription: goal.GoalDescription,
				Completed:       goal.Completed,
			}

			repo.UpdatePlayerGoal(updateParams)
		}
	}

}
