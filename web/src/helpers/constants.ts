//all constants, interfaces, types, etc.

//contants

//api
export const API = 'https://...';
export const DEVMODE = true;
export const StatusWindowTime = 3000;

//statusWindow codes
export enum StatusCodes {
  'error', 'info', 'loading', 'success'
};

export const SCHEDULE_TARGET_TEXT = ['Выберите группу для просмотра', 'Выберите проподавателя для просмотра', 'Выберите факультет для просмотра'];

//types

export type TMaybeNumber = number | null;

//interfaces

//validator.ts interface
export interface IValidAnswer{
  value: string,
  error: string,
}

//userList item interface
export interface IUsersList{
  id: number,
  name: string,
  role: number,
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
  isThisMonth: boolean,
};

// schedule item interface
export interface IScheduleItem{
  time: string,
  type: string,
  title: string,
  place: string,
  groups: string[],
};

//api interfaces

export interface IAPI_Login_Request{
  login: string,
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
  'Заведующий кафедры' = 3,
  'Преподаватель' = 4,
  'Староста группы' = 5,
  'Студент' = 6,
};

export const ROLES_SET_PRORECTOR = [2, 4];

export const ROLES_SET_DECAN_TO_PREPOD = [3, 4];
export const ROLES_SET_DECAN_TO_STUDENT = [5, 6];

export const ROLES_SET_ZAV_CAF = [2, 4];