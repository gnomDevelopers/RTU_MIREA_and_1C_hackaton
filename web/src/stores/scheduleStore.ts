import { defineStore } from "pinia";
import { type IScheduleItem } from "@/helpers/constants";

export const useScheduleStore = defineStore('schedule', {
  state() {
    return{
      scheduleType: 0 as number, // 0 - my schedule, 1 - general schedule
      scheduleTarget: 0 as number, // 0 - group, 1 - teacher, 2 - faculty

      selectedTarget: '' as string, // selected in search

      scheduleTable: [] as IScheduleItem[], // таблица расписания

      ScheduleData: [] as IScheduleItem[][], // расписание на весь месяц
    }
  },
  actions: {
    loadScheduleTable(){
      //API
      const mainSchedule = [
        
      ];

      this.scheduleTable = [
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