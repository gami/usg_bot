package tournament

import (
	"usg_bot/entity"
)

// Tournament reporesents a Round-robin Tournament.
type Tournament struct {
	Sheet Sheet
}

// NewTournament is a function to initialize a touranment.
func NewTournament(sheet Sheet) *Tournament {
	return &Tournament{
		Sheet: sheet,
	}
}

func (t *Tournament) CheckIn(team *entity.Team) error {
	t.Sheet.Append("点呼", team.ToRow())
	return nil
}

func (t *Tournament) AddResult(res *entity.Result) error {
	t.Sheet.Append("予選報告", res.ToRow())
	return nil
}
