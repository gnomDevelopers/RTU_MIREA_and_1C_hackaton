import { defineStore } from "pinia";
import { StatusCodes, type IAPI_Audience_Update, type IAPI_Campus_Update, type IUserGet, AUDITORY_PROFILE_LIST, AUDITORY_TYPE_LIST } from "@/helpers/constants";
import { useStatusWindowStore } from "./statusWindowStore";

const statusWindow = useStatusWindowStore();

export const useUniversityStore = defineStore('university', {
  state() {
    return{
      campusList: [] as IAPI_Campus_Update[],
      auditoriesList: [] as IAPI_Audience_Update[],

      //уточнить тип!
      decansList: [] as IUserGet[],
      educationDepartmentsList: [] as IUserGet[],
      zavCafsList: [] as IUserGet[],
      teachersList: [] as IUserGet[],
      studentsList: [] as IUserGet[],
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

      for(let i = 1; i < 110; i++) {
        const data: IUserGet = {id: i, surname: `Деканов${i}`, name: `Декан${i}`, thirdname: `Деканович${i}`, role: 2, faculty_id: 1, department_id: -1, educational_direction: '',};
        this.decansList.push(data);
      }

      for(let i = 1; i < 110; i++) {
        const data: IUserGet = {id: i, surname: `УчОтделов${i}`, name: `УчОтдел${i}`, thirdname: `УчОтделович${i}`, role: 3, faculty_id: 1, department_id: -1, educational_direction: '',};
        this.educationDepartmentsList.push(data);
      }
  
      for(let i = 1; i < 110; i++) {
        const data: IUserGet = {id: i, surname: `ЗавКафов${i}`, name: `ЗавКаф${i}`, thirdname: `ЗавКафович${i}`, role: 4, faculty_id: 1, department_id: 1, educational_direction: '',};
        this.zavCafsList.push(data);
      }

      for(let i = 1; i < 110; i++) {
        const data: IUserGet = {id: i, surname: `Преподов${i}`, name: `Препод${i}`, thirdname: `Преподович${i}`, role: 5, faculty_id: 1, department_id: 1, educational_direction: '',};
        this.teachersList.push(data);
      }

      for(let i = 1; i < 110; i++) {
        const data: IUserGet = {id: i, surname: `Студентов${i}`, name: `Студент${i}`, thirdname: `Студентович${i}`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',};
        this.studentsList.push(data);
      }
  
      for(let i = 1; i < 110; i++) {
        const data: IAPI_Audience_Update = {  
          id: i,
          campus_id: i % 3,
          capacity: 123,
          name: `A-${i}`,
          profile: AUDITORY_PROFILE_LIST[i % 8],
          type: AUDITORY_TYPE_LIST[i % 5],
        };
        this.auditoriesList.push(data);
      }
    }
  }
});