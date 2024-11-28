import { type IScheduleItem, type ITimeTable } from "./constants";

const timeSlots = [
  { time_start: "09-00", time_end: "10-30" },
  { time_start: "10-40", time_end: "12-10" },
  { time_start: "12-40", time_end: "14-10" },
  { time_start: "14-20", time_end: "15-50" },
  { time_start: "16-20", time_end: "17-50" },
  { time_start: "18-00", time_end: "19-30" },
];

export function transformSchedule(schedule: IScheduleItem[]): ITimeTable[] {

  const scheduleByDate: { [date: string]: IScheduleItem[] } = {};
  // Group entries by date
  for (const entry of schedule) {
    const date = entry.date;
    if (!scheduleByDate[date]) {
      scheduleByDate[date] = [];
    }
    scheduleByDate[date].push(entry);
  }

  // Transform and sort
  const result: ITimeTable[] = [];
  for (const date in scheduleByDate) {
    const entries = scheduleByDate[date];
    //Sort by time_start
    entries.sort((a, b) => a.time_start.localeCompare(b.time_start));
    result.push({ date, timeTable: entries });
  }
  return result;
}

export function extendTimetable(timetable: ITimeTable[]): ITimeTable[] {
  const extendedTimetable: ITimeTable[] = [];
  const numWeeks = 16;
  const firstWeek = timetable.filter(item => item.timeTable[0].week === 1);
  const secondWeek = timetable.filter(item => item.timeTable[0].week === 2);

  // Function to add days to a date string (MM-DD-YY format)
  const addDays = (dateString: string, numDays: number): string => {
    const [month, day, year] = dateString.split('-').map(Number);
    const date = new Date(2000 + year, month - 1, day);
    date.setDate(date.getDate() + numDays);
    const newDay = String(date.getDate()).padStart(2, '0');
    const newMonth = String(date.getMonth() + 1).padStart(2, '0');
    const newYear = String(date.getFullYear()).padStart(2, '0');
    return `${newDay}.${newMonth}.${newYear}`;
  };

  for (let i = 0; i < numWeeks; i++) {
    const weekNum = i + 1;
    const isOddWeek = weekNum % 2 !== 0; //Check if it's an odd week
    const weekSchedule = isOddWeek ? firstWeek : secondWeek;

    for (const entry of weekSchedule) {
      const newDate = addDays(entry.date, (i - (i % 2)) * 7);
      const newTableDay:ITimeTable = {date: newDate, timeTable: [] as IScheduleItem[]};
      
      for (const slot of timeSlots) {
        const matchingEntry = entry.timeTable.find(item => item.time_start === slot.time_start);
        newTableDay.timeTable.push(matchingEntry || {
          auditory: "",
          date: newDate,
          group_names: [''],
          id: 0,
          name: "",
          teacher_names: [''],
          time_end: slot.time_end,
          time_start: slot.time_start,
          type: "",
          university: entry.timeTable[0].university,
          week: weekNum,
          weekday: entry.timeTable[0].weekday + i * 7,
        });
      }

      extendedTimetable.push(newTableDay);
    }
  }

  // Sort the extended timetable by date
  extendedTimetable.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime());
  return extendedTimetable;
}