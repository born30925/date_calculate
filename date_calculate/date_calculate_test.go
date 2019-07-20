package date_calculate

import (
	"testing"
	"time"
)

func Test_calculateSeconds_Input_startDate_5_2_1996_endDate_20_7_2019_Should_Be_740188800(t * testing.T)  {


	expected := 740188800
	startDate := time.Date(1996,2,5,0,0,0,0,time.UTC)
	endDate :=  time.Date(2019,7,20,0,0,0,0,time.UTC)

	actual :=  calculateSeconds(startDate,endDate)


	if actual != expected {
		t.Errorf("Expected %d but it got %d",expected,actual)
	}

}
func Test_calculateMinutes_Input_startDate_5_2_1996_endDate_20_7_2019_Should_Be_12336480(t * testing.T)  {


	expected := 12336480
	startDate := time.Date(1996,2,5,0,0,0,0,time.UTC)
	endDate :=  time.Date(2019,7,20,0,0,0,0,time.UTC)

	actual :=  calculateMinutes(startDate,endDate)


	if actual != expected {
		t.Errorf("Expected %d but it got %d",expected,actual)
	}

}
func Test_calculateHours_Input_startDate_5_2_1996_endDate_20_7_2019_Should_Be_205608(t * testing.T)  {


	expected := 205608
	startDate := time.Date(1996,2,5,0,0,0,0,time.UTC)
	endDate :=  time.Date(2019,7,20,0,0,0,0,time.UTC)

	actual :=  calculateHours(startDate,endDate)


	if actual != expected {
		t.Errorf("Expected %d but it got %d",expected,actual)
	}

}

func Test_calculateSeconds_Input_startDate_2_8_1995_endDate_20_7_2019_Should_Be_756345600(t * testing.T)  {


	expected := 756345600
	startDate := time.Date(1995,8,2,0,0,0,0,time.UTC)
	endDate :=  time.Date(2019,7,20,0,0,0,0,time.UTC)

	actual :=  calculateSeconds(startDate,endDate)


	if actual != expected {
		t.Errorf("Expected %d but it got %d",expected,actual)
	}

}