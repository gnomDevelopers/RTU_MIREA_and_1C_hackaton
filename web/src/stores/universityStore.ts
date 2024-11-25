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
} from "@/helpers/constants";

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
        {id: 1, address: 'Малая Пироговская улица, 1', name: 'МП-1', university_id: 1}, 
        {id: 2, address: 'Проспект Вернадского, 78', name: 'В-78', university_id: 1}, 
        {id: 3, address: 'Проспект Вернадского, 86', name: 'В-86', university_id: 1},
      ];

      // for(let i = 1; i < 110; i++) {
      //   const data: IUserGet = {id: i, surname: `Деканов${i}`, name: `Декан${i}`, thirdname: `Деканович${i}`, role: 2, faculty_id: 1, department_id: -1, educational_direction: '',};
      //   this.decansList.push(data);
      // }
      this.decansList = [
        {id: 1, surname: `Иванова`, name: `Алина`, thirdname: `Сергеевна`, role: 2, faculty_id: 1, department_id: -1, educational_direction: '',},
        {id: 2, surname: `Петров`, name: `Дмитрий`, thirdname: `Алексеевич `, role: 2, faculty_id: 1, department_id: -1, educational_direction: '',},
      ];  
      // for(let i = 1; i < 110; i++) {
      //   const data: IUserGet = {id: i, surname: `УчОтделов${i}`, name: `УчОтдел${i}`, thirdname: `УчОтделович${i}`, role: 3, faculty_id: 1, department_id: -1, educational_direction: '',};
      //   this.educationDepartmentsList.push(data);
      // }
      this.educationDepartmentsList = [
        {id: 3, surname: `Смирнова`, name: `Елена`, thirdname: `Владимировна`, role: 3, faculty_id: 1, department_id: -1, educational_direction: '',},
        {id: 4, surname: `Кузнецов`, name: `Сергей`, thirdname: `Николаевич`, role: 3, faculty_id: 1, department_id: -1, educational_direction: '',},
      ];  
      // for(let i = 1; i < 110; i++) {
      //   const data: IUserGet = {id: i, surname: `ЗавКафов${i}`, name: `ЗавКаф${i}`, thirdname: `ЗавКафович${i}`, role: 4, faculty_id: 1, department_id: 1, educational_direction: '',};
      //   this.zavCafsList.push(data);
      // }
      this.zavCafsList = [
        {id: 5, surname: `Попова`, name: `Ольга`, thirdname: `Павловна`, role: 4, faculty_id: 1, department_id: 1, educational_direction: '',},
        {id: 6, surname: `Соколов`, name: `Андрей`, thirdname: `Михайлович`, role: 4, faculty_id: 1, department_id: 1, educational_direction: '',},
      ];    
      // for(let i = 1; i < 110; i++) {
      //   const data: IUserGet = {id: i, surname: `Преподов${i}`, name: `Препод${i}`, thirdname: `Преподович${i}`, role: 5, faculty_id: 1, department_id: 1, educational_direction: '',};
      //   this.teachersList.push(data);
      // }
      this.teachersList = [
        {id: 7, surname: `Морозова`, name: `Наталья`, thirdname: `Юрьевна`, role: 5, faculty_id: 1, department_id: 1, educational_direction: '',},
        {id: 8, surname: `Васильев`, name: `Алексей`, thirdname: `Иванович`, role: 5, faculty_id: 1, department_id: 1, educational_direction: '',},
      ];  
      // for(let i = 1; i < 110; i++) {
      //   const data: IUserGet = {id: i, surname: `Студентов${i}`, name: `Студент${i}`, thirdname: `Студентович${i}`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',};
      //   this.studentsList.push(data);
      // }
      this.studentsList = [
        {id: 9, surname: `Новикова`, name: `Екатерина`, thirdname: `Александровна`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
        {id: 10, surname: `Волков`, name: `Павел`, thirdname: `Дмитриевич`, role: 6, faculty_id: 1, department_id: 1, educational_direction: 'Фуллстек разработка',},
      ];  

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

      this.groupMembersList = this.sortPeople(this.groupMembersList);

      this.groupMembersScores = [];
      for(let user of this.groupMembersList){
        const data = {user: user, scores: [], avg: 0} as IGroupScores;
        let summ = 0;
        let count = 0;
        for(let i = 0; i < 16; i++) {
          let j = Math.ceil(Math.random() * 100) <= 30 ? Math.ceil(Math.random() * 5) : 0;
          summ += j;
          if(j !== 0) count ++;
          data.scores.push(j);
        }
        data.avg = summ / count;
        this.groupMembersScores.push(data);
      }


        
      for(let i = 1; i < 10; i++) {
        const data: IAPI_Audience_Update = {  
          id: i,
          campus_id: i % 3 + 1,
          capacity: Math.ceil(Math.random() * 100) + 20,
          name: `A-${i}`,
          profile: AUDITORY_PROFILE_LIST[i % AUDITORY_PROFILE_LIST.length],
          type: AUDITORY_TYPE_LIST[i % AUDITORY_TYPE_LIST.length],
        };
        this.auditoriesList.push(data);
      }

      for(let i = 1; i < 10; i++) {
        const data: IGroup = {id: i, name: `ЭФБО-0${i}-23`};
        this.groupsList.push(data);
      }

      this.facultiesList = [ // факультеты
        {id: 1, name: 'ИПТИП'},
        {id: 2, name: 'ИИТ'},
        {id: 3, name: 'ИКБ'},
      ];
      this.deparmentsList = [ // кафедры
        {id: 1, name: 'Физика'},
        {id: 2, name: 'Высшая математика'},
        {id: 3, name: 'Индустриальне программирование'},
      ];
      this.educationalDirectionsList = [ // название направления
        {id: 1, name: 'Фуллстек разработка'},
        {id: 2, name: 'Программная инженерия'},
        {id: 3, name: 'Компьютерный дизайн'},
      ];
      this.facultativesList = [
        {id: 1, name: 'Мобильная разработка Kotlin'},
        {id: 2, name: 'DevOps курсы'},
        {id: 3, name: 'Автошкола'},
      ];
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
    }
  }
});