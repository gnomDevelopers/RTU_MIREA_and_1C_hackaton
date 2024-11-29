//all constants, interfaces, types, etc.

//contants

//api
export const API = '/api';
export const DEVMODE = true;
export const StatusWindowTime = 3000;

export const AUDITORY_TYPE_LIST = ['Лекторий', 'Практика', 'Спортивный зал', 'Лаборантская', 'Склад'];
export const AUDITORY_PROFILE_LIST = ['Обычная', 'Компьютерный класс', 'Химический класс', 'Физический класс', 'Биологический класс', 'Географический класс', 'Литературный класс', 'Исторический класс'];

//statusWindow codes
export enum StatusCodes {
  'error', 'info', 'loading', 'success'
};

export const MONTH_NAMES = ['Январь','Февраль','Март','Апрель','Май','Июнь','Июль','Август','Сентябрь','Октябрь','Ноябрь','Декабрь'];

export const WEEK_DAYS = ['Понедельник', 'Вторник', 'Среда', 'Четверг', 'Пятница', 'Суббота', 'Воскресенье'];

export const SCHEDULE_TARGET_TEXT = ['Выберите группу для просмотра', 'Выберите преподавателя для просмотра', 'Выберите факультет для просмотра'];

//types

export type TMaybeNumber = number | null;
export type TMaybeBoolean = boolean | null;
export type TMaybeString = string | null;

//interfaces

//validator.ts interface
export interface IValidAnswer{
  value: string,
  error: string,
}

//searchList item interface
export interface ISearchList{
  id: number,
  search_field: string,
  data: any,
}

//itemList item interface
export interface IItemList{
  id: number,
  name: string,
}

//groupCOrrect item interface
export interface IGroupCorrectUserItem{
  //
}

//statusWindow interface
export interface IStatusWindow{
  id: number,
  status: StatusCodes,
  text: string,
  time: number,
};

// calendar item interface
export interface Day{
  year: number,
  month: number,
  day: number,
};

// schedule get interface
export interface IScheduleItem {
  auditory: string,
  date: string,
  group_names: string[],
  id: number,
  name: string,
  teacher_names: string[],
  time_end: string,
  time_start: string,
  type: string,
  university: string,
  week: number,
  weekday: number,
}

// user interface
export interface IUser{
  department_id: number,
  educational_direction: string,
  faculty_id: number,
  father_name: string,
  first_name: string,
  group_id: number,
  last_name: string,
  role: number,
  university_id: number,
}

//userList item interface
export interface IUserGet extends IUser{
  id: number,
  // role: string,
}

//user create interface
export interface IUserCreate extends IUser{
  password: string,
}

export interface IGroup {
  id: number,
  name: string,
}

export interface ISubject {
  name: string,
}

export interface IGroupScores {
  user: IUserGet,
  scores: number[],
  avg: number,
  gpa: number,
}

export interface IGroupAttendance {
  user: IUserGet,
  attendace: number,
}

export interface ITimeTable{
  date: string,
  timeTable: IScheduleItem[],
}

export interface IUniversity{
  id: number,
  name: string,
  postfix: string,
}

//api interfaces

export interface IAPI_Login_Request{
  email: string,
  password: string,
};

export interface IAPI_Audience_Create{
  campus: string,
  capacity: number,
  name: string,
  profile: string,
  type: string
};

export interface IAPI_Audience_Update extends IAPI_Audience_Create{
  id: number,
};

export interface IAPI_Campus_Create {
  address: string,
  name: string,
  university: string,
};

export interface IAPI_Campus_Update extends IAPI_Campus_Create {
  id: number,
};

export interface IAPI_Class_Create{
  academic_discipline_id: number,
  auditory_id: number,
  date: string,
  group_names: string[],
  name: string,
  teacher_names: string[],
  time_end: string,
  time_start: string,
  type: string,
  week: number,
  weekday: number
};

export interface IAPI_Class_Update extends IAPI_Class_Create{
  id: number,
};

export interface IAPI_User_Schedule_Create{
  date: string,
  name: string,
  time_end: string,
  time_start: string
};

export interface IAPI_User_Schedule_Update extends IAPI_User_Schedule_Create{
  id: number,
};

export interface IAPI_University_Create{
  name: string,
};

export interface IAPI_University_Update extends IAPI_University_Create{
  id: number,
};

export interface IAPI_Class_Visit_QR{
  id: number,
  classID: number,
}

//rules

export const ROLES = [0, 1, 2, 3, 4, 5, 6];

export enum ROLES_NAME {
  'Администратор' = 0,
  'Проректор' = 1,
  'Декан' = 2,
  'Учебный отдел' = 3,
  'Заведующий кафедры' = 4,
  'Преподаватель' = 5,
  'Студент' = 6,
};

export const ROLES_SET_PRORECTOR = [2, 3, 4, 5, 6]; //  

export const ROLES_SET_DECAN_TO_PREPOD = [4, 5]; // 

export const ROLES_SET_ZAV_CAF = [2, 4]; // ?

export enum INFO_FIELDS {
  'faculty' = 0, 
  'department' = 1,
  'educationDirection' = 2,
};

export const INFO_FIELDS_ROLE = [
  [], // 0
  [], // 1
  ['faculty'], // 2
  ['faculty'], // 3
  ['faculty', 'department'], // 4
  ['faculty', 'department'], // 5
  ['faculty', 'department', 'educationDirection'], // 6
];

// functions

export function CHECK_COOKIE_EXIST(cookieName: string): boolean{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    if(cookie.split('=')[0].trim() === cookieName) return true;
  }
  return false;
}

export function GET_COOKIE(cookieName: string): string{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    let [key, value] = cookie.split('=');
    if(key.trim() === cookieName) return value;
  }
  return '';
}

export function GET_ROLEID_BY_ROLENAME(roleName: string): TMaybeNumber{
  for(let i = 0; i < ROLES.length; i++){
    if(ROLES_NAME[i] === roleName) return i;
  }
  return null;
}

export function GET_CORRECT_DATE(day: number, month: number, year: number) {
  const Day = String(day).padStart(2, '0');
  const Month = String(month).padStart(2, '0');
  const Year = String(year).padStart(2, '0');

  return `${Day}.${Month}.${Year}`;
}

export function GET_DAY_OF_WEEK(dateString: string): string {
  const [day, month, year] = dateString.split('.').map(Number);

  const date = new Date(year, month - 1, day);
  const dayIndex = date.getDay();

  const adjustedDayIndex = (dayIndex === 0) ? 6 : dayIndex -1 ;
  return WEEK_DAYS[adjustedDayIndex];
}

export function GET_WEEK_NUMBER(date: Date | string): number {
  if(typeof date === 'string') date = GET_DATE_FROM_STRING(date);
  // Copy date so don't modify original
  date = new Date(Date.UTC(date.getFullYear(), date.getMonth(), date.getDate()));
  // Get first day of year
  const yearStart = new Date(Date.UTC(date.getUTCFullYear(), 0, 1));
  // Calculate full weeks
  return Math.ceil((((date.getTime() - yearStart.getTime()) / 86400000) + yearStart.getDay() + 1) / 7);
}

export function GET_DATE_FROM_STRING(dateString: string): Date {
  const [day, month, year] = dateString.split('.').map(Number);
  return new Date(year, month - 1, day);
}

export function GET_WEEK_DIFFERENCE(startDate: Date, targetDate: Date): number {
  // Ensure both dates are at 00:00:00 hours to avoid time-related discrepancies
  startDate.setHours(0, 0, 0, 0);
  targetDate.setHours(0, 0, 0, 0);

  // Calculate the difference in milliseconds
  const diffInMilliseconds = targetDate.getTime() - startDate.getTime();

  // Calculate the difference in days
  const diffInDays = Math.floor(diffInMilliseconds / (1000 * 60 * 60 * 24));

  //Calculate the week number
  //This assumes weeks start on Monday and end on Sunday
  return Math.floor(diffInDays / 7) + 1;
}