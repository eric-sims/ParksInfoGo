package apis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"parksInfoGo/models"
)

const baseURL = "https://queue-times.com"

func GetDisneyWaitTimesAllRides(ctx context.Context) (*models.DisneyQueues, error) {
	return getQueues(16, ctx)
}

func GetCaliforniaAdventureWaitTimesAllRides(ctx context.Context) (*models.DisneyQueues, error) {
	return getQueues(17, ctx)
}

func getQueues(id int, ctx context.Context) (*models.DisneyQueues, error) {
	var ret models.QueueTimesResponse

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/parks/%d/queue_times.json", baseURL, id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// cast the body into QueueTimesResponse model
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}

	return transformToDisneyQueues(ret), nil
}

func transformToDisneyQueues(queueTimesResponse models.QueueTimesResponse) *models.DisneyQueues {
	var ret models.DisneyQueues
	for _, land := range queueTimesResponse.Lands {
		for _, ride := range land.Rides {
			ret.QueueTimes = append(ret.QueueTimes, ride)
		}
	}

	return &ret
}
