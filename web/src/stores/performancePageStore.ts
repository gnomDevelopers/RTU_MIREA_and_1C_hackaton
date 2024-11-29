import { defineStore } from "pinia";
import{ type IGroupAttendance, type IGroupScores, type TMaybeNumber, type TMaybeString } from "@/helpers/constants";

export const usePerformancePageStore = defineStore('performancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,
      selectedDiscipline: null as TMaybeString,

      groupMembersScores: [] as IGroupScores[],
    }
  }
});