import axios from "axios";
import { 
  API, DEVMODE, 
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
  type IAPI_University_Update 
} from '../helpers/constants';


//проверка аутентификации пользователя
export function API_Authenticate(){
  return new Promise((resolve, reject) => {
    axios.get(`${API}/login`)
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

