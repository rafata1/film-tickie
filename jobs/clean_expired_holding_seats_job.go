package jobs

import (
	"fmt"
	"github.com/templateOfService/connectors/mysql"
	"github.com/templateOfService/models"
	"time"
)

var cleaningQuery = "UPDATE orders SET status = ? WHERE status = ? AND created_at < ?"

func CleanExpiredHoldingSeats() {
	conn := mysql.GetMySQLInstance()
	for {
		time.Sleep(5 * time.Second)
		now := time.Now()
		_, err := conn.Exec(cleaningQuery, models.Expired, models.Holding, now.Add(-10*time.Minute))
		if err != nil {
			fmt.Println("error cleaning expired seats")
		}
	}
}
