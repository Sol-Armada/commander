package ship

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/sol-armada/commander/position"
)

type uexShip struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	NameFull  string  `json:"name_full"`
	UUID      string  `json:"uuid"`
	SCU       float64 `json:"scu"`
	Crew      string  `json:"crew"`
	UrlPhotos string  `json:"url_photos"`
}

func fetch() ([]Ship, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.uexcorp.space/2.0/vehicles", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch ships")
	}

	body := make(map[string]any)
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		panic(err)
	}

	if body["http_code"].(float64) != 200 {
		return nil, errors.New("failed to fetch ships")
	}

	ships := []Ship{}
	if bodyRaw, ok := body["data"].([]any); ok {
		for _, itemJson := range bodyRaw {
			s := uexShip{}
			b, err := json.Marshal(itemJson)
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(b, &s); err != nil {
				return nil, err
			}

			crewSplit := strings.Split(s.Crew, ",")
			crewStr := crewSplit[0]
			if len(crewSplit) > 1 {
				crewStr = crewSplit[1]
			}

			crew, err := strconv.Atoi(strings.TrimSpace(crewStr))
			if err != nil {
				continue
			}

			imageUrl := ""
			if s.UrlPhotos != "" {
				imageUrls := []string{}
				if err := json.Unmarshal([]byte(s.UrlPhotos), &imageUrls); err != nil {
					return nil, err
				}
				if len(imageUrls) > 0 {
					imageUrl = imageUrls[0]
				}
			}

			ship := Ship{
				Id:        s.UUID,
				Name:      s.Name,
				Crew:      crew,
				Image:     imageUrl,
				SCU:       int(s.SCU),
				Positions: []*position.Position{},
			}

			for i := range s.Crew {
				name := "Crew"
				if i == 0 {
					name = "Pilot"
				}
				ship.Positions = append(ship.Positions, &position.Position{
					Id:   i,
					Name: name,
				})
			}

			ships = append(ships, ship)
		}
	}

	return ships, nil
}
