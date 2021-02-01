package setting

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-ini/ini"
)

type Config struct {
	Name     string `yaml:"name"`
	Registry struct {
		Consul struct {
			URL   string `yaml:"url"`
			Token string `yaml:"token"`
		}
	} `yaml:"registry"`
	Data struct {
		Dir string `yaml:"dir"`
	} `yaml:"data"`
}

type App struct {
	PageSize        int    `yaml:"PageSize"`
	PrefixUrl       string `yaml:"PrefixUrl"`
	RuntimeRootPath string `yaml:"RuntimeRootPath"`

	ImageSavePath  string   `yaml:"ImageSavePath"`
	ImageMaxSize   int      `yaml:"ImageMaxSize"`
	ImageAllowExts []string `yaml:"ImageAllowExts"`

	ExportSavePath string `yaml:"ExportSavePath"`
	QrCodeSavePath string `yaml:"QrCodeSavePath"`
	FontSavePath   string `yaml:"FontSavePath"`

	LogSavePath    string `yaml:"LogSavePath"`
	LogSaveName    string `yaml:"LogSaveName"`
	LogFileExt     string `yaml:"LogFileExt"`
	LogLevel       string `yaml:"LogLevel"`
	TimeFormat     string `yaml:"TimeFormat"`
	LogPushToRedis bool   `yaml:"LogPushToRedis"`

	DivisionPrecision int `yaml:"DivisionPrecision"`

	CdnUrl     string `yaml:"CdnUrl"`
	QBoxBucket string `yaml:"QBoxBucket"`
	QBoxAccess string `yaml:"QBoxAccess"`
	QBoxSecret string `yaml:"QBoxSecret"`

	//MinipStaffName                    string `yaml:"MinipStaffName"`
	//MinipStaffQrCodeURL               string `yaml:"MinipStaffQrCodeURL"`
	//MinipStaffAppID                   string `yaml:"MinipStaffAppID"`
	//MinipStaffSecret                  string `yaml:"MinipStaffSecret"`
	//MinipStaffOriID                   string `yaml:"MinipStaffOriID"`
	//MinipStaffToken                   string `yaml:"MinipStaffToken"`
	//MinipStaffEncodedAESKey           string `yaml:"MinipStaffEncodedAESKey"`
	//MinipStaffTemplateMsgManageNotice string `yaml:"MinipStaffTemplateMsgManageNotice"`
	//MinipStaffTemplateMsgOrderTimeout string `yaml:"MinipStaffTemplateMsgOrderTimeout"`
	//MinipStaffTemplateMsgOrderRemind  string `yaml:"MinipStaffTemplateMsgOrderRemind"`
	//MinipStaffTemplateMsgOrderFinish  string `yaml:"MinipStaffTemplateMsgOrderFinish"`
	//MinipStaffTemplateMsgOrderNew     string `yaml:"MinipStaffTemplateMsgOrderNew"`

	//MinipServiceProviderMchID  string `yaml:"MinipServiceProviderMchID"`
	//MinipServiceProviderAppID  string `yaml:"MinipServiceProviderAppID"`
	//MinipServiceProviderApiKey string `yaml:"MinipServiceProviderApiKey"`

	NoticeList string `yaml:"NoticeList"`

	DispatchOrderMaxRemindTimes int `yaml:"DispatchOrderMaxRemindTimes"`
}

var AppSetting = &App{}

type Remote struct {
	PmsBaseUrl        string `yaml:"PmsBaseUrl"`
	MinipUserBaseUrl  string `yaml:"MinipUserBaseUrl"`
	MinipStaffBaseUrl string `yaml:"MinipStaffBaseUrl"`
	PayCenterBaseUrl  string `yaml:"PayCenterBaseUrl"`
	EsServiceUrl      string `yaml:"EsServiceUrl"`
	SentryDSN         string `yaml:"SentryDSN"`
	WfeBaseUrl        string `yaml:"WfeBaseUrl"`
}

var KafkaSetting = &Kafka{}

type Kafka struct {
	BootstrapServers string `yaml:"BootstrapServers"`
	Topic            string `yaml:"Topic"`
}

var RemoteSetting = &Remote{}

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"HttpPort"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
}

var ServerSetting = &Server{}

type Consul struct {
	Address string `yaml:"Address"`
	Token   string `yaml:"Token"`
}

var ConsulSetting = &Consul{}

type Database struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	PmsName     string `yaml:"PmsName"`
	TablePrefix string `yaml:"TablePrefix"`
}

var DatabaseSetting = &Database{}
var RecordDatabaseSetting = &Database{}
var GmtIDDatabaseSetting = &Database{}

type DatabaseFlow struct {
	Type        string `yaml:"Type"`
	User        string `yaml:"User"`
	Password    string `yaml:"Password"`
	Host        string `yaml:"Host"`
	Name        string `yaml:"Name"`
	TablePrefix string `yaml:"TablePrefix"`
}

var DatabaseFlowSetting = &DatabaseFlow{}

type Redis struct {
	Host        string        `yaml:"Host"`
	Password    string        `yaml:"Password"`
	MaxIdle     int           `yaml:"MaxIdle"`
	MaxActive   int           `yaml:"MaxActive"`
	IdleTimeout time.Duration `yaml:"IdleTimeout"`
	MachineryDB int           `yaml:"MachineryDB"`
}

var RedisSetting = &Redis{}

type Sso struct {
	Host           string `yaml:"Host"`
	WxSsoAppID     string `yaml:"WxSsoAppID"`
	WxSsoAppSecret string `yaml:"WxSsoAppSecret"`
}

var SsoSetting = &Sso{}

type Passport struct {
	Host string `yaml:"Host"`
}

var PassportSetting = &Passport{}

type Cas struct {
	BaseUrl string `yaml:"BaseUrl"`
	AppID   string `yaml:"AppID"`
	ApiKey  string `yaml:"ApiKey"`
}

var CasSetting = &Cas{}

type PayCenter struct {
	BaseUrl string `yaml:"BaseUrl"`
	AppID   string `yaml:"AppID"`
	ApiKey  string `yaml:"ApiKey"`
}

var PayCenterSetting = &PayCenter{}

type Urm struct {
	URL       string `yaml:"URL"`
	AppID     string `yaml:"AppID"`
	AppSecret string `yaml:"AppSecret"`
}

var UrmSetting = &Urm{}

type Wfe struct {
	BaseUrl string `yaml:"BaseUrl"`
	AppID   string `yaml:"AppID"`
	ApiKey  string `yaml:"ApiKey"`
}

var WfeSetting = &Wfe{}

type Aliyun struct {
	RegionID     string `yaml:"RegionID"`
	AccessID     string `yaml:"AccessID"`
	AccessSecret string `yaml:"AccessSecret"`
	SliderAppKey string `yaml:"SliderAppKey"`
}

var AliyunSetting = &Aliyun{}

type TencentCloud struct {
	// 帐户层面
	AppId     string `yaml:"AppId"`
	SecretId  string `yaml:"SecretId"`
	SecretKey string `yaml:"SecretKey"`

	// 验证码
	CaptchaAppId        uint64 `yaml:"CaptchaAppId"`
	CaptchaAppSecretKey string `yaml:"CaptchaAppSecretKey"`
	// 无感验证码
	SmartCaptchaAppId        uint64 `yaml:"SmartCaptchaAppId"`
	SmartCaptchaAppSecretKey string `yaml:"SmartCaptchaAppSecretKey"`
	// 对象存储cos
	CosBucket string `yaml:"CosBucket"`
	CosRegion string `yaml:"CosRegion"`
	//短信
	SmsAppID  string `yaml:"SmsAppID"`
	SmsAppKey string `yaml:"SmsAppKey"`
	SmsSignID string `yaml:"SmsSignID"`
}

var TencentCloudSetting = &TencentCloud{}

type Machinery struct {
	ResultBackend   string `yaml:"ResultBackend"`
	Broker          string `yaml:"Broker"`
	ResultsExpireIn int    `yaml:"ResultsExpireIn"`
}

var MachinerySetting = &Machinery{}

type Email struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

var EmailSetting = &Email{}

type Vision struct {
	BaseURL   string `yaml:"BaseURL"`
	AppID     string `yaml:"AppID"`
	AppSecret string `yaml:"AppSecret"`
}

var VisionSetting = &Vision{}

type YnkInvoice struct {
	AppId        string `yaml:"AppId"`
	AppSecret    string `yaml:"AppSecret"`
	AppHost      string `yaml:"AppHost"`
	DeviceNo     string `yaml:"DeviceNo"`
	DeviceTaxNo  string `yaml:"DeviceTaxNo"`
	CreatedBy    string `yaml:"CreatedBy"`
	DefaultEmail string `yaml:"DefaultEmail"`
}

var YnkInvoiceSetting = &YnkInvoice{}

type Approval struct {
	Url string `yaml:"Url"`
}

var ApprovalSetting = &Approval{}

type Fpms struct {
	BaseUrl string `yaml:"BaseUrl"`
}

var FpmsSetting = &Fpms{}

type Auth struct {
	AppID  string `yaml:"app_id"`
	Secret string `yaml:"secret"`
}

var AuthSetting = &Auth{}

var cfg *ini.File

func Setup(env *string) {
	var (
		iniFile string
	)

	switch *env {
	case "dev":
		iniFile = "conf/dev.ini"
	case "test":
		iniFile = "conf/test.ini"
	case "prod":
		iniFile = "conf/prod.ini"
	default:
		panic("invalid env")
	}

	//试图从yaml加载配置文件
	cnt, err := ioutil.ReadFile(iniFile)
	if err == nil && len(cnt) > 0 {
		fmt.Printf("ini文件存在，现从ini开始加载\n")
		setupFromIni(iniFile)
	} else {
		fmt.Printf("ini文件不存在\n")
	}
}

//增加ini文件的名字的支持
func SetupV2(env *string, projectName *string) {
	var (
		iniFile string
	)

	switch *env {
	case "dev":
		iniFile = fmt.Sprintf("conf/%s.dev.ini", *projectName)
	case "test":
		iniFile = fmt.Sprintf("conf/%s.test.ini", *projectName)
	case "prod":
		iniFile = fmt.Sprintf("conf/%s.prod.ini", *projectName)
	default:
		panic("invalid env")
	}

	//试图从yaml加载配置文件
	cnt, err := ioutil.ReadFile(iniFile)
	if err == nil && len(cnt) > 0 {
		fmt.Printf("ini文件存在，现从ini开始加载\n")
		setupFromIni(iniFile)
	} else {
		fmt.Printf("ini文件不存在\n")
	}
}

//从ini加载配置文件
func setupFromIni(cfgFile string) {
	cfgHandler, err := ini.Load(cfgFile)
	if err != nil {
		fmt.Printf("setting.Setup, fail to parse '%s': %v\n", cfgFile, err)
	} else {
		fmt.Printf("setting.Setup, parse '%s' success\n", cfgFile)
	}
	cfg = cfgHandler

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("consul", ConsulSetting)
	mapTo("database", DatabaseSetting)
	mapTo("database_record", RecordDatabaseSetting)
	mapTo("database_gmtid", GmtIDDatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("kafka", KafkaSetting)
	mapTo("sso", SsoSetting)
	mapTo("passport", PassportSetting)
	mapTo("database_flow", DatabaseFlowSetting)
	mapTo("remote", RemoteSetting)
	mapTo("cas", CasSetting)
	mapTo("urm", UrmSetting)
	mapTo("wfe", WfeSetting)
	mapTo("aliyun", AliyunSetting)
	mapTo("tencent_cloud", TencentCloudSetting)
	mapTo("pay_center", PayCenterSetting)
	mapTo("machinery", MachinerySetting)
	mapTo("email", EmailSetting)
	mapTo("vision", VisionSetting)
	mapTo("ynk_invoice", YnkInvoiceSetting)
	mapTo("approval", ApprovalSetting)
	mapTo("fpms", FpmsSetting)
	mapTo("auth", AuthSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("Cfg.MapTo %s err: %v\n", section, err)
	}
}
