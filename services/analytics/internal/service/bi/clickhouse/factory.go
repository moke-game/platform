package clickhouse

import (
	"go.uber.org/zap"

	"github.com/moke-game/platform.git/services/analytics/internal/service/bi"
	"github.com/moke-game/platform.git/services/analytics/internal/service/bi/clickhouse/internal"
)

func NewDataProcessor(
	logger *zap.Logger, rootPath, addr, dbname, uname, passwd string,
) (processor bi.DataProcessor, err error) {
	p := new(internal.Processor)
	if e := p.Init(logger, rootPath, addr, dbname, uname, passwd); e != nil {
		err = e
	} else {
		processor = p
	}
	return
}

type UserGuide struct {
	Uid         string `json:"uid"`          //用户ID
	DeviceId    string `json:"device_id"`    //设备ID
	Guide       string `json:"guide"`        //引导名称
	GuideName   string `json:"gname"`        //引导名称
	GuideStep   int32  `json:"step"`         //引导步骤
	IsGuest     int32  `json:"is_guest"`     //是否游客
	CreateDay   string `json:"create_day"`   //创建日期
	CreateTime  string `json:"create_time"`  //创建时间
	IP          string `json:"ip"`           //ip		String
	System      int32  `json:"system"`       //操作系统 0:u3d 1:android 2:ios
	CountryCode string `json:"country_code"` //国家代码		String
	Country     string `json:"country"`      //国家		String		用户所在国家，根据 IP 地址生成							客户端
	Province    string `json:"province"`     //省份		String		用户所在省份，根据 IP 地址生成							客户端
	City        string `json:"city"`         //城市		string		用户所在城市，根据 IP 地址生成							客户端
	Brand       string `json:"brand"`        //手机型号		String		手机型号，例如：samsung，oppo，vivo							客户端
	Network     string `json:"network"`      //网络类型		String		网络类型，例如：Wifi							客户端
	Language    string `json:"language"`     //系统语言		String		系统语言，例如：中文简体							客户端
}

type step struct {
	Key   int32
	Value string
}

var Guides = map[string]*step{
	"HotfixUIOpen":                      {1, "进度条界面打开"},
	"HotfixVersionStart":                {2, "加载资源标记开始"},
	"HotfixVersionFinish":               {3, "加载资源标记结束"},
	"HotfixManifestStart":               {4, "加载清单开始"},
	"HotfixManifestFinish":              {5, "加载清单结束"},
	"HotfixUpdateStart":                 {6, "更新开始"},
	"HotfixUpdateFinish":                {7, "更新结束"},
	"HotfixUIClose":                     {8, "进度条界面关闭"},
	"LaunchStart":                       {9, "游戏启动"},
	"LaunchPrelaodStart":                {10, "预加载开始"},
	"LaunchPrelaodFInish":               {11, "预加载结束"},
	"LaunchFinish":                      {12, "启动结束"},
	"RegisterPageDisplayRequestedEvent": {13, "注册界面调起请求"},
	"RegisterPageDisplayedEvent":        {14, "注册界面调起结果"},
	"RegisteredRequestEvent":            {15, "注册请求事件"},
	"RegisteredEvent":                   {16, "注册结果"},
	"LoginDsiplay":                      {17, "登陆界面展示"},
	"LoginPageDisplayRequestedEvent":    {18, "打开UMS的登陆界面"},
	"LoginPageDisplayedEvent":           {19, "登录界面调起结果"},
	"LoginRequestEvent":                 {20, "登陆请求事件"},
	"LoginConnectStart":                 {21, "开始连接游戏服务器"},
	"LoggedinEvent":                     {22, "登录结果"},
	"LoginConnectFinish":                {23, "游戏服务器连接完成"},
	"LoginStart":                        {24, "开始登陆"},
	"LoginFinish":                       {25, "登陆结束"},
	"GuideHandHold":                     {26, "展示握持手机的方式"},
	"GuidePreambleInfo":                 {27, "游戏背景介绍"},
	"LoginDone":                         {28, "登陆完成进入游戏"},
	"GuideMoveGuide":                    {29, "移动引导, 拖动"},
	"GuideAttackGuide":                  {30, "普通攻击引导, 点击"},
	"GuideSkillGuide":                   {31, "技能攻击引导, 拖动"},
	"GuideHealthGuide":                  {32, "血量引导"},
	"GuideFlashGuide":                   {33, "闪现引导"},
	"GuideSpecialGuide":                 {34, "大招引导"},
	"GuideBattleFinish":                 {35, "战斗引导结束"},
	"GuideTaskReward":                   {36, "领取任务奖励"},
	"GuideHeroListEnter":                {37, "点击英雄按钮"},
	"GuideHeroSelectEnter":              {38, "点击选中主角当前英雄"},
	"GuideHeroUpgradeEnter":             {39, "英雄升级进入"},
	"GuideHeroUpgrade":                  {40, "英雄升级"},
	"GuideHeroUpgradeExit":              {41, "英雄升级退出"},
	"GuideHeroSelectExit":               {42, "英雄升级退出"},
	"GuideBattleMatch":                  {43, "匹配进入"},
	"GuideBattleResultStart":            {44, "完成首次匹配战斗看到结算页面"},
	"GuideBattleResultExit":             {45, "关闭新手引导结算页面返回主城"},
	"GuideBattleResultReplay":           {46, "关闭新手引导结算页面再来一场"},
	"GuideMedalRoadEnter":               {47, "打开成长之路"},
	"GuideGuideOver":                    {48, "引导结束"},
}
