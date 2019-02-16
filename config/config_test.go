package config

import (
	"os"
	"testing"
)

func Test_makeTournamentSheetIDs(t *testing.T) {
	sids := makeTournamentSheetIDs()

	if nr, ok := sids["NR"]; !ok || len(nr) == 0 {
		t.Errorf("faied to make config : env->%v but sid->%v", os.Getenv("TOURNAMENT_SHEET_ID_NR"), sids["NR"])
	}
}
