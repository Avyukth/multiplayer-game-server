package models

type gameModes struct {
	BattleRoyal    int32 `bson:"battle_royal"`
	TeamDeathMatch int32 `bson:"team_death_match"`
	CaptureTheFlag int32 `bson:"capture_the_flag"`
}

type GameStat struct {
	AreaCode  int32     `bson:"area_code"`
	GameModes gameModes `bson:"game_modes"`
}
