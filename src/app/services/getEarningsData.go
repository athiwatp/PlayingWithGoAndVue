package services

import (
	. "app/clients"
	"app/models"
	"encoding/json"
)

func GetEarningsData(symbol string) (models.EarningsView, error) {
	var body []byte
	var err error

	if body, err = RestClientGet("https://api.iextrading.com/1.0/stock/" + symbol + "/earnings"); err != nil {
		return models.EarningsView{}, err
	}

	var e models.Earnings
	err = json.Unmarshal(body, &e)

	result := models.EarningsView{EarningsData: e, Headers: []string{"Period", "Actual", "Consensus", "Estimated", "Year Ago"}, ColumnWidths: []string{"20%", "20%", "20%", "20%", "20%"}}

	return result, err

}
