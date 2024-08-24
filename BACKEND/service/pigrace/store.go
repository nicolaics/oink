package pigrace

import (
	"database/sql"
	"fmt"

	"github.com/nicolaics/oink/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPigRaceDataByID(id int) (*types.PigRace, error) {
	rows, err := s.db.Query("SELECT * FROM pig_race WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	pigRaceData := new(types.PigRace)

	for rows.Next() {
		pigRaceData, err = scanRowIntoPigRaceData(rows)

		if err != nil {
			return nil, err
		}
	}

	if pigRaceData.ID == 0 {
		return nil, fmt.Errorf("pig race data not found")
	}

	return pigRaceData, nil
}

func (s *Store) UpdateFinalDistance(pigRaceData *types.PigRace, distance float64) error {
	_, err := s.db.Exec("UPDATE pig_race JOIN users ON pig_race.user_id = users.id SET final_distance_to_goal = ? WHERE users.id = ? ",
							distance, pigRaceData.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Store) UpdatePigStamina(pigRaceData *types.PigRace, stamina float64) error {
	_, err := s.db.Exec("UPDATE pig_race JOIN users ON pig_race.user_id = users.id SET stamina = ? WHERE users.id = ? ",
							(pigRaceData.PigStamina + stamina), pigRaceData.ID)
	if err != nil {
		return err
	}
	
	return nil
}

func scanRowIntoPigRaceData(rows *sql.Rows) (*types.PigRace, error) {
	pigRaceData := new(types.PigRace)

	err := rows.Scan(
		&pigRaceData.ID,
		&pigRaceData.UserID,
		&pigRaceData.PigStamina,
		&pigRaceData.FinalDistanceToGoal,
	)

	if err != nil {
		return nil, err
	}

	return pigRaceData, nil
}
