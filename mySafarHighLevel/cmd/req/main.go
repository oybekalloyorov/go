package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"mySafarHighLevel/internal/client"
	"mySafarHighLevel/internal/config"
	"mySafarHighLevel/internal/util"
)

func main() {
	cfg := config.Load()

	cl := client.New(cfg)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout+5)*time.Second)
	defer cancel()

	resp, err := cl.GetRecommendations(ctx)
	if err != nil {
		log.Fatalf("API error: %v", err)
	}

	if len(resp.Recommendations) == 0 {
		fmt.Println("No flights found (recommendations is null or empty). Try changing date/class/filters.")
		return
	}

	util.PrintPrettyJSON(resp)
}
