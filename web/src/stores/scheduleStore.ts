import { defineStore } from "pinia";
import { type IScheduleItem, StatusCodes } from "@/helpers/constants";
import { API_Schedule_Get_GroupName } from "@/api/api";
import { useStatusWindowStore } from "./statusWindowStore";

const statusWindow = useStatusWindowStore();

export const useScheduleStore = defineStore('schedule', {
  state() {
    return{
      scheduleType: 1 as number, // 0 - my schedule, 1 - general schedule
      scheduleTarget: 0 as number, // 0 - group, 1 - teacher, 2 - faculty

      selectedTarget: '' as string, // selected in search

      scheduleTableDay: [] as IScheduleItem[], // таблица расписания

      ScheduleData: [] as IScheduleItem[][], // расписание на весь семестр
    }
  },
  actions: {
    loadScheduleTable(){
      let stID = statusWindow.showStatusWindow(StatusCodes.loading, 'Получаем данные расписания...', -1);
      API_Schedule_Get_GroupName("ЭФБО-01-23")
      .then(response => {
        //
      })
      .catch(error => {
        statusWindow.showStatusWindow(StatusCodes.error, 'Неудалось получить данные расписания!');
      })
      .finally(() => {
        statusWindow.deteleStatusWindow(stID);

        const res = [
          { academic_discipline_id: 0, auditory_id: 0, date: '', group_names: ['ЭФБО-01-23'], id: 0, name: 'Философия', teacher_names: ['Teacher Teacherovich'], time_end: '9:00', time_start: '10:30', type: 'ПР', week: 0, weekday: 0 },
        ];
      });


      this.scheduleTableDay = [
        { time: '9.00-10:30',   type: '',   title: '', place: '', groups: [''] },
        { time: '10.40-12.10',  type: 'ПР', title: 'Создание программного обеспечения', place: 'В-407 (В-78)', groups: ['ЭФБО-01-23'] },
        { time: '12.40-14.10',  type: '',   title: '', place: '', groups: [''] },
        { time: '14.20-15.50',  type: 'ПР', title: 'История России', place: 'A-315 (МП-1)', groups: ['ЭФБО-01-23'] },
        { time: '16.20-17.50',  type: 'ЛК', title: 'Программирование корпоративных систем', place: 'A-16 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
        { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
      ];
    },
  }
});