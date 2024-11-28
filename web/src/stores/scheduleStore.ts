import { defineStore } from "pinia";
import { type IScheduleItem, type ITimeTable, type TMaybeNumber } from "@/helpers/constants";
import { API_Schedule_Get_GroupName, API_University_Groups_Schedule_Get } from "@/api/api";
import { transformSchedule } from "@/helpers/scheduleParser";


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
    }
  },
  actions: {
    async loadScheduleGroups(){
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
          console.log('item: ', item);
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
        console.log('res: ', res);
        this.scheduleData = transformSchedule(res);

        console.log('scheduleData: ', this.scheduleData);
      })
      .catch(error => {
        this.scheduleData = [];
      });

    },
    loadScheduleTableByTeacherName(teacherName: string){

    },
    loadScheduleTableByClassName(className: string){

    },
  }
});