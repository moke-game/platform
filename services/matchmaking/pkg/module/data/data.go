package data

type MatchResult struct {
	MatchRoomId    string                `json:"match_room_id"`
	GameMode       int32                 `json:"game_mode"`
	PlayId         int32                 `json:"play_id"`
	MapId          int32                 `json:"map_id"`
	Members        map[string]int32      `json:"members"` //key uid val 阵营ID
	Robots         map[int32][]int32     `json:"robots"`  //key 阵营ID val 机器人战力 val长度为机器人个数
	Players        map[string]PlayerData `json:"players"`
	HallUrl        string                `json:"hall_url"`
	HallID         string                `json:"hall_id"`
	BattleRoomAddr string                `json:"battle_room_addr"`
	IsFirstEnter   bool                  `json:"is_first_enter"`
}

type PlayerData struct {
	Uid          string            `json:"uid"`
	Nickname     string            `json:"nickname"`
	Avatar       string            `json:"avatar"`
	HeroId       int32             `json:"hero_id"`
	SkinId       int32             `json:"skin_id"`
	HeroLevel    int32             `json:"hero_level"`
	Attribute    map[int32]float64 `json:"attribute"`
	PetProfileId int64             `json:"pet_profile_id"`
	HeroCups     int32             `json:"hero_cups"`
	PetSkill     map[int32]int32   `json:"pet_skill"`
	IsAgain      bool              `json:"is_again"`
}
