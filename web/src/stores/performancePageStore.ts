import { defineStore } from "pinia";
import{ type IGroupAttendance, type IGroupScores, type TMaybeNumber, type TMaybeString } from "@/helpers/constants";
import { API_Grades_Group_Discipline_Get } from "@/api/api";

export const usePerformancePageStore = defineStore('performancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,
      selectedDiscipline: null as TMaybeString,

      groupGrades: [] as IGroupScores[],
    }
  },
  actions: {
    async loadGroupGrades(){
      try{
        if(this.selectedGroup === null || this.selectedDiscipline === null) return;
        const response: any = API_Grades_Group_Discipline_Get(this.selectedGroup, this.selectedDiscipline);

      }catch(error){
        this.groupGrades = [];
      }
    }
  }
});