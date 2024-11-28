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
        this.scheduleData = transformSchedule(response.date);

        console.log('scheduleData: ', this.scheduleData);
      })
      .catch(error => {
        this.scheduleData = [];
      });

      // this.scheduleData = [
      //   {
      //     date: '25.11.2024', 
      //     timeTable: [
      //       { time: '9.00-10:30',   id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: 'ПР', title: 'Иностранный язык', place: 'И-320 (В-78)', groups: ['ЭФБО-01-23'] },
      //       { time: '10.40-12.10',  id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: 'ЛК', title: 'Технологии индустриального программирования', place: 'А-10 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
      //       { time: '12.40-14.10',  id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: 'ЛК', title: 'Создание программного обеспечения', place: 'А-15 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
      //       { time: '14.20-15.50',  id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: 'ПР', title: 'Физическая культура и спорт', place: 'ФОК-9 (В-78)', groups: ['ЭФБО-01-23'] },
      //       { time: '16.20-17.50',  id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: '',   title: '', place: '', groups: [''] },
      //       { time: '18.00-19.30',  id: 1, week: 1, weekDay: 1, teacherNames: ['Сафронов'], type: '',   title: '', place: '', groups: [''] },
      //     ]
      //   },
      // ];
      // this.scheduleTableDay = this.scheduleData[0].timeTable;

    },
    loadScheduleTableByTeacherName(teacherName: string){
      // let stID = statusWindow.showStatusWindow(StatusCodes.loading, 'Получаем данные расписания...', -1);
      // API_Schedule_Get_TeacherName(teacherName)
      // .then(response => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.success, 'Расписание преподавателя получено!');
      // })
      // .catch(error => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.error, 'Неудалось получить данные расписания!');
      // })
    },
    loadScheduleTableByClassName(className: string){
      // let stID = statusWindow.showStatusWindow(StatusCodes.loading, 'Получаем данные расписания...', -1);
      // API_Schedule_Get_ClassName(className)
      // .then(response => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.success, 'Расписание предмета получено!');
      // })
      // .catch(error => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.error, 'Неудалось получить данные расписания!');
      // })
    },
    loadGroups(){
      // let stID = statusWindow.showStatusWindow(StatusCodes.loading, 'Получаем данные о группах...', -1);
      // API_Schedule_Get_ClassName(className)
      // .then(response => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.success, 'Расписание предмета получено!');
      // })
      // .catch(error => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.error, 'Неудалось получить данные расписания!');
      // })
    }
  }
});