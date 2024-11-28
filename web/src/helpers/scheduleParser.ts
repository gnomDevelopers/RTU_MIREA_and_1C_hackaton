import { type IScheduleItem, type ITimeTable } from "./constants";


export function transformSchedule(schedule: IScheduleItem[]): ITimeTable[] {
  console.log('start transform');
  // Group entries by date
  const scheduleByDate = schedule.reduce((acc, entry) => {
    const date = entry.date;
    acc[date] = acc[date] || [];
    acc[date].push(entry);
    return acc;
  }, {} as { [date: string]: IScheduleItem[] });

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
