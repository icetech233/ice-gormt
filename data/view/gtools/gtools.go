package gtools

import (
	"fmt"
	"os/exec"

	"github.com/xxjwxc/public/mylog"

	"github.com/icetech233/ice-gormt/data/dlg"
	"github.com/icetech233/ice-gormt/data/view/model"

	"github.com/icetech233/ice-gormt/data/config"

	"github.com/icetech233/ice-gormt/data/view/model/genmysql"
	"github.com/xxjwxc/public/tools"
)

// Execute exe the cmd
func Execute() {
	if config.GetIsGUI() {
		dlg.WinMain()
	} else {
		showCmd()
	}
}

func showCmd() {
	// var tt oauth_db.UserInfoTbl
	// tt.Nickname = "ticket_001"
	// orm.Where("nickname = ?", "ticket_001").Find(&tt)
	// fmt.Println(tt)
	var modeldb model.IModel
	switch config.GetDbInfo().Type {
	case 0: // mysql
		modeldb = genmysql.GetModel()
	}
	if modeldb == nil {
		mylog.Error(fmt.Errorf("modeldb not fund : please check db_info.type (0:mysql , 1:sqlite , 2:mssql) "))
		return
	}

	pkg := modeldb.GenModel()
	// gencnf.GenOutPut(&pkg)
	// just for test
	// out, _ := json.Marshal(pkg)
	// tools.WriteFile("test.txt", []string{string(out)}, true)

	list, _ := model.Generate(pkg)

	for _, v := range list {
		path := config.GetOutDir() + "/" + v.FileName
		tools.WriteFile(path, []string{v.FileCtx}, true)

		mylog.Info("formatting differs from goimport's:")
		cmd, _ := exec.Command("goimports", "-l", "-w", path).Output()
		mylog.Info(string(cmd))

		mylog.Info("formatting differs from gofmt's:")
		cmd, _ = exec.Command("gofmt", "-l", "-w", path).Output()
		mylog.Info(string(cmd))
	}
}
