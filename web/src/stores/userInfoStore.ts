import { defineStore } from "pinia";
import { API_Authenticate, API_UserInfo } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean } from "@/helpers/constants";


export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      userID: null as TMaybeNumber, // id пользователя
      first_name: '', // имя
      last_name: '', // фамилия
      father_name: '', //отчество
      university_id: null as TMaybeNumber, // id университет
      faculty_id: null as TMaybeNumber, // id факультета
      department_id: null as TMaybeNumber, // id кафедры
      educationalDirection: '', // навзание направления
      role: 1 as number, // роль
    }
  },
  actions: {
    async Authenticate(){
      try{
        const response:any = await API_Authenticate();
        if(response.data.authorized !== undefined) this.authorized = response.data.authorized;
        else this.authorized = false;
      }catch (error){
        this.authorized = false;
      }
    },
    loadUserData(){
      if(this.userID === null) return;
      API_UserInfo(this.userID)
      .then(response => {

      })
      .catch(error => {

      });
    }
  }
});