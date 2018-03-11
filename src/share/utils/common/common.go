package common

func SwitchMonth(month string) int64{

	var mon int64
	switch month {
	case "January":
		mon = 1
	case "February":
		mon = 2
	case "March":
		mon = 3
	case "April":
		mon = 4
	case "May":
		mon = 5
	case "June":
		mon = 6
	case "July":
		mon = 7
	case "August":
		mon = 8
	case "September":
		mon = 9
	case "October":
		mon = 10
	case "November":
		mon = 11
	case "December":
		mon = 12
	default:
		mon = 0
	}
	return mon
}
