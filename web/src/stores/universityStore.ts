import { defineStore } from "pinia";
import { 
  AUDITORY_PROFILE_LIST, 
  AUDITORY_TYPE_LIST, 
  type IAPI_Audience_Update, 
  type IAPI_Campus_Update, 
  type IUserGet, 
  type IGroup, 
  type IItemList,
  type IGroupScores,
  ROLES_NAME,
} from "@/helpers/constants";
import { useUserInfoStore } from "./userInfoStore";
import { API_Audience_Get, API_Campus_Get_University, API_Departments_Get, API_Facultatives_Get, API_Faculties_Get, API_University_Users_Get } from "@/api/api";

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

      groupsList: [] as IGroup[],
      facultativesList: [] as IItemList[],
      groupMembersList: [] as IUserGet[],
      groupMembersScores: [] as IGroupScores[],

      facultiesList: [] as IItemList[],
      deparmentsList: [] as IItemList[],
      educationalDirectionsList: [] as IItemList[],

      tmpuserID: 100,
      selectedDate: '25.11.2024',
    }
  },
  actions: {
    async loadUniversityInfo(){
      const userInfo = useUserInfoStore();

      if(userInfo.university !== null) API_Campus_Get_University(userInfo.university)
      .then((response: any) => {
        this.campusList = response.data
      })
      .catch(error => {
        this.campusList = [];
      })

      let UniversityUsers: any[] = [];
      if(userInfo.university !== null) API_University_Users_Get(userInfo.university)
      .then((response: any) => {
        UniversityUsers = response.data;
      })
      .catch(error => {
        UniversityUsers = [];
      })

      for(let user of UniversityUsers){
        const userData = {
          id: user.id, 
          surname: user.last_name, 
          name: user.first_name, 
          thirdname: user.father_name, 
          role: user.role, 
          faculty_id: user.faculty_id, 
          department_id: user.department_id, 
          educational_direction: user.educational_direction
        } as IUserGet;

        switch(user.role){
          case ROLES_NAME[2]: this.decansList.push(userData); break;
          case ROLES_NAME[3]: this.educationDepartmentsList.push(userData); break;
          case ROLES_NAME[4]: this.zavCafsList.push(userData); break;
          case ROLES_NAME[5]: this.teachersList.push(userData); break;
          case ROLES_NAME[6]: this.studentsList.push(userData); break;
        }
      }

      this.groupMembersList = [
        {id: 11, surname: `Романова`, name: `Анастасия`, thirdname: `Игоревна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 12, surname: `Белов`, name: `Кирилл`, thirdname: `Олегович`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 13, surname: `Крылова`, name: `Ирина`, thirdname: `Викторовна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 14, surname: `Орлов`, name: `Денис`, thirdname: `Сергеевич`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 15, surname: `Савина`, name: `Светлана`, thirdname: `Николаевна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 16, surname: `Богданов`, name: `Евгений`, thirdname: `Александрович`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 17, surname: `Титова`, name: `Мария`, thirdname: `Дмитриевна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 18, surname: `Сидоров`, name: `Даниил`, thirdname: `Алексеевич`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 19, surname: `Волкова`, name: `Юлия`, thirdname: `Андреевна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 20, surname: `Лебедев`, name: `Антон`, thirdname: `Валерьевич`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
      ];

      this.groupMembersScores = [];
      for(let user of this.groupMembersList){
        const data = {user: user, scores: [], avg: 0, gpa: 0} as IGroupScores;
        let summ = 0;
        let count = 0;
        for(let i = 0; i < 16; i++) {
          let j = Math.ceil(Math.random() * 100) <= 30 ? Math.ceil(Math.random() * 5) : 0;
          summ += j;
          if(j !== 0) count ++;
          data.scores.push(j);
        }
        data.avg = summ / count;
        data.gpa = data.avg + ((Math.ceil(Math.random() * 10)) > 5 ? 1 : -1) * (Math.random()*0.3);
        this.groupMembersScores.push(data);
      }
      this.groupMembersScores = this.sortByName(this.groupMembersScores);
        

      if(userInfo.university !== null) API_Audience_Get(userInfo.university)
      .then((response:any) => {
        this.auditoriesList = response.data;
      })
      .catch(error => {
        this.auditoriesList = [];
      });


      for(let i = 1; i < 10; i++) {
        const data: IGroup = {id: i, name: `ЭФБО-0${i}-23`};
        this.groupsList.push(data);
      }

      if(userInfo.university !== null) API_Faculties_Get(userInfo.university) // факультеты
      .then((response:any) => {
        this.facultiesList = [];
        for(let item of response.data) this.facultiesList.push({id: item.id, name: item.name});
      })
      .catch(error => {
        this.facultiesList = [];
      })

      if(userInfo.university !== null) API_Departments_Get(userInfo.university) // кафедры
      .then((response:any) => {
        this.deparmentsList = response.data;
      })
      .catch(error => {
        this.deparmentsList = [];
      })

      this.educationalDirectionsList = [ // название направления
        {id: 1, name: 'Фуллстек разработка'},
        {id: 2, name: 'Программная инженерия'},
        {id: 3, name: 'Компьютерный дизайн'},
      ];


      API_Facultatives_Get()
      .then((response:any) => {
        if(response.data) this.facultativesList = response.data;
      })
      .catch(error => {
        this.facultativesList = [];
      });
    },
    sortByName(people: IGroupScores[]): IGroupScores[] {
      return [...people].sort((a, b) => {
        const surnameComparison = a.user.surname.localeCompare(b.user.surname);
        if (surnameComparison !== 0) {
          return surnameComparison;
        }

        const nameComparison = a.user.name.localeCompare(b.user.name);
        if (nameComparison !== 0) {
          return nameComparison;
        }

        return a.user.thirdname.localeCompare(b.user.thirdname);
      });
    },
    sortByGpa(students: IGroupScores[]): IGroupScores[] {
      return [...students].sort((a, b) => b.gpa - a.gpa);
    },
    sortPeople(people: IUserGet[]): IUserGet[] {
      return [...people].sort((a, b) => {
        const surnameComparison = a.surname.localeCompare(b.surname);
        if (surnameComparison !== 0) {
          return surnameComparison;
        }

        const nameComparison = a.name.localeCompare(b.name);
        if (nameComparison !== 0) {
          return nameComparison;
        }

        return a.thirdname.localeCompare(b.thirdname);
      });
    },
  }
});