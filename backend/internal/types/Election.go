package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Election struct {
	ElectionID int    `json:"election_id"` // PK
	Start      string `json:"start"`
	End        string `json:"end"`
	Location   string `json:"location"`
}

type ElectionState struct {
	ElectionID int              `json:"election_id"`
	Start      string           `json:"start"`
	End        string           `json:"end"`
	Location   string           `json:"location"`
	Result     []CandidateVotes `json:"result"`
}

type CandidateVotes struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Party       string `json:"party"`
	PartyColour string `json:"party_colour"`
	Votes       int    `json:"votes"`
}

type ElectionRequest struct {
	Start      string             `json:"start"`
	End        string             `json:"end"`
	Location   string             `json:"location"`
	Candidates []CandidateRequest `json:"candidates"`
}

func (req ElectionRequest) Validate() error {
	if req.Start == "" {
		return fmt.Errorf("start is empty")
	}

	if req.End == "" {
		return fmt.Errorf("end is empty")
	}

	if req.Location == "" {
		return fmt.Errorf("location is empty")
	}

	startDate, err := StringToDate(req.Start)
	if err != nil {
		return fmt.Errorf("failed to decode start date: %w", err)
	}

	endDate, err := StringToDate(req.End)
	if err != nil {
		return fmt.Errorf("failed to decode end date: %w", err)
	}

	if startDate.Unix() > time.Now().Unix() {
		return fmt.Errorf("attempted to create an election that starts in the past")
	}

	if endDate.Unix() < startDate.Unix() {
		return fmt.Errorf("ending date of election is before start date")
	}

	for _, candidate := range req.Candidates {
		if err := candidate.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func StringToDate(date string) (time.Time, error) {
	splitDate := strings.Split(date, "-")
	if len(splitDate) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	year, err := StringToInt(splitDate[0])
	if err != nil {
		return time.Time{}, err
	}

	month, err := StringToInt(splitDate[1])
	if err != nil {
		return time.Time{}, err
	}

	day, err := StringToInt(splitDate[2])
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func StringToInt(value string) (int, error) {
	year, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return -1, fmt.Errorf("invalid number in date")
	}

	return int(year), nil
}
