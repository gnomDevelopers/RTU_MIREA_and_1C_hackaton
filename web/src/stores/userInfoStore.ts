import { defineStore } from "pinia";
import { API_Authenticate, API_UserInfo } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean, type TMaybeString, ROLES_NAME, ROLES, GET_ROLEID_BY_ROLENAME, CHECK_COOKIE_EXIST, GET_COOKIE } from "@/helpers/constants";
import { useUniversityStore } from "./universityStore";

// const universityStore = useUniversityStore();

export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      userID: null as TMaybeNumber, // id пользователя

      email: '',
      first_name: '', // имя
      last_name: '', // фамилия
      father_name: '', //отчество
      role: null as TMaybeNumber, // роль
      group_id: null as TMaybeNumber,
      university: null as TMaybeString, // id университет
      faculty_id: null as TMaybeNumber, // id факультета
      department_id: null as TMaybeNumber, // id кафедры
      educationalDirection: null as TMaybeString, // название направления
      isPasswordChanged: true, // сменен ли пароль

      group_name: '',

      
    }
  },
  actions: {
    async Authenticate(){
      try{
        const response = await API_Authenticate();
        await this.onAuthorized(response);

      }catch (error){
        this.authorized = false;
        this.userID = null;
      }
    },
    async loadUserData(){
      if(this.userID === null) return; // пользователь не авторизован
      try{
        const response:any = await API_UserInfo(this.userID);

        this.first_name = response.data.first_name;
        this.last_name = response.data.last_name,
        this.father_name = response.data.father_name;
        this.university = response.data.university_name;
        this.faculty_id = response.data.faculty_id;
        this.department_id = response.data.department_id;
        this.educationalDirection = response.data.educational_direction;
        this.role = GET_ROLEID_BY_ROLENAME(response.data.role);
        this.email = response.data.email;
        this.group_id = response.data.group_id;
        this.isPasswordChanged = response.data.is_password_changed;
      }catch(error){
        this.email = '';
        this.first_name = '';
        this.last_name = '';
        this.father_name = '';
        this.role = null;
        this.group_id = null;
        this.university = null;
        this.faculty_id = null;
        this.department_id = null;
        this.educationalDirection = null;
        this.isPasswordChanged = true;
      }
    },
    async onAuthorized(response: any){
      const universityStore = useUniversityStore();
      this.authorized = true;
      this.userID = response.data.id;
      
      if(!CHECK_COOKIE_EXIST('acceptCookie') || GET_COOKIE('acceptCookie') === 'true'){
        document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=strict`;
        document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=strict`;
      }
      
      if(this.authorized){
        await this.loadUserData(); // загрузка данных о пользователе
        await universityStore.loadUniversityInfo(); // загрузка данных об институте
      }
    },
  }
});