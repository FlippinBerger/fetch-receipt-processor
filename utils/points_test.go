package utils

import (
	"testing"
)

func TestGetAlphaNumericCount(t *testing.T) {
	count := GetAlphaNumericCount("hello")
	if count != 5 {
		t.Errorf("got %d expected %d", count, 5)
	}

	count = GetAlphaNumericCount("hello1")
	if count != 6 {
		t.Errorf("got %d expected %d", count, 6)
	}

	count = GetAlphaNumericCount("hello 1")
	if count != 6 {
		t.Errorf("got %d expected %d", count, 6)
	}

	count = GetAlphaNumericCount("hello!@#")
	if count != 5 {
		t.Errorf("got %d expected %d", count, 5)
	}

	count = GetAlphaNumericCount("")
	if count != 0 {
		t.Errorf("got %d expected %d", count, 0)
	}

	count = GetAlphaNumericCount("      ")
	if count != 0 {
		t.Errorf("got %d expected %d", count, 0)
	}

	count = GetAlphaNumericCount("&@$&  @^#")
	if count != 0 {
		t.Errorf("got %d expected %d", count, 0)
	}

	count = GetAlphaNumericCount("       hello    ")
	if count != 5 {
		t.Errorf("got %d expected %d", count, 5)
	}
}

func TestGetCentsPoints(t *testing.T) {
	points := GetCentsPoints("")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetCentsPoints("100")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetCentsPoints("hello")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetCentsPoints("1.00")
	if points != 75 {
		t.Errorf("got %d expected %d", points, 75)
	}

	points = GetCentsPoints("1.12")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetCentsPoints("1.25")
	if points != 25 {
		t.Errorf("got %d expected %d", points, 25)
	}

	points = GetCentsPoints("1.50")
	if points != 25 {
		t.Errorf("got %d expected %d", points, 25)
	}

	points = GetCentsPoints("1.75")
	if points != 25 {
		t.Errorf("got %d expected %d", points, 25)
	}
}

func TestGetPairPoints(t *testing.T) {
	points := GetPairPoints(0)
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetPairPoints(1)
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetPairPoints(2)
	if points != 5 {
		t.Errorf("got %d expected %d", points, 5)
	}

	points = GetPairPoints(3)
	if points != 5 {
		t.Errorf("got %d expected %d", points, 5)
	}

	points = GetPairPoints(4)
	if points != 10 {
		t.Errorf("got %d expected %d", points, 10)
	}
}

func TestIsOddDayInDateString(t *testing.T) {
	isOdd := IsOddDayInDateString("")
	if isOdd {
		t.Errorf("shouldn't be odd")
	}

	isOdd = IsOddDayInDateString("03")
	if isOdd {
		t.Errorf("shouldn't be odd")
	}

	isOdd = IsOddDayInDateString("2020-03")
	if isOdd {
		t.Errorf("shouldn't be odd")
	}

	isOdd = IsOddDayInDateString("2020-03-01")
	if !isOdd {
		t.Errorf("should be odd")
	}

	isOdd = IsOddDayInDateString("2020-03-03")
	if !isOdd {
		t.Errorf("should be odd")
	}

	isOdd = IsOddDayInDateString("2020-03-3")
	if !isOdd {
		t.Errorf("should be odd")
	}

	isOdd = IsOddDayInDateString("2020-03-22")
	if isOdd {
		t.Errorf("shouldn't be odd")
	}

	isOdd = IsOddDayInDateString("2020-03-25")
	if !isOdd {
		t.Errorf("should be odd")
	}
}

func TestIsBetween2And4(t *testing.T) {
	between := IsBetweenTwoAndFourPM("")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("2:01")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("14:00")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("14:01")
	if !between {
		t.Errorf("should be between")
	}

	between = IsBetweenTwoAndFourPM("15:59")
	if !between {
		t.Errorf("should be between")
	}

	between = IsBetweenTwoAndFourPM("16:00")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("22:00")
	if between {
		t.Errorf("shouldn't be between")
	}

	between = IsBetweenTwoAndFourPM("00:00")
	if between {
		t.Errorf("shouldn't be between")
	}
}

func TestItemDescriptionPoints(t *testing.T) {
	points := GetItemDescriptionPoints("", "")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetItemDescriptionPoints("", "1.00.00")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetItemDescriptionPoints("qqq", "1.00.00")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetItemDescriptionPoints("qqq", "1.00")
	if points != 1 {
		t.Errorf("got %d expected %d", points, 1)
	}

	points = GetItemDescriptionPoints("qqq qq", "1.00")
	if points != 1 {
		t.Errorf("got %d expected %d", points, 1)
	}

	points = GetItemDescriptionPoints("qqq qq", "10.00")
	if points != 2 {
		t.Errorf("got %d expected %d", points, 2)
	}

	points = GetItemDescriptionPoints("qqq qqq", "10.00")
	if points != 0 {
		t.Errorf("got %d expected %d", points, 0)
	}

	points = GetItemDescriptionPoints("qqq qq", "10.50")
	if points != 3 {
		t.Errorf("got %d expected %d", points, 3)
	}
}
