import { defineStore } from "pinia";
import { API_Authenticate } from "@/api/api";
import { type TMaybeNumber } from "@/helpers/constants";

export const useUserInfoStore = defineStore('userInfo', {
  state() {
    return{
      authenticated: false,

      first_name: '', // имя
      last_name: '', // фамилия
      father_name: '', //отчество
      university_id: null as TMaybeNumber, // id университета
      permission_id: null as TMaybeNumber, // id уровня доступа
      faculty_id: null as TMaybeNumber, // id факультета
      department_id: null as TMaybeNumber, // id 
    }
  },
  actions: {
    async Authenticate(){
      const data = await API_Authenticate();

      //...
    }
  }
});