import axios from "axios";
import { api, devMode } from '../helpers/constants';


//проверка аутентификации пользователя
export function API_Authenticate(){
  return new Promise((resolve, reject) => {
    axios.post(`${api}/auth`)
    .then(response => {
      if(devMode) console.log('Authentication success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(devMode) console.log('Authentication error: ', error);
      reject(error);
    })
  });
}