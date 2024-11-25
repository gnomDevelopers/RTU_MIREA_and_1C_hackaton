import { defineStore } from "pinia";
import { API_Authenticate, API_UserInfo } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean, type TMaybeString } from "@/helpers/constants";


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
      educationalDirection: null as TMaybeString, // название направления
      role: 1 as number, // роль
      email: '',
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

      this.authorized = true;
      this.userID = 1;
    },
    loadUserData(){
      // if(this.userID === null) return;
      if(this.userID === null) this.userID = 1;
      API_UserInfo(this.userID)
      .then(response => {

      })
      .catch(error => {

      })
      .finally(() => {
        this.first_name = 'Даниил';
        this.last_name = 'Постников',
        this.father_name = 'Сергеевич';
        this.university_id = 1;
        this.faculty_id = null;
        this.department_id = null;
        this.educationalDirection = null;
        this.role = 1;
        this.email = 'postnikov_d_s';
      });
    }
  }
});