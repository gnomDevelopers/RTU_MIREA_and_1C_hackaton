import axios from "axios";
import { API, DEVMODE, type IAPI_Login_Request } from '../helpers/constants';


//проверка аутентификации пользователя
export function API_Authenticate(){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/auth`)
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
interface IAPI_Audience_Create{
  campus_id: number,
  capacity: number,
  name: string,
  profile: string,
  type: string
};
export function API_Audience_Create(data: IAPI_Audience_Create){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/audience`, data)
    .then(response => {
      if(DEVMODE) console.log('Audience create success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Audience create error: ', error);
      reject(error);
    })
  });
};

//изменение аудитории
interface IAPI_Audience_Update extends IAPI_Audience_Create{
  id: number,
};
export function API_Audience_Update(data: IAPI_Audience_Update){
  return new Promise((resolve, reject) => {
    axios.put(`${API}/audience`, data)
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
    axios.delete(`${API}/audience/${audienceID}`)
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

/////// campus ///////

//создание кампуса
interface IAPI_Campus_Create {
  address: string,
  name: string,
  university_id: number
};
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
interface IAPI_Campus_Update extends IAPI_Campus_Create {
  id: number,
};
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
export function API_Campus_Get_University(universityID: number){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/campus/university_id/${universityID}`)
    .then(response => {
      if(DEVMODE) console.log('Campus get by university success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Campus get by university error: ', error);
      reject(error);
    })
  });
};

/////// classes ///////

interface IAPI_Class_Create{
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

interface IAPI_Class_Update extends IAPI_Class_Create{
  id: number,
}
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
  if (DEVMODE) for(let item of data) { console.log('File with schedule: ', item[1]); break; }
  
  return new Promise((resolve, reject) => {
    axios.post(`${API}/schedule/parse`, data, {
      headers: {
        'Content-Type': 'multipart/form-data'
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

/////// user schedule ///////

interface IAPI_User_Schedule_Create{
  date: string,
  name: string,
  time_end: string,
  time_start: string
}
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

interface IAPI_User_Schedule_Update extends IAPI_User_Schedule_Create{
  id: number,
}
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

interface IAPI_University_Create{
  name: string,
}
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

interface IAPI_University_Update extends IAPI_University_Create{
  id: number,
}
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
    axios.get(`${API}/university/all`)
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