package config

const (
	//service port
	//ItemSaverPort = 1234
	//WorkerPort0   = 9000

	//Elastic Search
	ElasticIndex = "dating_profile"

	//RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.SaveItem"
	CrawlServiceRpc = "CrawlService.Process"

	//Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	//Rate limiting
	Qps = 20
)
