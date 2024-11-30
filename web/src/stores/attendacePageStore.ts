import { defineStore } from "pinia";
import{ GET_CORRECT_DATE, type IGroupAttendance, type TMaybeNumber, type TMaybeString } from "@/helpers/constants";
import { API_University_Groups_Schedule_Get } from "@/api/api";

export const useAttendacePageStore = defineStore('attendancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,

      attendanceGroupMembers: [] as IGroupAttendance[],

    }
  },
  actions: {


  }
});