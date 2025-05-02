package ship

import (
	"log/slog"
	"os"

	"github.com/sol-armada/commander/position"
)

type Ship struct {
	Id        string               `json:"id"`
	Name      string               `json:"name"`
	Positions []*position.Position `json:"positions"`
	Image     string               `json:"image"`
	Crew      int                  `json:"crew"`
	SCU       int                  `json:"scu"`
}

var ships = []Ship{}

// //go:embed ships.json
// var shipEmbed embed.FS

func init() {
	s, err := fetch()
	if err != nil {
		slog.Error("failed to fetch ships", "err", err)
		os.Exit(1)
	}
	ships = append(ships, s...)

	// Read the ships.json file from the embedded filesystem
	// s, err := shipEmbed.ReadFile("ships.json")
	// if err != nil {
	// 	slog.Error("failed to read ships.json", "err", err)
	// 	os.Exit(1)
	// }

	// mShips := []any{}
	// if err := json.Unmarshal(s, &mShips); err != nil {
	// 	slog.Error("failed to unmarshal ships.json", "err", err)
	// 	os.Exit(1)
	// }

	// for _, mShip := range mShips {
	// 	ship := &Ship{}
	// 	mShipBytes, _ := json.Marshal(mShip)
	// 	if err := json.Unmarshal(mShipBytes, ship); err != nil {
	// 		slog.Error("failed to unmarshal ships.json", "err", err)
	// 		os.Exit(1)
	// 	}
	// 	ships = append(ships, ship)
	// }
}

func GetShips() []Ship {
	return ships
}
