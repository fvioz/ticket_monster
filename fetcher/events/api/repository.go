package api

import (
	"encoding/xml"
	"fetcher/libs"
	"io"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetEvents() ([]BasePlan, error) {
	xmlEvents, err := r.GetXmlEvents()
	if err != nil {
		return nil, err
	}

	var planList PlanList
	xml.Unmarshal(xmlEvents, &planList)

	return planList.Output.BasePlans, nil
}

func (r *Repository) GetXmlEvents() ([]byte, error) {
	logger := libs.LoggerInstance()
	client := libs.NewHttpClient()

	resp, err := client.Get("https://provider.code-challenge.feverup.com/api/events")
	if err != nil {
		logger.Error("Failed to fetch events", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
