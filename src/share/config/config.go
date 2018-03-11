package config

const (
	ConfigPrefix = "class/"
	MysqlDSN     = "root:@(localhost:3306)/mtbsystem"
	Namespace    = "com.mtbsystem."
	LogPath      = "/data/goapp/log/"
	Num          = 20 // 分页每次取多少
	TickingNow   = 1  // 正在上映
	TickingWill  = 2  // 即将上映
	ActorType    = 1  // 演员
	DirectorType = 2  // 导演
)

const (
	ServiceNameUser    = "user"
	ServiceNameFilm    = "film"
	ServiceNameComment = "comment"
	ServiceNameCinema  = "cinema"
	ServiceNameOrder   = "order"
	ServiceNameCMS     = "cms"
)
