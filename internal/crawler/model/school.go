package model

type SchoolListResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data      Data   `json:"data"`
	Location  string `json:"location"`
	Encrydata string `json:"encrydata"`
}
type Item struct {
	Admissions      string `json:"admissions"`
	Answerurl       string `json:"answerurl"`
	Belong          string `json:"belong"`
	Central         string `json:"central"`
	CityID          string `json:"city_id"`
	CityName        string `json:"city_name"`
	CodeEnroll      string `json:"code_enroll"`
	CollegesLevel   string `json:"colleges_level"`
	CountyID        string `json:"county_id"`
	CountyName      string `json:"county_name"`
	Department      string `json:"department"`
	Doublehigh      string `json:"doublehigh"`
	DualClass       string `json:"dual_class"`
	DualClassName   string `json:"dual_class_name"`
	F211            int    `json:"f211"`
	F985            int    `json:"f985"`
	Hightitle       string `json:"hightitle"`
	InnerRate       int    `json:"inner_rate"`
	IsRecruitment   string `json:"is_recruitment"`
	Level           string `json:"level"`
	LevelName       string `json:"level_name"`
	Name            string `json:"name"`
	Nature          string `json:"nature"`
	NatureName      string `json:"nature_name"`
	OuterRate       int    `json:"outer_rate"`
	ProvinceID      string `json:"province_id"`
	ProvinceName    string `json:"province_name"`
	Rank            string `json:"rank"`
	RankType        string `json:"rank_type"`
	Rate            int    `json:"rate"`
	SchoolID        int    `json:"school_id"`
	SchoolType      string `json:"school_type"`
	TagName         string `json:"tag_name"`
	Type            string `json:"type"`
	TypeName        string `json:"type_name"`
	ViewMonth       string `json:"view_month"`
	ViewTotal       string `json:"view_total"`
	ViewTotalNumber string `json:"view_total_number"`
	ViewWeek        string `json:"view_week"`
}
type Data struct {
	Item     []Item `json:"item"`
	NumFound int    `json:"numFound"`
}
