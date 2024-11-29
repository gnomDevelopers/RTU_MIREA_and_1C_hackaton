import { defineStore } from "pinia";
import{ type IGroupAttendance, type IDataItem, type TMaybeNumber, type TMaybeString, type IReaorganizedGroupScore } from "@/helpers/constants";
import { API_GPA_Get, API_Grades_Group_Discipline_Get } from "@/api/api";
import { reorganizeGroupScores, transformData } from "@/helpers/gradesParser";

export const usePerformancePageStore = defineStore('performancePage', {
  state() {
    return{
      selectedGroup: null as TMaybeString,
      selectedDiscipline: null as TMaybeString,

      groupGrades: [] as IReaorganizedGroupScore[],
      isGroupGradesEmpty: true,
    }
  },
  actions: {
    async loadGroupGrades(){
      try{
        if(this.selectedGroup === null || this.selectedDiscipline === null) return;
        const response: any = await API_Grades_Group_Discipline_Get(this.selectedGroup, this.selectedDiscipline);

        const transformedData: IDataItem = transformData({ 
          grade_class: response.data.grade_class, 
          group_member: response.data.group_member, 
          users_score: response.data.users_score ? response.data.users_score : [],
        });

        console.log('added empty grades: ', transformedData);

        this.groupGrades = reorganizeGroupScores( transformedData );

        console.log('reorganized data: ', this.groupGrades);
        for(let i = 0; i < this.groupGrades.length; i++){
          API_GPA_Get(this.groupGrades[i].id)
          .then((response: any) => {
            this.groupGrades[i].gpa = Number(response.data.value);
            if(isNaN(this.groupGrades[i].gpa)) this.groupGrades[i].gpa = 0;
          })
          .catch(error => {
            this.groupGrades[i].gpa = 0;
          })
        }
        console.log('after get gpa: ', this.groupGrades);

        
        this.isGroupGradesEmpty = false;
        //запросить gpa

      }catch(error){
        this.groupGrades = [{
          first_name: '',
          last_name: '',
          father_name: '',
          id: 0,
          grades: [{
              date: '',
              class_id: 0,
              id: 0,
              user_id: 0,
              value: 0,
          }],
          average_score: 0,
          sum_score: 0,
          user_id: 0,
          gpa: 0,
        }];
        this.isGroupGradesEmpty = true;
      }
    },
    // sortByName(people: IGroupScores[]): IGroupScores[] {
    //   return [...people].sort((a, b) => {
    //     const surnameComparison = a.user.last_name.localeCompare(b.user.last_name);
    //     if (surnameComparison !== 0) {
    //       return surnameComparison;
    //     }

    //     const nameComparison = a.user.first_name.localeCompare(b.user.first_name);
    //     if (nameComparison !== 0) {
    //       return nameComparison;
    //     }

    //     return a.user.father_name.localeCompare(b.user.father_name);
    //   });
    // },
    // sortByGpa(students: IGroupScores[]): IGroupScores[] {
    //   return [...students].sort((a, b) => b.gpa - a.gpa);
    // },
  }
});