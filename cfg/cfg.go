package cfg

type Config struct {
	debug  bool
	server CfgServer
	mysql  CfgMysql
}

type CfgServer struct {
	port string
}

type CfgMysql struct {
	url  string
	port string
}
