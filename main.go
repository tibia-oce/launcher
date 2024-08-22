package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"launcher/internal/config"
	"launcher/internal/launcher"
	"launcher/internal/logger"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	executable, err := os.Executable()
	if err != nil {
		fmt.Printf("Failed to get executable path: %v", err)
		os.Exit(1)
	}

	executable = filepath.Base(executable)
	appName := strings.TrimSuffix(executable, filepath.Ext(executable))
	baseURL := "https://raw.githubusercontent.com/luan/tibia-client/main/"

	logger.Init("info")
	cfg := config.LoadConfig(appName)

	app := launcher.NewApp(baseURL, appName, cfg)

	if err := app.DoUpdate(baseURL); err != nil {
		logger.Error(fmt.Errorf("Failed to update: %v", err))
		os.Exit(1)
	}

	err = wails.Run(&options.App{
		Title:  appName + " Launcher",
		Width:  760,
		Height: 440,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		DisableResize:    true,
		Frameless:        true,
		BackgroundColour: &options.RGBA{R: 13, G: 25, B: 51, A: 0},
		OnStartup:        app.Startup,
		Windows: &windows.Options{
			ZoomFactor:           1.0,
			WebviewIsTransparent: true,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: true,
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
