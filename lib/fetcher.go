package mantela_fetcher

import (
	"encoding/json"
	"net/http"
)

func ParseMantela(resp *http.Response) (*Mantela, error) {
	var mantela Mantela
	if err := json.NewDecoder(resp.Body).Decode(&mantela); err != nil {
		return nil, err
	}
	return &mantela, nil
}
