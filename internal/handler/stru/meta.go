package stru

type UserPortraitData struct {
	Id        int64  `form:"id" json:"id"`
	OtsCode   string `form:"otsCode" json:"otsCode" validate:"required"`
	RedisCode string `form:"redisCode" json:"redisCode" validate:"required"`
	TableName string `form:"tableName" json:"tableName" validate:"required"`
	Field     string `form:"field" json:"field" validate:"required"`
	Column    string `form:"column" json:"column" validate:"required"`
	Key       string `form:"key" json:"key" validate:"required"`
	Timeout   int32  `form:"timeout" json:"timeout" validate:"required"`
	Code      string `form:"code" json:"code"`
	Desc      string `form:"desc" json:"desc"`
	Status    int8   `form:"status" json:"status"`
}

type ItemPortraitData struct {
	Id         int64  `form:"id" json:"id"`
	OtsCode    string `form:"otsCode" json:"otsCode" validate:"required"`
	RedisCode  string `form:"redisCode" json:"redisCode" validate:"required"`
	TableName  string `form:"tableName" json:"tableName" validate:"required"`
	Field      string `form:"field" json:"field" validate:"required"`
	Column     string `form:"column" json:"column" validate:"required"`
	Prop       string `form:"prop" json:"prop"`
	Key        string `form:"key" json:"key" validate:"required"`
	ParentProp string `form:"parentProp" json:"parentProp"`
	Timeout    int32  `form:"timeout" json:"timeout" validate:"required"`
	//NextId           int64  `form:"nextId" json:"nextId"`
	//Conditional      string `form:"conditional" json:"conditional"`
	//ConditionalField string `form:"conditionalField" json:"conditionalField"`
	Code   string `form:"code" json:"code"`
	Desc   string `form:"desc" json:"desc"`
	Status int8   `form:"status" json:"status"`
}
type DictDomainData struct {
	Id             int64  `form:"id" json:"id"`
	Code           string `form:"code" json:"code" validate:"required"`
	Desc           string `form:"desc" json:"desc"`
	RmipId         int64  `form:"rmipId" json:"rmipId" validate:"required"`
	Domain         string `form:"domain" json:"domain" validate:"required"`
	Feature        string `form:"feature" json:"feature"`
	FlowProportion int32  `form:"flowProportion" json:"flowProportion"`
	RecallSize     int32  `form:"recallSize" json:"recallSize"`
	JumpType       int8   `form:"jumpType" json:"jumpType"`
	Status         int8   `form:"status" json:"status"`
}
type DictSceneData struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code" validate:"required"`
	Desc   string `form:"desc" json:"desc"`
	Option string `form:"option" json:"option" validate:"required"`
	Status int8   `form:"status" json:"status"`
}
type DictModeData struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code" validate:"required"`
	Desc   string `form:"desc" json:"desc"`
	Status int8   `form:"status" json:"status"`
}
type DictSortData struct {
	Id      int64  `form:"id" json:"id"`
	Code    string `form:"code" json:"code" validate:"required"`
	Desc    string `form:"desc" json:"desc"`
	OssCode string `form:"ossCode" json:"ossCode" validate:"required"`
	Algo    int8   `form:"algo" json:"algo" validate:"required"`
	Bucket  string `form:"bucket" json:"bucket" validate:"required"`
	Address string `form:"address" json:"address" validate:"required"`
	Status  int8   `form:"status" json:"status"`
}

type FilterListParams struct {
	PageNum           int32  `form:"pageNum" json:"pageNum" validate:"required"`
	PageSize          int32  `form:"pageSize" json:"pageSize" validate:"required"`
	Id                int64  `form:"id" json:"id"`
	FilterId          int64  `form:"filterId" json:"filterId"`
	RmdId             int64  `form:"rmdId" json:"rmdId"`
	Mode              int8   `form:"mode" json:"mode"`
	Status            int8   `form:"status" json:"status"`
	ReviewStatus      int8   `form:"reviewStatus" json:"reviewStatus"`
	PublishActionCode string `form:"publishActionCode" json:"publishActionCode"`
}

type ReviewContentListParams struct {
	PageNum            int32    `form:"pageNum" json:"pageNum"  validate:"required"`
	PageSize           int32    `form:"pageSize" json:"pageSize"  validate:"required"`
	Id                 int64    `form:"id" json:"id"`
	OperationId        int64    `form:"operationId" json:"operationId"`
	ProjectIds         []int64  `form:"projectIds[]" json:"projectIds[]"`
	PublishActionCodes []string `form:"publishActionCodes[]" json:"publishActionCodes[]"`
	ReviewStatuses     []int8   `form:"reviewStatuses[]" json:"reviewStatuses[]"`
}

type FilterSaveData struct {
	Id       int64  `form:"id" json:"id"`
	RmdId    int64  `form:"rmdId" json:"rmdId" validate:"required"`
	Mode     int8   `form:"mode" json:"mode" validate:"required"`
	Param    string `form:"param" json:"param"`
	Priority int32  `form:"priority" json:"priority"`
}
type FilterSaveParam struct {
	Action string          `form:"action" json:"action" validate:"required"`
	Data   *FilterSaveData `form:"data" json:"data" validate:"required,dive,required"`
}
type MetaDomainListParams struct {
	Id     int64 `form:"id" json:"id"`
	RaiId  int64 `form:"raiId" json:"raiId"`
	Status int8  `form:"status" json:"status"`
	Limit  int32 `form:"limit" json:"limit"`
}
type MetaBoostListParams struct {
	PageNum           int32  `form:"pageNum" json:"pageNum" validate:"required"`
	PageSize          int32  `form:"pageSize" json:"pageSize" validate:"required"`
	Id                int64  `form:"id" json:"id"`
	BoostId           int64  `form:"boostId" json:"boostId"`
	RmdId             int64  `form:"rmdId" json:"rmdId"`
	Status            int8   `form:"status" json:"status"`
	ReviewStatus      int8   `form:"reviewStatus" json:"reviewStatus"`
	PublishActionCode string `form:"publishActionCode" json:"publishActionCode"`
}
type MetaBoostSaveParam struct {
	Action string             `form:"action" json:"action" validate:"required"`
	Data   *MetaBoostSaveData `form:"data" json:"data" validate:"required,dive,required"`
}
type MetaBoostSaveData struct {
	Id       int64   `form:"id" json:"id"`
	RmdId    int64   `form:"rmdId" json:"rmdId" validate:"required"`
	Field    string  `form:"field" json:"field" validate:"required"`
	Pattern  string  `form:"pattern" json:"pattern" validate:"required"`
	Weight   float64 `form:"weight" json:"weight" validate:"required"`
	Priority int32   `form:"priority" json:"priority"`
}
type AbtestStreamListParams struct {
	Id     int64  `form:"id" json:"id"`
	Code   string `form:"code" json:"code"`
	Scene  int64  `form:"scene" json:"scene"`
	Mode   int64  `form:"mode" json:"mode"`
	Status int8   `form:"status" json:"status"`
	Limit  int32  `form:"limit" json:"limit"`
}
type AbtestStreamData struct {
	Id     int64  `form:"id" json:"id" validate:"required"`
	Code   string `form:"code" json:"code" validate:"required"`
	Desc   string `form:"desc" json:"desc"`
	Scene  int64  `form:"scene" json:"scene" validate:"required"`
	Mode   int64  `form:"mode" json:"mode" validate:"required"`
	Param  string `form:"param" json:"param"`
	Status int8   `form:"status" json:"status" validate:"required"`
}
type AbtestStreamMultiData struct {
	Id             int64             `form:"id" json:"id"`
	Code           string            `form:"code" json:"code" validate:"required"`
	Desc           string            `form:"desc" json:"desc"`
	Scene          int64             `form:"scene" json:"scene" validate:"required"`
	Mode           int64             `form:"mode" json:"mode" validate:"required"`
	Param          string            `form:"param" json:"param"`
	Status         int8              `form:"status" json:"status"`
	CreatedBy      string            `form:"createdBy" json:"createdBy"`
	UpdatedBy      string            `form:"updatedBy" json:"updatedBy"`
	AbTestInfoList []*AbTestInfoData `form:"abTestInfoList" json:"abTestInfoList"`
}

type AbTestInfoData struct {
	Id             int64                  `form:"id" json:"id"`
	Code           string                 `form:"code" json:"code"`
	Desc           string                 `form:"desc" json:"desc"`
	RrCode         string                 `form:"rrCode" json:"rrCode"`
	RasId          int64                  `form:"rasId" json:"rasId"`
	RmupId         int64                  `form:"rmupId" json:"rmupId"`
	StartTime      string                 `form:"startTime" json:"startTime"`
	EndTime        string                 `form:"endTime" json:"endTime"`
	Status         int8                   `form:"status" json:"status"`
	CreatedBy      string                 `form:"createdBy" json:"createdBy"`
	UpdatedBy      string                 `form:"updatedBy" json:"updatedBy"`
	MetaDomainList []*MetaDomainMutliData `form:"metaDomainList" json:"metaDomainList"`
}
type MetaDomainMutliData struct {
	Id               int64                 `form:"id" json:"id"`
	RaiId            int64                 `form:"raiId" json:"raiId"`
	DomainCode       string                `form:"domainCode" json:"domainCode"`
	Percentage       int32                 `form:"percentage" json:"percentage"`
	Status           int8                  `form:"status" json:"status"`
	CreatedBy        string                `form:"createdBy" json:"createdBy"`
	UpdatedBy        string                `form:"updatedBy" json:"updatedBy"`
	MetaBehaviorList []*MetaBehaviorData   `form:"metaBehaviorList" json:"metaBehaviorList"`
	MetaBoostList    []*MetaBoostMultiData `form:"metaBoostList" json:"metaBoostList"`
	MetaExposureList []*MetaExposureData   `form:"metaExposureList" json:"metaExposureList"`
	MetaFilterList   []*MetaFilterData     `form:"metaFilterList" json:"metaFilterList"`
	MetaFlowList     []*MetaFlowData       `form:"metaFlowList" json:"metaFlowList"`
	MetaRecallList   []*MetaRecallData     `form:"metaRecallList" json:"metaRecallList"`
	MetaScatterList  []*MetaScatterData    `form:"metaScatterList" json:"metaScatterList"`
	MetaSortList     []*MetaSortData       `form:"metaSortList" json:"metaSortList"`
}
type MetaBehaviorData struct {
	Id                     int64                     `form:"id" json:"id"`
	RmdId                  int64                     `form:"rmdId" json:"rmdId"`
	OtsCode                string                    `form:"otsCode" json:"otsCode"`
	TableName              string                    `form:"tableName" json:"tableName"`
	Interval               int64                     `form:"interval" json:"interval"`
	Size                   int64                     `form:"size" json:"size"`
	Field                  int64                     `form:"field" json:"field"`
	Rt                     int32                     `form:"rt" json:"rt"`
	Factor                 float64                   `form:"factor" json:"factor"`
	Weaken                 int32                     `form:"weaken" json:"weaken"`
	Status                 int8                      `form:"status" json:"status"`
	CreatedBy              string                    `form:"createdBy" json:"createdBy"`
	UpdatedBy              string                    `form:"updatedBy" json:"updatedBy"`
	MetaBehaviorRecallList []*MetaBehaviorRecallData `form:"metaBehaviorRecallList" json:"metaBehaviorRecallList"`
	MetaBehaviorScoreList  []*MetaBehaviorScoreData  `form:"metaBehaviorScoreList" json:"metaBehaviorScoreList"`
}
type MetaBehaviorRecallData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	Field     string `form:"field" json:"field"`
	Count     int64  `form:"count" json:"count"`
	IsSplit   int8   `form:"isSplit" json:"isSplit"`
	Separator string `form:"separator" json:"separator"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
type MetaBehaviorScoreData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	Bhv       string `form:"bhv" json:"bhv"`
	Weight    int32  `form:"weight" json:"weight"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
type MetaBoostMultiData struct {
	Id        int64   `form:"id" json:"id"`
	RmdId     int64   `form:"rmdId" json:"rmdId"`
	Field     string  `form:"field" json:"field"`
	Pattern   string  `form:"pattern" json:"pattern"`
	Weight    float64 `form:"weight" json:"weight"`
	Priority  int32   `form:"priority" json:"priority"`
	Status    int8    `form:"status" json:"status"`
	CreatedBy string  `form:"createdBy" json:"createdBy"`
	UpdatedBy string  `form:"updatedBy" json:"updatedBy"`
}
type MetaExposureData struct {
	Id                       int64                       `form:"id" json:"id"`
	RmdId                    int64                       `form:"rmdId" json:"rmdId"`
	StartTime                string                      `form:"startTime" json:"startTime"`
	EndTime                  string                      `form:"endTime" json:"endTime"`
	Count                    int64                       `form:"count" json:"count"`
	Records                  int64                       `form:"records" json:"records"`
	Offset                   int64                       `form:"offset" json:"offset"`
	RedisCode                string                      `form:"redisCode" json:"redisCode"`
	Key                      string                      `form:"key" json:"key"`
	Algo                     string                      `form:"algo" json:"algo"`
	Constitute               string                      `form:"constitute" json:"constitute"`
	Status                   int8                        `form:"status" json:"status"`
	CreatedBy                string                      `form:"createdBy" json:"createdBy"`
	UpdatedBy                string                      `form:"updatedBy" json:"updatedBy"`
	MetaExposureIntervalList []*MetaExposureIntervalData `form:"metaExposureIntervalList" json:"metaExposureIntervalList"`
}
type MetaExposureIntervalData struct {
	Id        int64  `form:"id" json:"id"`
	RmeId     int64  `form:"rmeId" json:"rmeId"`
	RrCode    string `form:"rrCode" json:"rrCode"`
	TimePoint string `form:"timePoint" json:"timePoint"`
	Count     int64  `form:"count" json:"count"`
	Priority  int32  `form:"priority" json:"priority"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
type MetaFilterData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	Mode      int8   `form:"mode" json:"mode"`
	Param     string `form:"param" json:"param"`
	Priority  int32  `form:"priority" json:"priority"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}

type MetaFlowData struct {
	Id             int64  `form:"id" json:"id"`
	Code           string `form:"code" json:"code"`
	Desc           string `form:"desc" json:"desc"`
	RmdId          int64  `form:"rmdId" json:"rmdId"`
	ActiveDays     int32  `form:"activeDays" json:"activeDays"`
	ContinuousDays int32  `form:"continuousDays" json:"continuousDays"`
	DayBehaviors   int32  `form:"dayBehaviors" json:"dayBehaviors"`
	Level          int32  `form:"level" json:"level"`
	RedisCode      string `form:"redisCode" json:"redisCode"`
	Key            string `form:"key" json:"key"`
	Algo           string `form:"algo" json:"algo"`
	Constitute     string `form:"constitute" json:"constitute"`
	Status         int8   `form:"status" json:"status"`
	CreatedBy      string `form:"createdBy" json:"createdBy"`
	UpdatedBy      string `form:"updatedBy" json:"updatedBy"`
}

type MetaRecallData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	OtsCode   string `form:"otsCode" json:"otsCode"`
	TableName string `form:"tableName" json:"tableName"`
	AlgoType  int8   `form:"algoType" json:"algoType"`
	Algo      string `form:"algo" json:"algo"`
	Field     string `form:"field" json:"field"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
type MetaScatterData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	Algo      int8   `form:"algo" json:"algo"`
	Field     string `form:"field" json:"field"`
	Model     int8   `form:"model" json:"model"`
	Params    string `form:"params" json:"params"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
type MetaSortData struct {
	Id        int64  `form:"id" json:"id"`
	RmdId     int64  `form:"rmdId" json:"rmdId"`
	RdsId     int64  `form:"rdsId" json:"rdsId"`
	Priority  int32  `form:"priority" json:"priority"`
	Status    int8   `form:"status" json:"status"`
	CreatedBy string `form:"createdBy" json:"createdBy"`
	UpdatedBy string `form:"updatedBy" json:"updatedBy"`
}
