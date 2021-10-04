package absPathFinder

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func getAbsPath()string{

	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	exePath := filepath.Dir(ex)
	// fmt.Println("absPath:", dir)

	return exePath
}

func getWorkingDirPath()string{
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func GetCurrentAbPath() string {
	dir := getAbsPath() // 執行黨絕對路徑
	tmpDir :=os.TempDir() // run go 產出的Temp 資料夾位置
	// fmt.Println(tmpDir)
	if strings.Contains(dir, tmpDir) {
		return getWorkingDirPath()
	}
	return dir
}