import { type IScheduleItem, type ITimeTable } from "./constants";


export function transformSchedule(schedule: IScheduleItem[]): ITimeTable[] {
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


// Example usage:
const schedule: IScheduleItem[] = [
  {
    auditory: "ауд. А-15 (В-78)",
    date: "09-02-24",
    group_names: ['ЭФБО-01-23', 'ЭФБО-02-23'],
    id: 311,
    name: "Физика",
    teacher_names: ['Сафронов А.А.'],
    time_end: "10-30",
    time_start: "09-00",
    type: "ЛК",
    university: "РТУ МИРЭА",
    week: 1,
    weekday: 1,
  },
  {
    auditory: "ауд. Б-20",
    date: "10-02-24",
    group_names: ['ЭФБО-03-23'],
    id: 313,
    name: "Химия",
    teacher_names: ['Петров П.П.'],
    time_end: "12-00",
    time_start: "11-00",
    type: "ЛК",
    university: "РТУ МИРЭА",
    week: 1,
    weekday: 1,
  },
    {
    auditory: "ауд. А-15 (В-78)",
    date: "09-02-24",
    group_names: ['ЭФБО-01-23', 'ЭФБО-02-23'],
    id: 312,
    name: "Математика",
    teacher_names: ['Иванов И.И.'],
    time_end: "11-30",
    time_start: "10-30",
    type: "ЛК",
    university: "РТУ МИРЭА",
    week: 1,
    weekday: 1,
  },
];

const transformedSchedule = transformSchedule(schedule);
console.log(JSON.stringify(transformedSchedule, null, 2));