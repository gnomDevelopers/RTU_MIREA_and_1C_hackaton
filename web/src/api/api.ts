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