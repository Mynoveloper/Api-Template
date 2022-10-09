package model

type Menu struct {
	UUID      string `bson:"-"`
	Separator string `bson:"Separator"`
	Label     string `bson:"Label"`
	Style     string `bson:"Style"`
	Class     string `bson:"Class"`
	PathTo    string `bson:"PathTo"`
	Url       string `bson:"Url"`
	Disabled  bool   `bson:"Disabled"`
	Target    any    `bson:"Target"`
	Icon      string `bson:"Icon"`
	Badge     string `bson:"Badge"`
	Items     []Menu `bson:"Items, omitempty"`
}

type Menus []Menu
