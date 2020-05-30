package server

// add by stefan

import (
	"GameServer/LogicMsg"
	"GameServer/dbo"
	"GameServer/rpc"
	"flag"
	"syscall"

	"github.com/Peakchen/xgameCommon/Config/LogicConfig"
	"github.com/Peakchen/xgameCommon/Config/serverConfig"
	"github.com/Peakchen/xgameCommon/HotUpdate"
	"github.com/Peakchen/xgameCommon/Kcpnet"
	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/define"
)

func init() {
	var CfgPath string
	flag.StringVar(&CfgPath, "serverconfig", "serverconfig", "default path for configuration files")
	serverConfig.LoadSvrAllConfig(CfgPath)
	dbStatistics.InitDBStatistics()
	LogicMsg.Init()
	rpc.Init()
}

func reloadConfig() {
	LogicConfig.LoadLogicAll()
}

func StartServer() {
	Gamecfg := serverConfig.GGameconfigConfig.Get()
	server := Gamecfg.Zone + Gamecfg.No
	dbo.StartDBSerice(server)
	// for kill pid to emit signal to do action...
	HotUpdate.RunHotUpdateCheck(&HotUpdate.TServerHotUpdateInfo{
		Recvsignal: syscall.SIGTERM,
		HUCallback: reloadConfig,
	})
	gameSvr := Kcpnet.NewKcpClient(Gamecfg.Name,
		Gamecfg.Listenaddr,
		Gamecfg.Pprofaddr,
		define.ERouteId_ER_Game)

	gameSvr.Run()
	dbStatistics.DBStatisticsStop()
}
