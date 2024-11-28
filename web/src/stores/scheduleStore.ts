import { defineStore } from "pinia";
import { type IScheduleItem, type ITimeTable, type TMaybeNumber } from "@/helpers/constants";
import { API_Schedule_Get_GroupName, API_University_Groups_Schedule_Get } from "@/api/api";
import { extendTimetable, transformSchedule } from "@/helpers/scheduleParser";


export const useScheduleStore = defineStore('schedule', {
  state() {
    return{
      scheduleType: 0 as number, // 0 - my schedule, 1 - general schedule
      scheduleTarget: 0 as number, // 0 - group, 1 - teacher, 2 - faculty

      selectedTarget: '' as string, // selected in search

      scheduleTableDay: [] as IScheduleItem[], // таблица расписания

      //for schedule page
      selectedSheduleGroup: null as TMaybeNumber,
      //for attendance page
      selectedClass: null as TMaybeNumber,

      scheduleData: [] as ITimeTable[], // расписание на весь семестр

      scheduleGroups: [] as string[], // группы для расписания

      selectedDate: '',
    }
  },
  actions: {
    async loadScheduleGroups(){
      if(this.selectedDate === '') this.selectedDate = this.getCurrentDate();
      try{
        const groups: any = await API_University_Groups_Schedule_Get();
        this.scheduleGroups = [];
        for(let group of groups.data){
          this.scheduleGroups.push(group.group);
        }
      } catch(error) {
        this.scheduleGroups = [];
      }
    },
    loadScheduleTableByGroupName(groupName: string){
      API_Schedule_Get_GroupName(groupName)
      .then((response: any) => {
        this.scheduleData = [];
        const res: IScheduleItem[] = [];

        for(let item of response.data){
          const data:IScheduleItem = {
            auditory: item.auditory,
            date: item.date,
            group_names: item.group_names,
            id: item.id,
            name: item.name,
            teacher_names: item.teacher_names,
            time_end: item.time_end,
            time_start: item.time_start,
            type: item.type,
            university: item.university,
            week: item.week,
            weekday: item.weekday,
          };
          res.push(data);
        }

        this.scheduleData = transformSchedule(res);
        console.log('scheduleData: ', this.scheduleData);
        this.scheduleData = extendTimetable(this.scheduleData);
        console.log('scheduleDataFull: ', this.scheduleData);
      })
      .catch(error => {
        this.scheduleData = [];
      });

    },
    loadScheduleTableByTeacherName(teacherName: string){

    },
    loadScheduleTableByClassName(className: string){

    },
    getCurrentDate() {
      const today = new Date();
      const day = String(today.getDate()).padStart(2, '0');
      const month = String(today.getMonth() + 1).padStart(2, '0'); // Month is 0-indexed
      const year = String(today.getFullYear() % 100).padStart(2, '0'); //2-digit year
    
      return `${day}.${month}.${year}`;
    }
  }
});