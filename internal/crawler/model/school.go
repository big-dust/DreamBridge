package model

type SchoolListResponse struct {
	Code      string         `json:"code"`      // 响应码
	Message   string         `json:"message"`   // 响应消息
	Data      SchoolListData `json:"data"`      // 数据部分
	Location  string         `json:"location"`  // 位置信息
	Encrydata string         `json:"encrydata"` // 加密数据
}
type Item struct {
	// useful
	SchoolID      int    `json:"school_id"`       // 学校ID
	Name          string `json:"name"`            // 名称
	CodeEnroll    string `json:"code_enroll"`     // 招生代码
	CityName      string `json:"city_name"`       // 城市名称
	DualClassName string `json:"dual_class_name"` // 双一流分类名称
	F211          int    `json:"f211"`            // 是否211工程院校
	F985          int    `json:"f985"`            // 是否985工程院校
	Level         string `json:"level"`           // 等级
	// unuseful
	Admissions    string `json:"admissions"`     // 录取信息
	Answerurl     string `json:"answerurl"`      // 回答URL
	Belong        string `json:"belong"`         // 所属分类
	Central       string `json:"central"`        // 是否为中央直属
	CityID        string `json:"city_id"`        // 城市ID
	CollegesLevel string `json:"colleges_level"` // 学院等级
	CountyID      string `json:"county_id"`      // 县级ID
	CountyName    string `json:"county_name"`    // 县级名称
	Department    string `json:"department"`     // 部门/系
	Doublehigh    string `json:"doublehigh"`     // 是否双高学校（高水平大学建设高校）
	DualClass     string `json:"dual_class"`     // 双一流分类
	Hightitle     string `json:"hightitle"`      // 高层次称号
	InnerRate     int    `json:"inner_rate"`     // 内部评比率
	IsRecruitment string `json:"is_recruitment"` // 是否招生
	LevelName     string `json:"level_name"`     // 等级名称
	Nature        string `json:"nature"`         // 性质
	NatureName    string `json:"nature_name"`    // 性质名称
	OuterRate     int    `json:"outer_rate"`     // 外部评比率
	ProvinceID    string `json:"province_id"`    // 省份ID
	ProvinceName  string `json:"province_name"`  // 省份名称
	Rank          string `json:"rank"`           // 排名
	RankType      string `json:"rank_type"`      // 排名类型
	Rate          int    `json:"rate"`           // 评比率

	SchoolType      string `json:"school_type"`       // 学校类型
	TagName         string `json:"tag_name"`          // 标签名称
	Type            string `json:"type"`              // 类型
	TypeName        string `json:"type_name"`         // 类型名称
	ViewMonth       string `json:"view_month"`        // 月浏览量
	ViewTotal       string `json:"view_total"`        // 总浏览量
	ViewTotalNumber string `json:"view_total_number"` // 总浏览数
	ViewWeek        string `json:"view_week"`         // 周浏览量
}
type SchoolListData struct {
	Item     []Item `json:"item"`     // 学校列表项
	NumFound int    `json:"numFound"` // 找到的数量
}

type SchoolInfoResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    SchoolInfoData `json:"data"`
	Md5     string         `json:"md5"`
}
type SchoolInfoData struct {
	SchoolID          string        `json:"school_id"`
	DataCode          string        `json:"data_code"`
	Name              string        `json:"name"`
	Type              string        `json:"type"`
	SchoolType        string        `json:"school_type"`
	SchoolNature      string        `json:"school_nature"`
	Level             string        `json:"level"`
	CodeEnroll        string        `json:"code_enroll"`
	ZsCode            string        `json:"zs_code"`
	Belong            string        `json:"belong"`
	F985              string        `json:"f985"`
	F211              string        `json:"f211"`
	Department        string        `json:"department"`
	Admissions        string        `json:"admissions"`
	Central           string        `json:"central"`
	DualClass         string        `json:"dual_class"`
	IsSeal            string        `json:"is_seal"`
	AppliedGrade      string        `json:"applied_grade"`
	NumSubject        string        `json:"num_subject"`
	NumMaster         string        `json:"num_master"`
	NumDoctor         string        `json:"num_doctor"`
	NumAcademician    string        `json:"num_academician"`
	NumLibrary        string        `json:"num_library"`
	NumLab            string        `json:"num_lab"`
	ProvinceID        string        `json:"province_id"`
	CityID            string        `json:"city_id"`
	CountyID          string        `json:"county_id"`
	IsAds             string        `json:"is_ads"`
	IsRecruitment     string        `json:"is_recruitment"`
	CreateDate        string        `json:"create_date"`
	Area              int           `json:"area"`
	OldName           string        `json:"old_name"`
	IsFenxiao         string        `json:"is_fenxiao"`
	Status            string        `json:"status"`
	AddID             string        `json:"add_id"`
	AddTime           string        `json:"add_time"`
	UpdateID          string        `json:"update_id"`
	UpdateTime        string        `json:"update_time"`
	AdLevel           string        `json:"ad_level"`
	Short             string        `json:"short"`
	EPc               string        `json:"e_pc"`
	EApp              string        `json:"e_app"`
	RuankeRank        string        `json:"ruanke_rank"`
	Single            string        `json:"single"`
	CollegesLevel     string        `json:"colleges_level"`
	Doublehigh        string        `json:"doublehigh"`
	WslRank           string        `json:"wsl_rank"`
	QsRank            string        `json:"qs_rank"`
	XyhRank           string        `json:"xyh_rank"`
	IsSell            string        `json:"is_sell"`
	EolRank           string        `json:"eol_rank"`
	SchoolBatch       string        `json:"school_batch"`
	UsRank            string        `json:"us_rank"`
	IsLogo            string        `json:"is_logo"`
	NumMaster2        string        `json:"num_master2"`
	NumDoctor2        string        `json:"num_doctor2"`
	AiStatus          string        `json:"ai_status"`
	IsAds2            string        `json:"is_ads2"`
	CoopMoney         string        `json:"coop_money"`
	BdoldName         string        `json:"bdold_name"`
	GbhNum            string        `json:"gbh_num"`
	LevelName         string        `json:"level_name"`
	TypeName          string        `json:"type_name"`
	SchoolTypeName    string        `json:"school_type_name"`
	SchoolNatureName  string        `json:"school_nature_name"`
	DualClassName     string        `json:"dual_class_name"`
	SingleYear        int           `json:"single_year"`
	Remark            []interface{} `json:"remark"`
	ProvinceName      string        `json:"province_name"`
	CityName          string        `json:"city_name"`
	TownName          string        `json:"town_name"`
	Weiwangzhan       string        `json:"weiwangzhan"`
	Yjszs             string        `json:"yjszs"`
	Xiaoyuan          string        `json:"xiaoyuan"`
	Email             string        `json:"email"`
	SchoolEmail       string        `json:"school_email"`
	Address           string        `json:"address"`
	Postcode          string        `json:"postcode"`
	Site              string        `json:"site"`
	SchoolSite        string        `json:"school_site"`
	Phone             string        `json:"phone"`
	SchoolPhone       string        `json:"school_phone"`
	Miniprogram       string        `json:"miniprogram"`
	Content           string        `json:"content"`
	Video             Video         `json:"video"`
	VideoPc           VideoPc       `json:"video_pc"`
	IsVideo           int           `json:"is_video"`
	Dualclass         []Dualclass   `json:"dualclass"`
	SchoolSpecialNum  int           `json:"school_special_num"`
	Special           []Special     `json:"special"`
	NatureName        string        `json:"nature_name"`
	ProvinceScoreYear string        `json:"province_score_year"`
	QsWorld           string        `json:"qs_world"`
	Rank              Rank          `json:"rank"`
	Fenxiao           []Fenxiao     `json:"fenxiao"`
	GbhURL            string        `json:"gbh_url"`
	IsYikao           int           `json:"is_yikao"`
	YkFeature         []string      `json:"yk_feature"`
	YkType            []string      `json:"yk_type"`
}
type Video struct {
	SchoolID string `json:"school_id"`
	URL      string `json:"url"`
	URLType  string `json:"url_type"`
	ImgURL   string `json:"img_url"`
}
type VideoPc struct {
	SchoolID string `json:"school_id"`
	URL      string `json:"url"`
	URLType  string `json:"url_type"`
	ImgURL   string `json:"img_url"`
}
type Dualclass struct {
	ID       string `json:"id"`
	SchoolID string `json:"school_id"`
	Class    string `json:"class"`
}
type Special struct {
	ID               string `json:"id"`
	SchoolID         string `json:"school_id"`
	SpecialID        string `json:"special_id"`
	NationFeature    string `json:"nation_feature"`
	ProvinceFeature  string `json:"province_feature"`
	IsImportant      string `json:"is_important"`
	LimitYear        string `json:"limit_year"`
	Year             string `json:"year"`
	Level3Weight     string `json:"level3_weight"`
	NationFirstClass string `json:"nation_first_class"`
	XuekeRankScore   string `json:"xueke_rank_score"`
	IsVideo          int    `json:"is_video"`
	SpecialName      string `json:"special_name"`
	LevelName        string `json:"level_name"`
}
type Rank struct {
	RuankeRank string `json:"ruanke_rank"`
	XyhRank    string `json:"xyh_rank"`
	QsWorld    string `json:"qs_world"`
	UsRank     string `json:"us_rank"`
	RkMb       string `json:"rk_mb"`
	RkDl       string `json:"rk_dl"`
	RkZw       string `json:"rk_zw"`
	RkYy       string `json:"rk_yy"`
	RkCj       string `json:"rk_cj"`
	RkYuy      string `json:"rk_yuy"`
	RkZf       string `json:"rk_zf"`
	RkTy       string `json:"rk_ty"`
	RkMz       string `json:"rk_mz"`
	XyhGz1     string `json:"xyh_gz_1"`
	XyhGz2     string `json:"xyh_gz_2"`
	XyhGz3     string `json:"xyh_gz_3"`
	XyhMb1     string `json:"xyh_mb_1"`
	XyhMb2     string `json:"xyh_mb_2"`
	XyhMb3     string `json:"xyh_mb_3"`
	TwsChina   string `json:"tws_china"`
}
type Yuanxi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type Fenxiao struct {
	FxName string   `json:"fx_name"`
	Yuanxi []Yuanxi `json:"yuanxi"`
}

type JobDetailResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Jobrate struct {
			Job struct {
				One string `json:"1"`
			} `json:"job"`
			Postgraduate struct {
				One string `json:"1"`
			} `json:"postgraduate"`
			Abroad struct {
				One string `json:"1"`
			} `json:"abroad"`
		} `json:"jobrate"`
		Province map[string]struct {
			ID           string `json:"id"`
			SchoolID     string `json:"school_id"`
			Province     string `json:"province"`
			Rate         string `json:"rate"`
			Num          string `json:"num"`
			Sort         string `json:"sort"`
			UpdateID     string `json:"update_id"`
			UpdateTime   string `json:"update_time"`
			Year         string `json:"year"`
			ProvinceName string `json:"province_name"`
		} `json:"province"`
		Attr struct {
			Minyingqiye     string `json:"民营企业"`
			Guoyouqiye      string `json:"国有企业"`
			Dangzhengjiguan string `json:"党政机关"`
			Keyandwunit     string `json:"科研单位"`
			Qitashiyedanwei string `json:"其他事业单位"`
			Waiziquanye     string `json:"外资企业"`
		} `json:"attr"`
		Company map[string]string `json:"company"`
		Gradute []struct {
			ID         string `json:"id"`
			SchoolID   string `json:"school_id"`
			FemaleNum  string `json:"female_num"`
			MenNum     string `json:"men_num"`
			MenRate    string `json:"men_rate"`
			FemaleRate string `json:"female_rate"`
			UpdateID   string `json:"update_id"`
			UpdateTime string `json:"update_time"`
			Year       string `json:"year"`
		} `json:"gradute"`
		Remark string `json:"remark"`
	} `json:"data"`
	Md5 string `json:"md5"`
}

type SpecialResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    SpecialData `json:"data"`
	Md5     string      `json:"md5"`
}
type SpecialData struct {
	NationFeature []NationFeature         `json:"nation_feature"`
	SpecialDetail map[string][]DetailInfo `json:"special_detail"`
}

type NationFeature struct {
	ID               string `json:"id"`
	SchoolID         string `json:"school_id"`
	SpecialID        string `json:"special_id"`
	NationFeature    string `json:"nation_feature"`
	ProvinceFeature  string `json:"province_feature"`
	IsImportant      string `json:"is_important"`
	LimitYear        string `json:"limit_year"`
	Year             string `json:"year"`
	Level3Weight     string `json:"level3_weight"`
	NationFirstClass string `json:"nation_first_class"`
	XuekeRankScore   string `json:"xueke_rank_score"`
	IsVideo          int    `json:"is_video"`
	SpecialName      string `json:"special_name"`
	SpecialType      string `json:"special_type"`
	TypeName         string `json:"type_name"`
	Level3Name       string `json:"level3_name"`
	Level3ID         string `json:"level3_id"`
	Level3Code       string `json:"level3_code"`
	Level2Name       string `json:"level2_name"`
	Level2ID         string `json:"level2_id"`
	Level2Code       string `json:"level2_code"`
	Code             string `json:"code"`
}

type DetailInfo struct {
	ID               string `json:"id"`
	SchoolID         string `json:"school_id"`
	SpecialID        string `json:"special_id"`
	NationFeature    string `json:"nation_feature"`
	ProvinceFeature  string `json:"province_feature"`
	IsImportant      string `json:"is_important"`
	LimitYear        string `json:"limit_year"`
	Year             string `json:"year"`
	Level3Weight     string `json:"level3_weight"`
	NationFirstClass string `json:"nation_first_class"`
	XuekeRankScore   string `json:"xueke_rank_score"`
	IsVideo          int    `json:"is_video"`
	SpecialName      string `json:"special_name"`
	SpecialType      string `json:"special_type"`
	TypeName         string `json:"type_name"`
	Level3Name       string `json:"level3_name"`
	Level3Code       string `json:"level3_code"`
	Level2Name       string `json:"level2_name"`
	Level2ID         string `json:"level2_id"`
	Level2Code       string `json:"level2_code"`
	Code             string `json:"code"`
}

type HistoryRecruitResponse struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    HistoryRecruitData `json:"data"`
	Md5     string             `json:"md5"`
	Time    string             `json:"time"`
}

// 2-7文科，1-7理科，2074-14历史类，2073-14物理类
type HistoryRecruitData map[string]map[string][]RecruitInfo

type RecruitInfo struct {
	Year         string `json:"year"`
	Type         string `json:"type"`
	Batch        string `json:"batch"`
	Name         string `json:"name"`
	Remark       string `json:"remark"`
	ElectiveInfo string `json:"elective_info"`
	Num          string `json:"num"`
}

type HistoryAdmissionResponse struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Data    HistoryRecruitData `json:"data"`
	Md5     string             `json:"md5"`
	Time    string             `json:"time"`
}

// 2-7文科，1-7理科，2074-14历史类，2073-14物理类
type HistoryAdmissionData map[string]map[string][]AdmissionInfo

type AdmissionInfo struct {
	Year         string `json:"year"`
	Type         string `json:"type"`
	Batch        string `json:"batch"`
	Name         string `json:"name"`
	Remark       string `json:"remark"`
	ElectiveInfo string `json:"elective_info"`
	Min          string `json:"min"`
	MinSection   string `json:"min_section"`
}
