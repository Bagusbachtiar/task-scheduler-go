package main 

import (
	"context"
	"log"
	"time"
	"os"
	"task-scheduler/scheduler"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)


func main() { 
	log.Println("⏰ Task Scheduler Engine Initializing...")

	//load environtment paarameters from local sevret .env files
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: no .env configuration file found. reading from system environment variables.")
	}

	dbConnectionString := os.Getenv("DATABASE_URL")
	if dbConnectionString == "" {
		log.Fatalf("critical Error: DATABASE_URL is missing from environtmet variables.")
	}

	// initialize high-performance postgres connection pool
	dbPool, err := pgxpool.New(context.Background(), dbConnectionString)
	if err != nil {
		log.Fatalf("unable to allocate pool resource context parameters: %v, err")
	}
	defer dbPool.Close()

	//verify target link connection is active 
	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatalf("Neon cloud target link ping operation failed: %v", err)
	}
	log.Println("Succesfully connected to Neon serverless Psotgres Cluster!")

	//Set up out rhytm clock tocker to pulse every 10 seconds
	// 1.1 Create a native Go Ticker that triggers a pulse every 2 seconds
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop() //Cean up the ticker when the program exits

	log.Println("🚀 Time-Driven Engine Online. Scanning for pending records.")

	//2. an infinite lopp that blocks and listens to the tickers internal channel 
	// every time 2 second passes, the ticker sends a message through 'ticker.C'
	for pulseTime := range ticker.C {
		log.Printf("\n⏱️ [Heartbeat Clock tick]: %v", pulseTime.Format("15:04:05"))
	
		//Database qUERY: SCAN for uncompleted tasks where execution time is less than or equal to now
		query := `
				SELECT id, phone_number, message, execution_time 
				FROM scheduled_tasks
				WHERE is_completed = FALSE AND execution_time <= $1;
			`

		rows, err := dbPool.Query(context.Background(), query, pulseTime)
		if err != nil {
			log.Printf("[Database Error]: Failed to poll pending records:  %v", err)
		    continue 
		}

		//Store slice array of tasks due in this current pulse cycle
		var tasksDue []scheduler.ScheduleTask
		for rows.Next() {
			var task scheduler.ScheduleTask
			err := rows.Scan(&task.ID, &task.PhoneNumber, &task.Message, &task.ExecutionTZ)
			if err != nil {
				log.Printf("[Processing Error]: Failed to scan record: %v", err)
				continue
			}
			tasksDue = append(tasksDue, task)
		}
		rows.Close() //Close the rows iterator to free up resources

		if len(tasksDue) == 0 {
			log.Println("😴 No tasks due in this cycle. Waiting for next tick...")
			continue
		}

		log.Printf("🔥 Found %d due tasks! Processing batch payload...", len (tasksDue))

		//PROCESS BATCH PIPELINE CONCURRENTLY
		for _, task := range tasksDue {
			//Spawns clean concurrent goroutine tracking calls directly to our refactored engine package 
			go scheduler.ProcessTasks(dbPool, task)
		}
	}
}