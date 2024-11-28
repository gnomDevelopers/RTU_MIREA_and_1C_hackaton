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
