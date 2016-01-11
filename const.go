package main

const (
	RailwayGinza      = "odpt.Railway:TokyoMetro.Ginza"
	RailwayMarunouchi = "odpt.Railway:TokyoMetro.Marunouchi"
	RailwayHibiya     = "odpt.Railway:TokyoMetro.Hibiya"
	RailwayTozai      = "odpt.Railway:TokyoMetro.Tozai"
	RailwayChiyoda    = "odpt.Railway:TokyoMetro.Chiyoda"
	RailwayYurakucho  = "odpt.Railway:TokyoMetro.Yurakucho"
	RailwayHanzomon   = "odpt.Railway:TokyoMetro.Hanzomon"
	RailwayNamboku    = "odpt.Railway:TokyoMetro.Namboku"
	RailwayFukutoshin = "odpt.Railway:TokyoMetro.Fukutoshin"
)

var RailwayToName map[string]string = map[string]string{
	RailwayGinza:      "銀座線",
	RailwayMarunouchi: "丸ノ内線",
	RailwayHibiya:     "日比谷線",
	RailwayTozai:      "東西線",
	RailwayChiyoda:    "千代田線",
	RailwayYurakucho:  "有楽町線",
	RailwayHanzomon:   "半蔵門線",
	RailwayNamboku:    "南北線",
	RailwayFukutoshin: "副都心線",
}

const (
	TypeTrain            = "odpt:Train"
	TypeTrainInfomation  = "odpt:TrainInformation"
	TypeStationTimetable = "odpt:StationTimetable"
	TypeStationFacility  = "odpt:StationFacility"
	TypePassengerSurvey  = "odpt:PassengerSurvey"
	TypeRailwayFare      = "odpt:RailwayFare"
	TypePoi              = "ug:Poi"
	TypeMLITStation      = "mlit:Station"
	TypeMLITRailway      = "mlit:Railway"
	TypeStation          = "odpt:Station"
	TypeRailway          = "odpt:Railway"
)

const (
	StatusSuspended         = "運転見合わせ"
	StatusShuttle            = "折返し運転"
	StatusTimetableDisarray = "ダイヤ乱れ"
	// 運転再開・ダイヤ乱れ
	StatusDelay          = "遅延"
	StatusPartiallyDelay = "一部列車遅延"
	StatusAbortDirect    = "直通運転中止"
	StatusRestartDirect  = "直通運転再開"
	// 快速運転中止
	// 快速運転再開
	// 準急運転中止
	// 準急運転再開
	// 急行運転中止
	// 急行運転再開
	// 通勤急行運転中止
	// 通勤急行運転再開
	// 女性専用車両中止
	// 運転再開見込
	// 運転再開
	// 運転規制
	// 速度規制
	StatusCancel = "運休"
	// メトロさがみ運休
	// メトロさがみ７０号運休
	// メトロさがみ８０号運休
	// メトロホームウエイ運休
	// メトロホームウエイ４１号運休
	// メトロホームウエイ４３号運休
	// メトロホームウエイ７１号運休
	// メトロはこね運休
	// メトロはこね２１号運休
	// メトロはこね２２号運休
	// メトロはこね２３号運休
	// メトロはこね２４号運休
	// 臨時特急ロマンスカー運休
	// 臨時列車運休
	// 振替輸送
	// バス代行
	// リフレッシュ工事
	// お知らせ
	StatusUnknown = "__UNKNOWN__"
)

const (
	LineColorGinza      = "#FF9500"
	LineColorMarunouchi = "#F62E36"
	LineColorHibiya     = "#B5B5AC"
	LineColorTozai      = "#009BBF"
	LineColorChiyoda    = "#00BB85"
	LineColorYurakucho  = "#C1A470"
	LineColorHanzomon   = "#8F76D6"
	LineColorNamnoku    = "#00AC9B"
	LineColorFukutoshin = "#9C5E31"
)
