package pkg

import "os"

func main() {
	CreateCalendar()
}

func CreateCalendar() {
	file, _ := os.Create("calendar.ics")
	cal := ""
	cal += "BEGIN:VCALENDAR\n" +
		"VERSION:2.0\n" +
		"PRODID:-//id//NONSGML v1.0//EN\n" +
		"CALSCALE:GREGORIAN\n" +
		"METHOD:PUBLISH\n"

	CreateEvent(&cal)
	cal += "END:VCALENDAR"
	file.WriteString(cal)
}

func CreateEvent(cal *string) {
	*cal += "BEGIN:VEVENT\nUID:12345678-1234-1234-1234-123456789012\nDTSTART:20240816T150000Z\nSUMMARY:title\nDESCRIPTION:description\nEND:VEVENT\n"
}
