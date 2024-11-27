import { defineStore } from "pinia";
import { API_Authenticate, API_UserInfo } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean, type TMaybeString } from "@/helpers/constants";
import { useUniversityStore } from "./universityStore";

// const universityStore = useUniversityStore();

export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      userID: null as TMaybeNumber, // id пользователя
      first_name: '', // имя
      last_name: '', // фамилия
      father_name: '', //отчество
      university: null as TMaybeString, // id университет
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
        const response = await API_Authenticate();
        this.onAuthorized(response);

      }catch (error){
        this.authorized = false;
        this.userID = null;
      }
    },
    async loadUserData(){
      if(this.userID === null) return; // пользователь не авторизован
      try{
        const response = await API_UserInfo(this.userID);

        // this.first_name = 'Денис';
        // this.last_name = 'Орлов',
        // this.father_name = 'Сергеевич';
        // this.university_id = 1;
        // this.faculty_id = 1;
        // this.department_id = 3;
        // this.educationalDirection = 'Фуллстек разработка';
        // this.role = 6;
        // this.email = 'orlov_d_s';
      }catch(error){

      }
    },
    async onAuthorized(response: any){
      const universityStore = useUniversityStore();
      this.authorized = true;
      this.userID = response.data.id;
      
      document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=strict`;
      document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=strict`;
      
      if(this.authorized){
        await this.loadUserData(); // загрузка данных о пользователе
        await universityStore.loadUniversityInfo(); // загрузка данных об институте
      }
    }
  }
});