import { type IScheduleItem, type ITimeTable } from "./constants";


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
      const newYear = String(date.getFullYear() % 100).padStart(2, '0');
      return `${newMonth}-${newDay}-${newYear}`;
  };

  for (let i = 0; i < numWeeks; i++) {
    const weekNum = i + 1;
    const isOddWeek = weekNum % 2 !== 0; //Check if it's an odd week
    const weekSchedule = isOddWeek ? firstWeek : secondWeek;

    for (const entry of weekSchedule) {
      const newDate = addDays(entry.date, (i - (i % 2)) * 7);
      const newTableDay:ITimeTable = {date: newDate, timeTable: [] as IScheduleItem[]};
      for(let item of entry.timeTable){
        newTableDay.timeTable.push({...item, date: newDate, weekday: item.weekday + i * 7});
      }
      extendedTimetable.push(newTableDay);
    }
  }

  // Sort the extended timetable by date
  extendedTimetable.sort((a, b) => new Date(a.date).getTime() - new Date(b.date).getTime());
  return extendedTimetable;
}