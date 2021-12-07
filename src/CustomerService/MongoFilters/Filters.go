package MongoFilters

import (
	"Api.Calisma/src/Common/Filters/FilterSettings"
	"Api.Calisma/src/Common/Filters/PaginationSettings"
	"Api.Calisma/src/CustomerService/Constants"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

func GetPaginator(ctx echo.Context) *options.FindOptions {
	return PaginationSettings.NewSettings().
		SetPagingKey(Constants.PaginationPageKey).
		SetOffSetKey(Constants.PaginationOffsetKey).
		SetMaxPerPage(Constants.PaginationPerPageMax).
		SetDefaultPerPage(Constants.PaginationPerPageDefault).
		SetShowAllKey(Constants.PaginationShowAllKey).
		GetPaginator(ctx.QueryParams())
}

func GetSearchFilter(ctx echo.Context) *bson.M {
	settings := FilterSettings.NewSettings()
	params := ctx.QueryParams()
	acceptableKeys := strings.Split(Constants.AcceptableSearchTexts, ",")
	for _, key := range acceptableKeys {
		if v := params.Get(key); v != "" {
			settings.AddField(key, v, "i")
		}
	}
	return settings.GetSearchFilter()
}
