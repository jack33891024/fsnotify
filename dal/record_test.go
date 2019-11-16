package db

import (
	"testing"
	"time"

	"github.com/learngo/inotify_directory/model"
)

func init() {
	dsn := "root:1234qwer@tcp(192.168.0.199:3306)/monitor_dir?parseTime=true"
	err := InitDB(dsn)
	if err != nil {
		return
	}
}

func TestInsertOperatorRecord(t *testing.T) {
	record := &model.OpRecord{}
	record.EventTime = time.Now().Format("2006/01/02 15:04:05")
	record.Operator = "创建文件"
	record.FilePath = "/data/app/vkgame/public/uploads/status.php"

	recordId, err := InsertOperatorRecord(record)
	if err != nil {
		t.Fatalf("insert record failed, err: %v\n", err)
	}

	t.Logf("insert success, insert id: %d\n", recordId)
}
