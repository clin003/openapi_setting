package app_conf

type OpenAPIConfModel struct {
	// # openAPI服务绑定端口
	Addr string `yaml:"addr"`
	// // # openAPI Server的名字
	// Name string `yaml:"name"`
	// # openAPI服务器的ip:port
	Url string `yaml:"url"`
	// # 自检服务重试的次数
	MaxPingCount int `yaml:"max_ping_count"`
	// #TLS服务域名(可选配置)
	AutotlsDomain string `yaml:"autotls_domain"`
	// #打开Tls配置，true开启TLS服务，false会关闭TLS服务。 (可选配置)
	AutotlsEnable bool   `yaml:"autotls_enable"`
	TlsKeyFile    string `yaml:"tls_key_file"`
	TlsPemFile    string `yaml:"tls_pem_file"`
	// openapi的日志相关配置。
	Log Log `yaml:"log"`
	// openapi的数据库相关配置说明
	Db Db `yaml:"db"`
	// 同步到七牛云空间开关
	SyncEnable bool  `yaml:"sync_enable"`
	Qiniu      Qiniu `yaml:"qiniu"`

	// #检测文案图片是否包含二维码
	CheckImageQrcode bool `yaml:"check_image_qrcode"`
	// # 允许客户端上传文件到openAPI服务器，
	UploadFile bool `yaml:"upload_file"`
	// # 启动WSSERVER服务，慧转发广播模式
	WsserverEnable bool `yaml:"wsserver_enable"`
	// ## 网页图文直播服务配置
	LiveEnbale        bool   `yaml:"live_enbale"`
	LiveTitle         string `yaml:"live_title"`
	LiveName          string `yaml:"live_name"`
	LiveWsserverUrl   string `yaml:"live_wsserver_url"`
	LiveWsserverToken string `yaml:"live_wsserver_token"`
	LiveMsgMaxCount   int    `yaml:"live_msg_max_count"`
	LiveTabLink0      string `yaml:"live_tab_link0"`
	LiveTabLink0Name  string `yaml:"live_tab_link0_name"`
	LiveTabLink1      string `yaml:"live_tab_link1"`
	LiveTabLink1Name  string `yaml:"live_tab_link1_name"`
	LiveTabLink2      string `yaml:"live_tab_link2"`
	LiveTabLink2Name  string `yaml:"live_tab_link2_name"`
	// ## 转链服务配置
	// # 转链服务开关
	TransformlinkEnbale bool `yaml:"transformlink_enbale"`
	// # 转链容错开关
	TransformlinkFailoverEnable bool `yaml:"transformlink_failover_enable"`
	// # 网页转链服务接口域名
	TransformlinkApiserverUrl string `yaml:"transformlink_apiserver_url"`
	// # 网页转链Title
	TransformlinkTitle string `yaml:"transformlink_title"`
	TransformlinkName  string `yaml:"transformlink_name"`
	// ## To慧广播安全验证token (可选配置)
	SenderWsserverToken string `yaml:"sender_wsserver_token"`
	// # # openAPI Server true 取消同步授权信息
	OpenapiServerEnable bool `yaml:"openapi_server_enable"`
	// #验证客户端Token和api_client端的hlty_token是否一致
	CheckToken bool `yaml:"check_token"`
	// # (同步日志、上传文件、慧广播消息、APP同步消息、淘宝转链备用接口)验证Token (可选配置)
	HltyToken string `yaml:"hlty_token"`
	// # 同步授权Token
	HltyUuid string `yaml:"hlty_uuid"`

	// # webhook配置
	WeixinWorkWebhookUrl      string `yaml:"weixin_work_webhook_url"`
	WeixinWorkWebhookEnable   bool   `yaml:"weixin_work_webhook_enable"`
	DingtalkWebhookUrl        string `yaml:"dingtalk_webhook_url"`
	DingtalkWebhookSecret     string `yaml:"dingtalk_webhook_secret"`
	DingtalkWebhookEnable     bool   `yaml:"dingtalk_webhook_enable"`
	DingtalkWebhookRichEnable bool   `yaml:"dingtalk_webhook_rich_enable"`

	Sdk Sdk `yaml:"sdk"`

	// # 微信公众号sdk的重试次数(可选配置)
	MaxRetryTimes int `yaml:"max_retry_times"`
	// # 微信公众号sdk刷新token间隔
	MaxRefreshTick int `yaml:"max_refresh_tick"`
	// # 微信公众号被关注自动回复消息
	WechatWelcomeMsg string `yaml:"wechat_welcome_msg"`
	// // # 微信公众号关键词自动回复(可选配置)
	// WechatKeyMsg WechatKeyMsg `yaml:"wechat_key_msg"`

	// # 打开淘宝联盟单品券高效转链接口配置，
	TbkprivilegeEnable bool `yaml:"tbkprivilege_enable"`
	// # 关闭淘宝联盟口令及链接识别转链,禁止识别淘宝一键转链
	TbkprivilegeDisable bool `yaml:"tbkprivilege_disable"`
	// # 自定义淘口令符号
	// # 淘口令自定义符号左侧部分
	TpwdSymbolLeft string `yaml:"tpwd_symbol_left"`
	// # 淘口令自定义符号右侧部分
	TpwdSymbolRight string `yaml:"tpwd_symbol_right"`
	// # 文案级关键词替换，格式 key>>key2|key3 (可选配置)
	ContentKeywordReplaceAll string `yaml:"content_keyword_replace_all"`
	// # 文案级关键词过滤，格式 key|key2 (可选配置)
	ContentKeywordFilter string `yaml:"content_keyword_filter"`
	// # 文案中淘礼金口令处理模式，可供选择有三个选项：0、转链，1、屏蔽，2、原链(可选配置)
	ContentTaolijinModel int `yaml:"content_taolijin_model"`
	// # 文案中特定关键词列表，格式 key|key2 (可选配置)
	ContentKeywordExFilter string `yaml:"content_keyword_ex_filter"`
	// # 文案中特定关键词处理模式(转链和原链模式都会经历文案级的关键词替换和屏蔽处理逻辑)，可供选择有三个选项：0、转链，1、屏蔽，2、原链 (可选配置)
	ContentKeywordExModel int `yaml:"content_keyword_ex_model"`
	// 客户端配置项目
	App App `yaml:"app"`
}

type Sdk struct {
	Taobao SdkTaobao `yaml:"taobao"`
	// Suning    SdkSuning   `yaml:"suning"`
	Kaola     SdkAppInfo  `yaml:"kaola"`
	Pinduoduo SdkAppInfo  `yaml:"pinduoduo"`
	Jingdong  SdkJingdong `yaml:"jingdong"`
	Wechat    SdkWechat   `yaml:"wechat"`
}
type SdkAppInfo struct {
	Appkey    string `yaml:"appkey"`
	Appsecret string `yaml:"appsecret"`
	// Adzoneid  string `yaml:"adzoneid"`
	// Siteid    string `yaml:"siteid"`
	Pid string `yaml:"pid"`
}

type SdkTaobao struct {
	Appkey    string `yaml:"appkey"`
	Appsecret string `yaml:"appsecret"`
	Adzoneid  string `yaml:"adzoneid"`
	Pid       string `yaml:"pid"`
}

type SdkSuning struct {
	Appkey    string `yaml:"appkey"`
	Appsecret string `yaml:"appsecret"`
	Promoter  string `yaml:"promoter"`
	Teamcpa   string `yaml:"teamcpa"`
	Adzoneid  string `yaml:"adzoneid"`
	Hongbao   string `yaml:"hongbao"`
}

// type SdkPinduoduo struct {
// 	SdkAppInfo
// 	// Appkey    string `yaml:"appkey"`
// 	// Appsecret string `yaml:"appsecret"`
// }

// type SdkKaola struct {
// 	SdkAppInfo
// 	// Appkey    string `yaml:"appkey"`
// 	// Appsecret string `yaml:"appsecret"`
// }
type SdkJingdong struct {
	Appkey    string `yaml:"appkey"`
	Appsecret string `yaml:"appsecret"`
	Adzoneid  string `yaml:"adzoneid"`
	Siteid    string `yaml:"siteid"`
	UnionId   string `yaml:"union_id"`
}
type SdkWechat struct {
	Appkey    string `yaml:"appkey"`
	Appsecret string `yaml:"appsecret"`
	Token     string `yaml:"token"`
	Aeskey    string `yaml:"aeskey"`
}

type Qiniu struct {
	DeleteAfterDays int    `yaml:"delete_after_days"`
	Ak              string `yaml:"ak"`
	Sk              string `yaml:"sk"`
	Bucket          string `yaml:"bucket"`
	Domain          string `yaml:"domain"`
	IsOnly          bool   `yaml:"is_only"`
}

// type WechatKeyMsg struct {
// 	key1 string `yaml:"关键词"`
// 	key2 string `yaml:"关键词2"`
// }

type Log struct {
	RollingPolicy  string `yaml:"rollingPolicy"`
	LogRotateDate  int    `yaml:"log_rotate_date"`
	LogRotateSize  int    `yaml:"log_rotate_size"`
	LogBackupCount int    `yaml:"log_backup_count"`
	Writers        string `yaml:"writers"`
	LoggerLevel    string `yaml:"logger_level"`
	LoggerFile     string `yaml:"logger_file"`
	LogFormatText  bool   `yaml:"log_format_text"`
}

type Db struct {
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbtype   string `yaml:"dbtype"`
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	Host     string `yaml:"host"`
}

type App struct {
	Tb    AppTb    `yaml:"tb"`
	Sn    AppSn    `yaml:"sn"`
	Kaola AppKaola `yaml:"kaola"`
	Pdd   AppPdd   `yaml:"pdd"`
	Jd    AppJd    `yaml:"jd"`
}
type AppTb struct {
	Pid string `yaml:"pid"`
}
type AppSn struct {
	Promoter string `yaml:"promoter"`
	Teamcpa  string `yaml:"teamcpa"`
	Adzoneid string `yaml:"adzoneid"`
	Hongbao  string `yaml:"hongbao"`
}
type AppKaola struct {
	Pid string `yaml:"pid"`
}
type AppPdd struct {
	Pid string `yaml:"pid"`
}
type AppJd struct {
	UnionId string `yaml:"union_id"`
}
