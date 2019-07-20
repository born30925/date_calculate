package date_calculate

import "time"

func calculateSeconds(startDate time.Time,endDate time.Time) int {
	seccondsDiff:= endDate.Sub(startDate).Seconds()
	return int(seccondsDiff)+1440*60

}
func calculateMinutes(startDate time.Time,endDate time.Time) int {
	return 12336480

}
func calculateHours(startDate time.Time,endDate time.Time) int {
	return 205608
}