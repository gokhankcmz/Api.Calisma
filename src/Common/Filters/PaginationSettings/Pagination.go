package PaginationSettings

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"strconv"
)

type PaginationSettings struct {
	DefaultPerPage int64
	MaxPerPage     int64
	OffsetKey      string
	PageKey        string
	ShowAllKey     string
}

func NewSettings() *PaginationSettings {
	return &PaginationSettings{
		DefaultPerPage: 10,
		MaxPerPage:     50,
		OffsetKey:      "offset",
		PageKey:        "page",
		ShowAllKey:     "",
	}
}

func (ps *PaginationSettings) SetMaxPerPage(maxpp int64) *PaginationSettings {
	ps.MaxPerPage = maxpp
	return ps
}

func (ps *PaginationSettings) SetDefaultPerPage(defaultpp int64) *PaginationSettings {
	ps.DefaultPerPage = defaultpp
	return ps
}

func (ps *PaginationSettings) SetShowAllKey(key string) *PaginationSettings {
	ps.ShowAllKey = key
	return ps
}

func (ps *PaginationSettings) SetPagingKey(key string) *PaginationSettings {
	ps.PageKey = key
	return ps
}

func (ps *PaginationSettings) SetOffSetKey(key string) *PaginationSettings {
	ps.OffsetKey = key
	return ps
}

func (ps *PaginationSettings) GetPaginator(QueryParams url.Values) *options.FindOptions {
	pOptions := options.Find()
	showAll, err := strconv.ParseBool(QueryParams.Get(ps.ShowAllKey))
	if !showAll || err != nil {

		ppValue, err := strconv.ParseInt(QueryParams.Get(ps.PageKey), 10, 64)
		if err != nil || ppValue <= 0 {
			ppValue = ps.DefaultPerPage
		}
		if ppValue > 50 {
			ppValue = ps.MaxPerPage
		}
		offsetValue, err := strconv.ParseInt(QueryParams.Get(ps.OffsetKey), 10, 64)
		if err != nil || offsetValue < 0 {
			offsetValue = 0
		}
		pOptions.SetSkip(offsetValue)
		pOptions.SetLimit(ppValue)
	}
	return pOptions
}
