package FilterSettings

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchOptions struct {
	Items []SearchItem
}

type SearchItem struct {
	Field      string
	SearchWord string
	Option     string
}

func NewSettings() *SearchOptions {
	return &SearchOptions{}
}

func (s *SearchOptions) AddField(FieldName string, SearchWord string, option string) *SearchOptions {
	s.Items = append(s.Items, SearchItem{
		Field:      FieldName,
		SearchWord: SearchWord,
		Option:     option,
	})
	return s
}

func (s *SearchOptions) GetSearchFilter() *bson.M {
	filter := bson.M{}
	fields := s.Items
	for _, v := range fields {
		filter[v.Field] = bson.M{
			"$regex": primitive.Regex{
				Pattern: v.SearchWord,
				Options: v.Option,
			},
		}
	}
	return &filter
}
