package services

import (
	"onez19/config"
	"onez19/models"
)

func GetActivitiesBySectionAndWorkspace(sectionID int, workspaceID string) ([]models.Activity, error) {
	var activities []models.Activity

	query := `
		SELECT a.id, a.name, a.description, a.start_date, a.end_date
		FROM activity AS a
		WHERE a.section_id = ? AND a.workspace_id = ?
	`

	rows, err := config.DB.Query(query, sectionID, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ID, &activity.Name, &activity.Description, &activity.StartDate, &activity.EndDate); err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}

	// ตรวจสอบข้อผิดพลาดจาก rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return activities, nil
}

func CreateActivity(activity models.Activity) error {
	_, err := config.DB.Exec("INSERT INTO activity (name, description, start_date, end_date, section_id, workspace_id) VALUES (?, ?, ?, ?, ?, ?)",
		activity.Name, activity.Description, activity.StartDate, activity.EndDate, activity.SectionID, activity.WorkspaceID)
	return err
}

func MoveActivity(activityID int, newSectionID int) error {
	_, err := config.DB.Exec("UPDATE activity SET section_id = ? WHERE id = ?", newSectionID, activityID)
	return err
}

func EditActivity(activity models.Activity) error {
	_, err := config.DB.Exec("UPDATE activity SET name = ?, description = ?, start_date = ?, end_date = ? WHERE id = ?",
		activity.Name, activity.Description, activity.StartDate, activity.EndDate, activity.ID)
	return err
}
