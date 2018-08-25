package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetWindowIcon(gui.NewQIcon5("qrc:/qml/icon.ico"))

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/qml/application.qml", 0))
	gui.QGuiApplication_Exec()
}
