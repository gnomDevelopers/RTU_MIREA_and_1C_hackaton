//all validation functions
import { type IValidAnswer } from "./constants";

//валидация логина пользователя
export function validUserLogin(value: string): IValidAnswer{
  if(value.length < 4) {
    return {value: '', error: 'Слишком короткий логин!'};
  }
  if(value.length > 20) {
    return {value: '', error: 'Слишком длинный логин!'};
  }
  if(value.match('/^[a-z_]+$/') === null) {
    return {value: '', error: 'Некорректный логин!'};
  }
  return {value: value, error: ''};
}

//валидация пароля пользователя
export function validUserPassword(value: string): IValidAnswer{
  if(value.length < 4) {
    return {value: '', error: 'Слишком короткий пароль!'};
  }
  if(value.match('/^[a-z_]+$/') === null) {
    return {value: '', error: 'Некорректный пароль!'};
  }
  return {value: value, error: ''};
}