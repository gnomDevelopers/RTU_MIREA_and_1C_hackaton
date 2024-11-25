import { defineStore } from "pinia";
import { type IScheduleItem, type ITimeTable, StatusCodes, type TMaybeNumber } from "@/helpers/constants";
import { API_Schedule_Get_ClassName, API_Schedule_Get_GroupName, API_Schedule_Get_TeacherName } from "@/api/api";
import { useStatusWindowStore } from "./statusWindowStore";

const statusWindow = useStatusWindowStore();

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
    }
  },
  actions: {
    loadScheduleTableByGroupName(groupName: string){
      // let stID = statusWindow.showStatusWindow(StatusCodes.loading, 'Получаем данные расписания...', -1);
      // API_Schedule_Get_GroupName(groupName)
      // .then(response => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.success, 'Расписание группы получено!');
      // })
      // .catch(error => {
      //   statusWindow.deteleStatusWindow(stID);
      //   statusWindow.showStatusWindow(StatusCodes.error, 'Неудалось получить данные расписания!');
      // });
      this.scheduleTableDay = [
        { time: '9.00-10:30',   type: 'ПР', title: 'Иностранный язык', place: 'И-320 (В-78)', groups: ['ЭФБО-01-23'] },
        { time: '10.40-12.10',  type: 'ЛК', title: 'Технологии индустриального программирования', place: 'А-10 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
        { time: '12.40-14.10',  type: 'ЛК', title: 'Создание программного обеспечения', place: 'А-15 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
        { time: '14.20-15.50',  type: 'ПР', title: 'Физическая культура и спорт', place: 'ФОК-9 (В-78)', groups: ['ЭФБО-01-23'] },
        { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
        { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
      ];
      this.scheduleData = [
        {
          date: '25.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: 'ПР', title: 'Иностранный язык', place: 'И-320 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '10.40-12.10',  type: 'ЛК', title: 'Технологии индустриального программирования', place: 'А-10 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
            { time: '12.40-14.10',  type: 'ЛК', title: 'Создание программного обеспечения', place: 'А-15 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
            { time: '14.20-15.50',  type: 'ПР', title: 'Физическая культура и спорт', place: 'ФОК-9 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
        {
          date: '26.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: '',   title: '', place: '', groups: [''] },
            { time: '10.40-12.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '12.40-14.10',  type: 'ПР', title: 'Системы искусственного интеллекта', place: 'Б-412 (МП-1)', groups: ['ЭФБО-01-23'] },
            { time: '14.20-15.50',  type: 'ПР', title: 'Системы искусственного интеллекта', place: 'Б-412 (МП-1)', groups: ['ЭФБО-01-23'] },
            { time: '16.20-17.50',  type: 'ЛК', title: 'Социальная психология и педагогика', place: 'A-61 (МП-1)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
            { time: '18.00-19.30',  type: 'ЛК', title: 'Системы искусственного интеллекта', place: 'А-61 (МП-1)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
          ]
        },
        {
          date: '27.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: 'ПР', title: 'Создание программного обеспечения', place: 'В-408 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '10.40-12.10',  type: 'ПР', title: 'Создание программного обеспечения', place: 'В-408 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '12.40-14.10',  type: 'ПР', title: 'Технологии индустриального программирования', place: 'В-317 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '14.20-15.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
        {
          date: '28.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: 'ПР', title: 'Математический анализ', place: 'А-403 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '10.40-12.10',  type: 'ПР', title: 'Программирование компоративных систем', place: 'В-401-2 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '12.40-14.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '14.20-15.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
        {
          date: '29.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: 'ПР', title: 'Фронтенд и бэкенд разработка', place: 'В-415 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '10.40-12.10',  type: 'ПР', title: 'Дискретная математика', place: 'В-218 (В-78)', groups: ['ЭФБО-01-23'] },
            { time: '12.40-14.10',  type: 'ЛК', title: 'Дискретная математика', place: 'А-18 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
            { time: '14.20-15.50',  type: 'ЛК', title: 'Математический анализ', place: 'А-18 (В-78)', groups: ['ЭФБО-01-23', 'ЭФБО-02-23', 'ЭФБО-03-23', 'ЭФБО-04-23', 'ЭФБО-05-23'] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
        {
          date: '30.11.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: '',   title: '', place: '', groups: [''] },
            { time: '10.40-12.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '12.40-14.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '14.20-15.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
        {
          date: '1.12.2024', 
          timeTable: [
            { time: '9.00-10:30',   type: '',   title: '', place: '', groups: [''] },
            { time: '10.40-12.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '12.40-14.10',  type: '',   title: '', place: '', groups: [''] },
            { time: '14.20-15.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '16.20-17.50',  type: '',   title: '', place: '', groups: [''] },
            { time: '18.00-19.30',  type: '',   title: '', place: '', groups: [''] },
          ]
        },
      ];
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