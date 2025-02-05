package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cast"

	"github.com/lekluge/gosumemory/config"

	"github.com/lekluge/gosumemory/mem"
	"github.com/lekluge/gosumemory/memory"
	"github.com/lekluge/gosumemory/pp"
	"github.com/lekluge/gosumemory/updater"
	"github.com/lekluge/gosumemory/web"
)

func main() {
	config.Init()
	updateTimeFlag := flag.Int("update", cast.ToInt(config.Config["update"]), "How fast should we update the values? (in milliseconds)")
	shouldWeUpdate := flag.Bool("autoupdate", true, "Should we auto update the application?")
	isRunningInWINE := flag.Bool("wine", cast.ToBool(config.Config["wine"]), "Running under WINE?")
	songsFolderFlag := flag.String("path", config.Config["path"], `Path to osu! Songs directory ex: /mnt/ps3drive/osu\!/Songs`)
	memDebugFlag := flag.Bool("memdebug", cast.ToBool(config.Config["memdebug"]), `Enable verbose memory debugging?`)
	memCycleTestFlag := flag.Bool("memcycletest", cast.ToBool(config.Config["memcycletest"]), `Enable memory cycle time measure?`)
	disablecgo := flag.Bool("cgodisable", cast.ToBool(config.Config["cgodisable"]), `Disable everything non memory-reader related? (pp counters)`)
	flag.Parse()
	cgo := *disablecgo
	mem.Debug = *memDebugFlag
	memory.MemCycle = *memCycleTestFlag
	memory.UpdateTime = *updateTimeFlag
	memory.SongsFolderPath = *songsFolderFlag
	memory.UnderWine = *isRunningInWINE
	if runtime.GOOS != "windows" && memory.SongsFolderPath == "auto" {
		log.Fatalln("Please specify path to osu!Songs (see --help)")
	}
	if memory.SongsFolderPath != "auto" {
		if _, err := os.Stat(memory.SongsFolderPath); os.IsNotExist(err) {
			log.Fatalln(`Specified Songs directory does not exist on the system! (try setting to "auto" if you are on Windows or make sure that the path is correct)`)
		}
	}
	if *shouldWeUpdate == true {
		updater.DoSelfUpdate()
	}

	go memory.Init()
	// err := db.InitDB()
	// if err != nil {
	// 	log.Println(err)
	// 	time.Sleep(5 * time.Second)
	// 	os.Exit(1)
	// }
	go web.SetupStructure()
	go web.SetupRoutes()
	if !cgo {
		go pp.GetData()
		go pp.GetFCData()
		go pp.GetMaxData()
		go pp.GetEditorData()
	}
	web.HTTPServer()

}
