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

export const SCHEDULE_TARGET_TEXT = ['Выберите группу для просмотра', 'Выберите преподавателя для просмотра', 'Выберите факультет для просмотра'];

//types

export type TMaybeNumber = number | null;
export type TMaybeBoolean = boolean | null;

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

// schedule item interface
export interface IScheduleItem{
  time: string,
  type: string,
  title: string,
  place: string,
  groups: string[],
};

// user interface
export interface IUser{
  name: string,
  surname: string,
  thirdname: string,
  role: number,
  faculty_id: number,
  department_id: number,
  educational_direction: string,
}

//userList item interface
export interface IUserGet extends IUser{
  id: number,
}


//api interfaces

export interface IAPI_Login_Request{
  email: string,
  password: string,
};

export interface IAPI_Audience_Create{
  campus_id: number,
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
  university_id: number
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
