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

export interface IAPI_Login_Request{
  login: string,
  password: string,
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