import { defineStore } from "pinia";
import { StatusCodes, type IAPI_Audience_Update, type IAPI_Campus_Update, type IUser } from "@/helpers/constants";
import { useStatusWindowStore } from "./statusWindowStore";

const statusWindow = useStatusWindowStore();

export const useUniversityStore = defineStore('university', {
  state() {
    return{
      campusList: [] as IAPI_Campus_Update[],
      auditoriesList: [] as IAPI_Audience_Update[],

      //уточнить тип!
      decansList: [] as IUser[],
      educationDepartmentsList: [] as IUser[],
      zavCafsList: [] as IUser[],
      teachersList: [] as IUser[],
      studentsList: [] as IUser[],
    }
  },
  actions: {
    loadUniversityInfo(){
      //API get campuses by universityID
      //API get audiences by universityID
      //API get decans by universityID
      //API get aducationDepartments by universityID
      //API get zavCafs by universityID
      //API get teachers by universityID
      //API get students by universityID

      this.campusList = [
        {id: 1, address: 'Ул. Пушкина', name: 'МП-1', university_id: 1}, 
        {id: 2, address: 'Ул. Пушкина', name: 'В-78', university_id: 1}, 
        {id: 3, address: 'Ул. Пушкина', name: 'В-82', university_id: 1},
      ];
    }
  }
});