package common

import "github.com/pkg/errors"

const (
	SUCCESS                       = 200   //成功
	ERROR                         = 500   //失败
	InvalidParams                 = 400   //参数错误
	NotLoggedIn                   = 1000  // 未登录
	ParameterIllegal              = 1001  // 参数不合法
	UnauthorizedUserId            = 1002  // 用户Id不合法
	Unauthorized                  = 1003  // 未授权
	ServerError                   = 1004  // 系统错误
	NotData                       = 1005  // 没有数据
	ModelAddError                 = 1006  // 添加错误
	ModelDeleteError              = 1007  // 删除错误
	ModelStoreError               = 1008  // 存储错误
	OperationFailure              = 1009  // 操作失败
	RoutingNotExist               = 1010  // 路由不存在
	ErrorUserExist                = 1011  // 用户已存在
	ErrorUserNotExist             = 1012  // 用户不存在
	ErrorNeedCaptcha              = 1013  // 需要验证码
	PasswordInvalid               = 1014  //密码不符合规范
	ErrorCheckTokenFail           = 10001 //用户信息获取失败
	ErrorCheckUserRoleFail        = 10002 // 用户信息获取失败
	ErrorReportDeviceInsertFail   = 20001 //上报设备信息失败
	ErrorDeviceUpdateFail         = 20002 // 设备更新失败
	ErrorReportDeviceUpsertFail   = 20003 //上报设备信息失败
	ErrorGetTrafficMapFail        = 20004 //获取人流地图数据失败
	ErrorDeviceDeleteFail         = 20005 //设备删除失败
	ErrorReportDeviceMonitorFail  = 20006 // 算法上报数据失败
	ErrorAlarmInsertFail          = 30001 //报警上报失败
	ErrorAlarmDealFail            = 30002 //报警处理失败
	ErrorAddPostConfigFail        = 40001 //查岗配置添加失败
	ErrorAddPatrolConfigFail      = 40002 //巡更点配置失败
	ErrorUpdatePatrolConfigFail   = 40003 //巡更更新失败
	ErrorAddPatrolRecordFail      = 40004 //巡查失败
	ErrorNoPatrolConfig           = 40005 //没有巡查点配置
	ErrorSendWarningCheckAuthFail = 40007 // 一键报警验证失败
	ErrorUploadSaveImageFail      = 50000 //图片上传失败
	ErrorUploadCheckImageFormat   = 50001 //格式错误
	ErrorUploadCheckImageFail     = 50002 //格式错误
	ErrorRTSPUrlInvalid           = 10003 //视频流地址错误
	LicenseNotExist               = 10004 //车牌号不存在
	ReachMaxPlayCount      		  = 10005 //到达播放上限
	LicenseExist                  = 10006 //车牌号不存在
)

var MsgFlags = map[int32]string{
	SUCCESS:                       "成功",
	ERROR:                         "失败",
	InvalidParams:                 "参数不合法",
	ErrorCheckTokenFail:           "用户校验失败，请尝试重新登录（10001）",
	ErrorCheckUserRoleFail:        "用户不存在",
	ErrorReportDeviceInsertFail:   "上报设备信息失败",
	ErrorReportDeviceUpsertFail:   "上报设备信息失败",
	ErrorGetTrafficMapFail:        "获取人流地图数据失败",
	ErrorDeviceUpdateFail:         "设备更新失败",
	ErrorDeviceDeleteFail:         "设备更新失败",
	ErrorAlarmInsertFail:          "报警上报失败",
	ErrorAlarmDealFail:            "报警上报失败",
	ErrorAddPostConfigFail:        "添加查岗配置失败",
	ErrorAddPatrolConfigFail:      "添加巡更配置失败",
	ErrorUpdatePatrolConfigFail:   "更新巡更配置失败",
	ErrorAddPatrolRecordFail:      "巡更失败",
	ErrorUploadSaveImageFail:      "上传图片失败",
	ErrorReportDeviceMonitorFail:  "摄像头数据上报失败",
	ErrorSendWarningCheckAuthFail: "一键报警验证失败",
	NotLoggedIn:                   "没有登录",
	ParameterIllegal:              "参数错误",
	UnauthorizedUserId:            "用户验证失败",
	Unauthorized:                  "用户验证失败",
	ErrorNeedCaptcha:              "用户验证失败",
	NotData:                       "无数据",
	ServerError:                   "服务器错误",
	ModelAddError:                 "添加失败",
	ModelDeleteError:              "删除失败",
	ModelStoreError:               "存储失败",
	OperationFailure:              "操作失败",
	RoutingNotExist:               "路由不存在",
	ErrorUserExist:                "用户已存在",
	ErrorUserNotExist:             "用户不存在",
	ErrorRTSPUrlInvalid:           "rtsp流地址错误",
	LicenseNotExist:               "车牌号不存在",
	ReachMaxPlayCount:             "播放视频到达上限",
	LicenseExist:                  "车牌号已存在",
	PasswordInvalid:               "密码不符合规范",
}

var MongoNoFound = errors.New("MONGO NO FOUND")
var MongoInsertError = errors.New("MONGO INSERT ERROR")
var DeviceNoFound = errors.New("DEVICE NO FOUND")
var FileNoFound = errors.New("FILE NO FOUND")
var StatusError = errors.New("STATUS ERROR")
var NoNeedInsert = errors.New("NO NEED INSERT")
var NoData = errors.New("NO DATA")

// GetMsg get error information based on Code
func GetMsg(code int32) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
