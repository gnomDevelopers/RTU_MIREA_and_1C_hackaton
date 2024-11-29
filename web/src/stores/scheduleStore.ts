import { defineStore } from "pinia";
import { WEEK_DAYS, type Day, type IScheduleItem, type ITimeTable, type TMaybeNumber } from "@/helpers/constants";
import { API_Schedule_Get_ClassName, API_Schedule_Get_GroupName, API_Schedule_Get_TeacherName, API_University_Groups_Schedule_Get, API_University_Teachers_Schedule_Get } from "@/api/api";
import { extendTimetable, transformSchedule } from "@/helpers/scheduleParser";


export const useScheduleStore = defineStore('schedule', {
  state() {
    return{
      scheduleType: 0 as number, // 0 - my schedule, 1 - general schedule
      scheduleTarget: 0 as number, // 0 - group, 1 - teacher, 2 - faculty

      selectedTarget: '' as string, // selected in search

      scheduleTableDay: [] as IScheduleItem[], // таблица расписания

      //for schedule page
      selectedSheduleTarget: null as TMaybeNumber,
      //for attendance page
      selectedClass: null as TMaybeNumber,

      scheduleData: [] as ITimeTable[], // расписание на весь семестр

      scheduleGroups: [] as string[], // группы для расписания
      scheduleTeachers: [] as string[], // преподаваатели для расписания
      scheduleFacultatives: [] as string[], // факультативы для расписания

      selectedDate: '',
      startDate: '02.09.2024',
    }
  },
  actions: {
    // подгружают списки доступных из расписания групп / учителей / факультативов
    async loadScheduleGroups(){
      try{
        const response: any = await API_University_Groups_Schedule_Get();
        this.scheduleGroups = [];
        for(let group of response.data){
          this.scheduleGroups.push(group.group);
        }
      } catch(error) {
        this.scheduleGroups = [];
      }
    },
    async loadScheduleTeachers(){
      try{
        const response: any = await API_University_Teachers_Schedule_Get();
        this.scheduleTeachers = [];
        for(let teacher of response.data){
          this.scheduleTeachers.push(teacher.teacher);
        }
      }catch(error){
        this.scheduleTeachers = [];
      }
    },
    async loadScheduleFacultatives(){
      try{
        const response: any = await API_University_Teachers_Schedule_Get();
        this.scheduleFacultatives = [];
        for(let facultative of response.data){
          this.scheduleFacultatives.push(facultative.name);
        }
      }catch(error){
        this.scheduleFacultatives = [];
      }
    },

    // подгружают расписание и конвертируют в scheduleData
    async loadScheduleTableByGroupName(groupName: string){
      try{
        const response:any = await API_Schedule_Get_GroupName(groupName);
        this.convertResponseToScheduleData(response);
      }catch(error){
        this.scheduleData = [];
      }
    },
    async loadScheduleTableByTeacherName(teacherName: string){
      try{
        const response:any = await API_Schedule_Get_TeacherName(teacherName);
        this.convertResponseToScheduleData(response);
      }catch(error){
        this.scheduleData = [];
      }
    },
    async loadScheduleTableByClassName(className: string){
      try{
        const response:any = await API_Schedule_Get_ClassName(className);
        this.convertResponseToScheduleData(response);
      }catch(error){
        this.scheduleData = [];
      }
    },


    selectScheduleDay(day: string){
      this.selectedDate = day;
      this.scheduleTableDay = [];
      for(let item of this.scheduleData){
        if(item.date === day){
          this.scheduleTableDay = item.timeTable;
          break;
        }
      }
    },

    convertResponseToScheduleData(response: any){
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
        // фильтруем, сортируем и растягиваем расписание на весь семестр
        this.scheduleData = extendTimetable(transformSchedule(res));
        
        console.log('scheduleData: ', this.scheduleData);
    }
  }
});