package cronjob

import (
	"log"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestCronjob(t *testing.T) {
	c := cron.New()
	err := c.AddFunc("* * * * *", func() {
		log.Println("tugas cron dijalankan", time.Now())
	})

	if err != nil {
		log.Println("gagal menjalankan cron", err)
		return
	}

	c.Start()
	// menunggu program berjalan selama 1 menit
	time.Sleep(1 * time.Minute)

	// menghentikan cron
	c.Stop()
}
