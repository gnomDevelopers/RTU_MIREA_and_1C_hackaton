//all constants, interfaces, types, etc.

//contants

//api
//export const API = 'http://localhost:8080';
export const API = 'https://gnomedeployer.ru/api';
export const DEVMODE = true;
export const StatusWindowTime = 3000;

//statusWindow codes
export enum StatusCodes {
  'error', 'info', 'loading', 'success'
};

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

// work user interfaces
export interface WorkUser {
  additional_experience: string;
  cv_path: string;
  father_name: string;
  first_name: string;
  gpa: number;
  id: number;
  last_name: string;
  phone_number: string;
  skills: string[];
  speciality: string;
  telegram: string;
  university: string;
  useful_links: string[];
  work_experience: string;
}

//api interfaces

export interface IAPI_Login_Request{
  email: string,
  password: string,
};

// functions

export function GET_COOKIE(cookieName: string): string{
  const allCookie = document.cookie.split(';');
  for(let cookie of allCookie){
    let [key, value] = cookie.split('=');
    if(key === cookieName) return value;
  }
  return '';
}


