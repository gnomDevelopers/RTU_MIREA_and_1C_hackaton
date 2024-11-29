import { defineStore } from "pinia";
import { API_Authenticate } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean } from "@/helpers/constants";


export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authorized: null as TMaybeBoolean, // проверка авторизованности
      first_name: '', // имя
      last_name: '', // фамилия
      father_name: '', //отчество
      university_id: null as TMaybeNumber, // id университета
      permission_id: null as TMaybeNumber, // id уровня доступа
      faculty_id: null as TMaybeNumber, // id факультета
      department_id: null as TMaybeNumber, // id 
      //role: -1 as number,
      role: 1 as number,
    }
  },
  actions: {
    async Authenticate(){
      try{
        const response:any = await API_Authenticate();
        await this.onAuthorized(response);
        if(response.data.authorized !== undefined) this.authorized = response.data.authorized;
        else this.authorized = false;
      }catch (error){
        this.authorized = false;
      }
    },
    async onAuthorized(response: any){
      this.authorized = true;

      console.log('Access Token:', response.data.access_token);
      console.log('Refresh Token:', response.data.refresh_token);
      
      document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=lax`;
      document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=lax`;
    }
  }
});