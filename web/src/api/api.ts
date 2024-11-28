import axios from "axios";
import { 
  API, DEVMODE, GET_COOKIE,
  type IAPI_Login_Request, 
  type IAPI_Audience_Create, 
  type IAPI_Audience_Update, 
  type IAPI_Campus_Create,
  type IAPI_Campus_Update, 
  type IAPI_Class_Create, 
  type IAPI_Class_Update, 
  type IAPI_User_Schedule_Create, 
  type IAPI_User_Schedule_Update, 
  type IAPI_University_Create, 
  type IAPI_University_Update, 
  type IUser,
  type IUserCreate
} from '../helpers/constants';


//проверка аутентификации пользователя
export function API_Authenticate(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/login`,  {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
     })
    .then(response => {
      if(DEVMODE) console.log('Authentication success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Authentication error: ', error);
      reject(error);
    })
  });
};

//вход в аккаунт
export function API_Login(data: IAPI_Login_Request){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/login`, data)
    .then(response => {
      if(DEVMODE) console.log('Login post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Login post error: ', error);
      reject(error);
    })
  });
};

//выход из аккаунта
export function API_Logout(){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/logout`)
    .then(response => {
      if(DEVMODE) console.log('Logout post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Logout post error: ', error);
      reject(error);
    })
  });
};

//получение данных об аккаунте
export function API_UserInfo(userID: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/user/${userID}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('UserInfo get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('UserInfo get error: ', error);
      reject(error);
    })
  });
};

//тест отправки файла на сервер
export function API_SendFile(data: FormData){
  if (DEVMODE) for(let item of data) console.log('API_SendFile: ', item[1]);

  return new Promise((resolve, reject) => {
    axios.post(`${API}/file`, data, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    .then(response => {
      if(DEVMODE) console.log('FileSend success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('FilSend error: ', error);
      reject(error);
    })
  });
};

/////// AUDIENCE ///////

//создание аудитории
export function API_Audience_Create(data: IAPI_Audience_Create[]){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/audience`, data, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Audiences create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Audiences create error: ', error);
      reject(error);
    })
  });
};

//изменение аудитории
export function API_Audience_Update(data: IAPI_Audience_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/auth/audience`, data, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Audience update success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Audience update error: ', error);
      reject(error);
    })
  });
};

//удаление аудитории
export function API_Audience_Delete(audienceID: number){
  return new Promise((resolve, reject) => {
    axios.delete(`${API}/auth/audience/${audienceID}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Audience delete success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Audience delete error: ', error);
      reject(error);
    })
  });
};

//получение всех аудиторий по универу
export function API_Audience_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/audience/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Audience get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Audience get error: ', error);
      reject(error);
    })
  });
};

/////// CAMPUS ///////

//создание кампуса
export function API_Campus_Create(data: IAPI_Campus_Create){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/campus`, data)
    .then(response => {
      if(DEVMODE) console.log('Campus create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Campus create error: ', error);
      reject(error);
    })
  });
};

//изменение кампуса
export function API_Campus_Update(data: IAPI_Campus_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/campus`, data)
    .then(response => {
      if(DEVMODE) console.log('Campus update success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Campus update error: ', error);
      reject(error);
    })
  });
};

//удаление кампуса
export function API_Campus_Delete(campusID: number){
  return new Promise((resolve, reject) => {
    axios.delete(`${API}/campus/${campusID}`)
    .then(response => {
      if(DEVMODE) console.log('Campus delete success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Campus delete error: ', error);
      reject(error);
    })
  });
};

//получение всех кампусов
export function API_Campus_Get_University(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/campus/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Campus get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Campus get error: ', error);
      reject(error);
    })
  });
};

/////// university users ///////

//создание пользователей для вуза
export function API_University_Users_Create(data: IUserCreate[]){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/user`, data, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Users create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Users create  error: ', error);
      reject(error);
    })
  });
};

//получение всех пользователей вуза
export function API_University_Users_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/user/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Users get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Users get  error: ', error);
      reject(error);
    })
  });
};

//получение всех групп вуза
export function API_University_Groups_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/groups/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Groups get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Groups get error: ', error);
      reject(error);
    })
  });
};

/////// classes ///////

//создание занятия
export function API_Class_Create(data: IAPI_Class_Create[]){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/class`, data)
    .then(response => {
      if(DEVMODE) console.log('Class create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Class create error: ', error);
      reject(error);
    })
  });
};

//обновление занятия
export function API_Class_Update(data: IAPI_Class_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/class`, data)
    .then(response => {
      if(DEVMODE) console.log('Class update success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Class update error: ', error);
      reject(error);
    })
  });
};

//удаление занятия
export function API_Class_Delete(classID: number){
  return new Promise((resolve, reject) => {
    axios.delete(`${API}/class/${classID}`)
    .then(response => {
      if(DEVMODE) console.log('Class delete success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Class delete error: ', error);
      reject(error);
    })
  });
};

/////// schedule ///////

//создание расписания из файла
export function API_Schedule_Create(data: FormData){
  if (DEVMODE) for(let item of data) { console.log('File schedule: ', item[1]); break; }
  
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/schedule/parse`, data, {
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Schedule create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Schedule create error: ', error);
      reject(error);
    })
  });
};

//получение расписания по названию группы
export function API_Schedule_Get_GroupName(groupName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/schedule/group/${groupName}`)
    .then(response => {
      if(DEVMODE) console.log('Schedule get by groupName success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Schedule get by groupName error: ', error);
      reject(error);
    })
  });
};

//получение расписания по имени преподавателя
export function API_Schedule_Get_TeacherName(teacherName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/schedule/teacher/${teacherName}`)
    .then(response => {
      if(DEVMODE) console.log('Schedule get by teacherName success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Schedule get by teacherName error: ', error);
      reject(error);
    })
  });
};

//получение расписания по названию предмета
export function API_Schedule_Get_ClassName(className: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/schedule/name/${className}`)
    .then(response => {
      if(DEVMODE) console.log('Schedule get by className success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Schedule get by className error: ', error);
      reject(error);
    })
  });
};

//получение факультативных занятий
export function API_Facultatives_Get(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/schedule/search/name`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Facultatives get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Facultatives get error: ', error);
      reject(error);
    })
  });
};

/////// user schedule ///////

//создание личного расписания
export function API_User_Schedule_Create(data: IAPI_User_Schedule_Create){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth/user_schedule`, data)
    .then(response => {
      if(DEVMODE) console.log('User schedule create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('User schedule create error: ', error);
      reject(error);
    })
  });
};

//обновление личного расписания
export function API_User_Schedule_Update(data: IAPI_User_Schedule_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/auth/user_schedule`, data)
    .then(response => {
      if(DEVMODE) console.log('User schedule update success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('User schedule update error: ', error);
      reject(error);
    })
  });
};

//удаление личного расписания
export function API_User_Schedule_Delete(user_schedule_ID: number){
  return new Promise((resolve, reject) => {
    axios.delete(`${API}/auth/user_schedule/${user_schedule_ID}`)
    .then(response => {
      if(DEVMODE) console.log('User schedule delete success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('User schedule delete error: ', error);
      reject(error);
    })
  });
};

//получение личного расписания
export function API_User_Schedule_Get(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/user_schedule`)
    .then(response => {
      if(DEVMODE) console.log('User schedule get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('User schedule get error: ', error);
      reject(error);
    })
  });
};

/////// university ///////

//создание университета
export function API_University_Create(data: IAPI_University_Create){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/university`, data)
    .then(response => {
      if(DEVMODE) console.log('University create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('University create error: ', error);
      reject(error);
    })
  });
};

//обновление университета
export function API_University_Update(data: IAPI_University_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/university`, data)
    .then(response => {
      if(DEVMODE) console.log('University update success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('University update error: ', error);
      reject(error);
    })
  });
};

//удаление университета
export function API_University_Delete(universityID: number){
  return new Promise((resolve, reject) => {
    axios.delete(`${API}/university/${universityID}`)
    .then(response => {
      if(DEVMODE) console.log('University delete success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('University delete error: ', error);
      reject(error);
    })
  });
};

//получение всех университетов
export function API_University_Get_All(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/university/all`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('University get all success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('University get all error: ', error);
      reject(error);
    })
  });
};

//получение университета по названию
export function API_University_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/university/${universityName}`)
    .then(response => {
      if(DEVMODE) console.log('University get by universityName success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('University get by universityName error: ', error);
      reject(error);
    })
  });
};

/////// education ///////

//получение всех кафедр
export function API_Departments_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/department/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Departments get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Departments get error: ', error);
      reject(error);
    })
  });
};

//получение всех факультетов
export function API_Faculties_Get(universityName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/auth/faculties/university/${universityName}`, {
      headers: {
        Authorization: 'Bearer ' + GET_COOKIE('access_token'),
      }
    })
    .then(response => {
      if(DEVMODE) console.log('Faculties get success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Faculties get error: ', error);
      reject(error);
    })
  });
};