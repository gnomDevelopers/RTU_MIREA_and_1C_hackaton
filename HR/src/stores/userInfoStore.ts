import { defineStore } from "pinia";
import { API_Authenticate } from "@/api/api";
import { type TMaybeNumber, type TMaybeBoolean } from "@/helpers/constants";


export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      userID: null as TMaybeNumber,
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

      // "additional_experience": "string",
      // "cv_path": "string",
      // "father_name": "string",
      // "first_name": "string",
      // "gpa": 0,
      // "id": 0,
      // "last_name": "string",
      // "phone_number": "string",
      // "skills": [
      //   "string"
      // ],
      // "speciality": "string",
      // "telegram": "string",
      // "university": "string",
      // "useful_links": [
      //   "string"
      // ],
      // "work_experience": "string"
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
      this.userID = response.data.id;
      
      document.cookie = `access_token=${response.data.access_token}; max-age=${60 * 60 * 2}; secure; samesite=lax`;
      document.cookie = `refresh_token=${response.data.refresh_token}; max-age=${60 * 60 * 6}; secure; samesite=lax`;
    }
  }
});