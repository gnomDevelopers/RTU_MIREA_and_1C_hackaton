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

//получение занятий по названию группы
export function API_Class_Get_GroupName(groupName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/class/group_name/${groupName}`)
    .then(response => {
      if(DEVMODE) console.log('Class get by groupName success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Class get by groupName error: ', error);
      reject(error);
    })
  });
};

//получение занятий по имени преподавателя
export function API_Class_Get_TeacherName(teacherName: string){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/class/teacher_name/${teacherName}`)
    .then(response => {
      if(DEVMODE) console.log('Class get by teacherName success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Class get by teacherName error: ', error);
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