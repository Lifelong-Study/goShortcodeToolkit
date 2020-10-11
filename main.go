package goShortcodeToolkit

import (
	"log"
	"time"
)

// 設置鬧鐘在特定時間
func NewAlarmAtSpecificTime(hour, minute int) *time.Timer {
	var now time.Time
	now = time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local)

	//
	start := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)

	// 當前時間比任務起始時間早 -> 傳回起始時間
	if now.Before(start) {
		log.Printf("任務將會在 %v 時開始輪詢", start)
		return time.NewTimer(start.Sub(time.Now()))
	} else {
		tomorrowNow := time.Date(now.Year(), now.Month(), now.Day() + 1, hour, minute, 0, 0, time.Local)
		log.Printf("任務將會在 %v 時開始輪詢", tomorrowNow)
		return time.NewTimer(tomorrowNow.Sub(time.Now()))
	}
}

// 設置鬧鐘，從特定時間開始，依循間隔分鐘是設置距離當前時間最靠近的鬧鐘
// 如果當前時間比鬧鈴時間早，則傳回鬧鈴時間；反之，傳回依間隔分鐘數且距離當前時間最近的時間
func NewAlarmAtSpecificTimeAndInterval(hour, minute, interval int) *time.Timer {
	var target int // 目標分鐘數
	var offset int // 差距分鐘數

	//
	var now time.Time
	now = time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local)

	//
	start := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, time.Local)

	// 當前時間比任務起始時間早 -> 傳回起始時間
	if now.Before(start) {
		log.Printf("任務將會在 %v 時開始輪詢", start)
		return time.NewTimer(start.Sub(time.Now()))
	}

	//
	if interval > now.Minute() {
		for {
			new := start.Add(time.Duration(interval) * time.Minute)

			if new.Before(now) {
				return time.NewTimer(new.Sub(time.Now()))
			}
		}
	} else {
		target = (1 + (now.Minute() / interval)) * interval
		offset = target - now.Minute()

		new := now.Add(time.Duration(offset) * time.Minute)

		log.Printf("任務將會在 %v 時開始輪詢", new)

		return time.NewTimer(new.Sub(time.Now()))
	}
}
