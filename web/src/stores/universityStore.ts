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
  GET_ROLEID_BY_ROLENAME,
} from "@/helpers/constants";
import { useUserInfoStore } from "./userInfoStore";
import { API_Audience_Get, API_Campus_Get_University, API_Departments_Get, API_Facultatives_Get, API_Faculties_Get, API_University_Groups_Get, API_University_Users_Get } from "@/api/api";

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
      facultativesList: [] as IGroup[],
      groupMembersList: [] as IUserGet[],
      groupMembersScores: [] as IGroupScores[],

      facultiesList: [] as IItemList[],
      deparmentsList: [] as IItemList[],

      tmpuserID: 100,
      selectedDate: '25.11.2024',
    }
  },
  actions: {
    async loadUniversityInfo(){
      const userInfo = useUserInfoStore();
      if(userInfo.university === null) {
        //все обнуляем
        this.campusList = [];
        this.auditoriesList = [];
  
        this.decansList = [];
        this.educationDepartmentsList = [];
        this.zavCafsList = [];
        this.teachersList = [];
        this.studentsList = [];
  
        this.groupsList = [];
        this.facultativesList = [];
        this.groupMembersList = [];
        this.groupMembersScores = [];
  
        this.facultiesList = [];
        this.deparmentsList = [];
        return;
      }

      this.loadCampus(userInfo.university);
      this.loadAuditories(userInfo.university);
      this.loadFaculties(userInfo.university);
      this.loadDepartments(userInfo.university);
      this.loadFacultatives();
      this.loadAllGroups();
      await this.loadUsers(userInfo.university);


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

    },


    async loadCampus(universityName: string){
      API_Campus_Get_University(universityName)
        .then((response: any) => {
          this.campusList = response.data;
        })
        .catch(error => {
          this.campusList = [];
        })
    },
    async loadUsers(universityName: string){
      let UniversityUsers: any[] = [];

      API_University_Users_Get(universityName)
      .then((response: any) => {
        UniversityUsers = response.data;
      })
      .catch(error => {
        UniversityUsers = [];
      });

      this.decansList = [];
      this.educationDepartmentsList = [];
      this.zavCafsList = [];
      this.teachersList = [];
      this.studentsList = [];

      for(let user of UniversityUsers){
        const userData = {
          id: user.id, 
          surname: user.last_name, 
          name: user.first_name, 
          thirdname: user.father_name, 
          role: GET_ROLEID_BY_ROLENAME(user.role), 
          faculty_id: user.faculty_id, 
          department_id: user.department_id, 
          educational_direction: user.educational_direction,
          group_id: user.group_id,
        } as IUserGet;

        switch(user.role){
          case ROLES_NAME[2]: this.decansList.push(userData); break;
          case ROLES_NAME[3]: this.educationDepartmentsList.push(userData); break;
          case ROLES_NAME[4]: this.zavCafsList.push(userData); break;
          case ROLES_NAME[5]: this.teachersList.push(userData); break;
          case ROLES_NAME[6]: this.studentsList.push(userData); break;
        }
      }
    },
    findGroupMembers(groupID: number){
      this.groupMembersList = [];
      for(let user of this.studentsList){
        this.groupMembersList = [];
        if(user.group_id === groupID) this.groupMembersList.push(user);
      }
    },
    async loadAuditories(universityName: string){
      API_Audience_Get(universityName)
      .then((response:any) => {
        this.auditoriesList = response.data;
      })
      .catch(error => {
        this.auditoriesList = [];
      });
    },
    async loadFaculties(universityName: string){
      API_Faculties_Get(universityName) // факультеты
      .then((response:any) => {
        this.facultiesList = [];
        for(let item of response.data) this.facultiesList.push({id: item.id, name: item.name});
      })
      .catch(error => {
        this.facultiesList = [];
      });
    },
    async loadDepartments(universityName: string){
      API_Departments_Get(universityName) // кафедры
      .then((response:any) => {
        this.deparmentsList = response.data;
      })
      .catch(error => {
        this.deparmentsList = [];
      })
    },
    async loadFacultatives(){
      API_Facultatives_Get()
      .then((response:any) => {
        this.facultativesList = response.data;
      })
      .catch(error => {
        this.facultativesList = [];
      });
    },
    async loadAllGroups(){
      API_University_Groups_Get()
      .then((response:any) => {
        this.groupsList = response.data.groups;
      })
      .catch(error => {
        this.groupsList = [];
      });
    },

    //вынести в константы!
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