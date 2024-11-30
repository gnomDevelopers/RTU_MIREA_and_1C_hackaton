<template>
  <div>
    <div class="flex flex-col md:flex-row w-full md:w-auto items-stretch gap-1 p-2 rounded-lg self-start bg-blue-300">
      <div class="flex flex-row flex-wrap gap-1">
        <input 
          @keyup.enter="$refs.inputName.focus()" 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="userSurname"
          placeholder="Фамилия">
        
        <input 
          ref="inputName" 
          @keyup.enter="$refs.inputThirdname.focus()" 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="userName"
          placeholder="Имя">
        
        <input 
          ref="inputThirdname" 
          @keyup.enter="$refs.inputRole.focus()" 
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="userThirdname"
          placeholder="Отчество">
        
        <select 
          ref="inputRole" 
          v-model="userRole" 
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите роль</option>
          <option class="text-lg" :value="2">{{getRolesName(2)}}</option>
          <option class="text-lg" :value="3">{{getRolesName(3)}}</option>
          <option class="text-lg" :value="4">{{getRolesName(4)}}</option>
          <option class="text-lg" :value="5">{{getRolesName(5)}}</option>
          <option class="text-lg" :value="6">{{getRolesName(6)}}</option>
        </select>

        <select 
          v-if="InfoFieldsRole[userRole]?.includes(InfoFields[0])"
          v-model="userFaculty" 
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите факультет</option>
          <option v-for="item in getFaculties" class="text-lg" :value="item.id">{{ item.name }}</option>
        </select>

        <select 
          v-if="InfoFieldsRole[userRole]?.includes(InfoFields[1])"
          v-model="userDepartment" 
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
          
          <option class="text-lg" :value="-1" disabled>Выберите кафедру</option>
          <option v-for="item in getDepartments" class="text-lg" :value="item.id">{{ item.name }}</option>
        </select>

        <input 
          v-if="InfoFieldsRole[userRole]?.includes(InfoFields[2])"
          class="min-w-20 w-full md:w-40 max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800" 
          type="text" 
          v-model="userEducationalDirections"
          placeholder="Направление">

          <select 
            v-if="InfoFieldsRole[userRole]?.includes(InfoFields[2])"
            v-model="userGroup" 
            class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg border-2 border-solid border-transparent focus:border-blue-800">
            
            <option class="text-lg" :value="-1" disabled>Выберите группу</option>
            <option v-for="item in getGroups" class="text-lg" :value="item.id">{{ item.name }}</option>
          </select>
      </div>
      
      <div class="flex flex-row flex-shrink-0 gap-1 items-center">
        <div @click="addUser" class="rounded-lg p-1 cursor-pointer btn">
          <img class="w-8 h-8" src="../assets/icons/icon-plus.svg"/>
        </div>
        <div v-if="usersList.length !== 0" @click="sendUsersList" class="px-2 h-10 rounded-lg cursor-pointer btn">
          <p class="text-white text-xl">Отправить</p>
        </div>
        <JsonExcel 
          :data="usersJsonData"
          :fields="usersJsonFields"
          type="xlsx"
          worksheet="My Worksheet"
          name="Данные_пользователей.xlsx"
          >
            Скачать данные
        </JsonExcel>
      </div>
    </div>

    <div v-if="usersList.length !== 0" class="flex flex-col gap-y-2 p-2 rounded-lg max-h-[400px] scrollable bg-white">
      <div v-for="user in usersList" class="flex flex-row gap-x-2 px-2 py-1 items-center rounded-lg bg-color-light">
        
        <div class="flex flex-row flex-grow justify-start items-center">
          <p class="text-lg">{{ user.last_name }} {{user.first_name}} {{ user.father_name }}</p>
        </div>

        <div class="flex flex-row flex-wrap gap-1 px-2">

          <p 
            class="font-semibold text-lg cursor-default">
            Роль: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">
              {{ getRolesName(user.role) }}
            </span>
          </p>

          <p 
            v-if="showFaculty(user.role)" 
            class="font-semibold text-lg cursor-default">
            Факультет: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">
              {{ getFacultyName(user.faculty_id) }}
            </span>
          </p>

          <p 
            v-if="showDepartment(user.role)" 
            class="font-semibold text-lg cursor-default">
            Кафедра: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">
              {{ getDepartmentName(user.department_id) }}
            </span>
          </p>

          <p 
            v-if="showEducationalDirection(user.role)" 
            class="font-semibold text-lg cursor-default">
            Направление: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">
              {{ user.educational_direction }}
            </span>
          </p>

          <p 
            v-if="showEducationalDirection(user.role)" 
            class="font-semibold text-lg cursor-default">
            Группа: 
            <span class=" font-normal text-base cursor-pointer text-blue-800">
              {{ getGroupName(user.group_id) }}
            </span>
          </p>

        </div>
        <div @click="deleteUser(user)" class="w-9 h-9 flex flex-row flex-shrink-0 justify-center items-center rounded-lg cursor-pointer btn">
          <img class="w-6 h-7" src="../assets/icons/icon-delete.svg"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUniversityStore } from '@/stores/universityStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { 
  ROLES_NAME, 
  INFO_FIELDS, 
  INFO_FIELDS_ROLE, 
  StatusCodes, 
  type IUser,
  type IUserCreate,
  type IUserCreateWithStringRole,
} from '../helpers/constants';
import { API_University_Users_Create } from '@/api/api';
import JsonExcel from "vue-json-excel3";

export default {
  components: {
    JsonExcel,
  },
  data(){
    return{
      RolesName: ROLES_NAME,
      InfoFields: INFO_FIELDS,
      InfoFieldsRole: INFO_FIELDS_ROLE,

      userName: '',
      userSurname: '',
      userThirdname: '',
      userRole: -1,
      userFaculty: -1,
      userDepartment: -1,
      userEducationalDirections: '',
      userGroup: -1,
      usersList: [] as IUserCreate[],

      usersJsonData: [] as { last_name: string; first_name: string; father_name: string; email: string; password: string; }[],
      usersJsonFields: {
        "last_name": "Фамилия",
        "first_name": "Имя",
        "father_name": "Отчество",
        "email": "Email",
        "password": "Пароль для входа",
      },
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useUniversityStore, useUserInfoStore),

    getFaculties(){
      return this.universityStore.facultiesList;
    },
    getDepartments(){
      return this.universityStore.deparmentsList;
    },
    getGroups(){
      return this.universityStore.groupsList;
    }
  },
  methods:{
    getRolesName(ind: number){
      return ROLES_NAME[ind];
    },
    addUser(){
      if(
        this.userName !== '' && 
        this.userSurname !== '' && 
        this.userRole !== -1 && 
        (this.showFaculty(this.userRole) ? this.showFaculty(this.userRole) && this.userFaculty !== -1 : true) &&
        (this.showDepartment(this.userRole) ? this.showDepartment(this.userRole) && this.userDepartment!== -1 : true) &&
        (this.showEducationalDirection(this.userRole) ? this.userEducationalDirections !== '' : true) &&
        (this.showEducationalDirection(this.userRole) ? this.userGroup !== -1 : true)
      ){
        this.usersList.push({
          first_name: this.userName, 
          last_name: this.userSurname, 
          father_name: this.userThirdname, 
          role: this.userRole,
          faculty_id: this.userFaculty === -1 ? 1 : this.userFaculty,
          department_id: this.userDepartment === -1 ? 1 : this.userDepartment,
          educational_direction: this.userEducationalDirections === '' ? 'null' : this.userEducationalDirections,
          group_id: this.userGroup === -1 ? 1 : this.userGroup,
          university_id: this.universityStore.getUniversityID(this.userInfoStore.university!),
          password: '',
        });

        this.userName = '';
        this.userSurname = '';
        this.userThirdname = '';
        this.userFaculty = -1;
        this.userDepartment = -1;
        this.userEducationalDirections = '';
        this.userGroup = -1;
        return;
      }
      if(this.userName === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите имя пользователя!');
      }
      if(this.userSurname === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите фамилию пользователя!');
      }
      if(this.userRole === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите роль пользователя!');
      }
      if(this.showFaculty(this.userRole) && this.userFaculty === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите факультет!');
      }
      if(this.showDepartment(this.userRole) && this.userDepartment === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Выберите кафедру!');
      }
      if(this.showEducationalDirection(this.userRole) && this.userEducationalDirections === ''){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите направление!');
      }
      if(this.showEducationalDirection(this.userRole) && this.userGroup === -1){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите группу!');
      }
    },
    deleteUser(userDelete: IUser){
      this.usersList = this.usersList.filter(user => user !== userDelete);
    },
    sendUsersList(){
      if(this.usersList.length === 0) return;
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);

      const body: IUserCreateWithStringRole[] = [];
      for(let item of this.usersList){
        body.push({
          password: item.password,
          department_id: item.department_id,
          educational_direction: item.educational_direction,
          faculty_id: item.faculty_id,
          father_name: item.father_name,
          first_name: item.first_name,
          group_id: item.group_id,
          last_name: item.last_name,
          role: ROLES_NAME[item.role],
          university_id: item.university_id,
        });
      }

      API_University_Users_Create(body)
      .then((response: any) => {

        //распределяем пользователей по соответствующим спискам
        for(let item of this.usersList){
          switch(item.role){
            case 2: this.universityStore.decansList.push({...item, group_id: 1, id: this.universityStore.tmpuserID++}); break;
            case 3: this.universityStore.educationDepartmentsList.push({...item, group_id: 1, id: this.universityStore.tmpuserID++}); break;
            case 4: this.universityStore.zavCafsList.push({...item, group_id: 1, id: this.universityStore.tmpuserID++}); break;
            case 5: this.universityStore.teachersList.push({...item, group_id: 1, id: this.universityStore.tmpuserID++}); break;
            case 6: this.universityStore.studentsList.push({...item, group_id: 1, id: this.universityStore.tmpuserID++}); break;
          }
        }

        this.usersJsonData = [];
        for(let item of response.data){
          const userData: { last_name: string; first_name: string; father_name: string; email: string; password: string; } = {
            last_name: item.last_name,
            first_name: item.first_name,
            father_name: item.father_name,
            email: item.email,
            password: item.password,
          }
          this.usersJsonData.push(userData);
        }
        console.log('usersJsonData: ', this.usersJsonData);
        //очистка буфера пользователей
        this.usersList = [];

        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Пользователи добавлены!');
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при добавлении пользователей!');
      });
    },

    showFaculty(role: number): boolean{
      return this.InfoFieldsRole[role]?.includes(this.InfoFields[0]);
    },
    showDepartment(role: number): boolean{
      return this.InfoFieldsRole[role]?.includes(this.InfoFields[1]);
    },
    showEducationalDirection(role: number): boolean{
      return this.InfoFieldsRole[role]?.includes(this.InfoFields[2]);
    },
    getFacultyName(facultyID: number): string{
      for(let item of this.universityStore.facultiesList){
        if(item.id === facultyID) return item.name;
      }
      return '';
    },
    getDepartmentName(departmentID: number): string{
      for(let item of this.universityStore.deparmentsList){
        if(item.id === departmentID) return item.name;
      }
      return '';
    },
    getGroupName(groupID: number): string{
      for(let item of this.universityStore.groupsList){
        if(item.id === groupID) return item.name;
      }
      return '';
    }
  }
}; 
</script>