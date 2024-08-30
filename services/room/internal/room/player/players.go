package player

import roompb "github.com/moke-game/platform/api/gen/room/api"

type Players struct {
	max     int32
	players map[string]*roompb.Player
}

func NewPlayers(max int32) *Players {
	return &Players{
		max:     max,
		players: make(map[string]*roompb.Player),
	}
}

func (p *Players) AddPlayer(player *roompb.Player) error {
	if int32(len(p.players)) >= p.max {
		return ErrPlayersFull
	}
	p.players[player.Uid] = player
	return nil
}

func (p *Players) RemovePlayer(uid string) bool {
	delete(p.players, uid)
	if len(p.players) == 0 {
		return true
	}
	return false
}

func (p *Players) GetAllPlayers() []*roompb.Player {
	players := make([]*roompb.Player, 0, len(p.players))
	for _, player := range p.players {
		players = append(players, player)
	}
	return players
}
