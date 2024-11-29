import { defineStore } from "pinia";
import{ type IGroupAttendance, type IDataItem, type TMaybeNumber, type TMaybeString, type IReaorganizedGroupScore } from "@/helpers/constants";
import { API_Grades_Group_Discipline_Get } from "@/api/api";
import { reorganizeGroupScores, transformData } from "@/helpers/gradesParser";

export const usePerformancePageStore = defineStore('performancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,
      selectedDiscipline: null as TMaybeString,

      groupGrades: [] as IReaorganizedGroupScore[],
    }
  },
  actions: {
    async loadGroupGrades(){
      try{
        if(this.selectedGroup === null || this.selectedDiscipline === null) return;
        const response: any = API_Grades_Group_Discipline_Get(this.selectedGroup, this.selectedDiscipline);

        this.groupGrades = [];

        const transformedData: IDataItem = transformData(response.data);
        console.log('added empty grades: ', transformedData);
        this.groupGrades = reorganizeGroupScores( transformedData );
        console.log('reorganized data: ', this.groupGrades);
        //запросить gpa
      }catch(error){
        this.groupGrades = [];
      }
    }
  }
});