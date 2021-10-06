package charts

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ChartsApi
	//SystemApiApi
	//AuthorityApi
	//AutoCodeApi
	//BaseApi
	//CasbinApi
	//DictionaryApi
	//DictionaryDetailApi
	//SystemApi
	//DBApi
	//JwtApi
	//OperationRecordApi
	//AuthorityMenuApi
}

var chartsService = service.ServiceGroupApp.ChartsServiceGroup
