import axios from "axios";
import { 
  API, DEVMODE, GET_COOKIE,
  type IAPI_Login_Request,
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

//вход в аккаунт студента
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

//вход в аккаунт ХР
export function API_LoginHR(data: IAPI_Login_Request){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/login/hr`, data)
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


