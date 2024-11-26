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
        this.authorized = response.data.authorized;
        this.userID = response.data.id;
      }catch (error){
        this.authorized = false;
        this.userID = -1;
      }
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
        this.first_name = 'Денис';
        this.last_name = 'Орлов',
        this.father_name = 'Сергеевич';
        this.university_id = 1;
        this.faculty_id = 1;
        this.department_id = 3;
        this.educationalDirection = 'Фуллстек разработка';
        this.role = 6;
        this.email = 'orlov_d_s';
      });
    }
  }
});