import { defineStore } from "pinia";
import{ type IGroupAttendance, type TMaybeNumber } from "@/helpers/constants";

export const useAttendacePageStore = defineStore('attendancePage', {
  state() {
    return{
      selectedGroupID: null as TMaybeNumber,

      attendanceGroupMembers: [] as IGroupAttendance[],
    }
  }
});