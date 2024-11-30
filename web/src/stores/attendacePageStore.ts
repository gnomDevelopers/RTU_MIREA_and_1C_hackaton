import { defineStore } from "pinia";
import{ GET_CORRECT_DATE, type IGroupAttendance, type TMaybeNumber, type TMaybeString } from "@/helpers/constants";
import { API_University_Groups_Schedule_Get } from "@/api/api";

export const useAttendacePageStore = defineStore('attendancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,

      attendanceGroupMembers: [] as IGroupAttendance[],

      attendanceGroups: [] as string[],

      selectedDate: GET_CORRECT_DATE(new Date().getDate(), new Date().getMonth() + 1, new Date().getFullYear()),
      startDate: '02.09.2024',
    }
  },
  actions: {
    async loadScheduleGroups(){
      try{
        const response: any = await API_University_Groups_Schedule_Get();
        this.attendanceGroups = [];
        for(let group of response.data){
          this.attendanceGroups.push(group.group);
        }
      } catch(error) {
        this.attendanceGroups = [];
      }
    },

  }
});