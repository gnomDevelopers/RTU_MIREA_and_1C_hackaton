<template>
  <div class="flex flex-col sm:flex-row items-start sm:items-center p-2 rounded-lg bg-color-light">
    <div class="flex-grow">
      <p class="text-lg">{{ data.last_name }} {{ data.first_name }} {{ data.father_name }}</p>
    </div>
    <div class="flex flex-row gap-x-2 items-center flex-shrink-0 ">
      <select 
          ref="inputRole" 
          v-model="userRole" 
          class="min-w-20 w-full md:w-auto max-w-none px-2 py-1 text-lg outline-none rounded-lg header-shadow border-2 border-solid border-transparent focus:border-blue-800">
          
          <option 
            v-for="roleID in enabledRoles" 
            :key="roleID" 
            :value="roleID" 
            :selected="data.role === roleID"
            class="text-lg">
            
            {{getRolesName(roleID)}}
          </option>

        </select>
      <div v-if="userRole !== data.role" @click="updateUser" class="py-1 px-2 rounded-lg cursor-pointer text-base select-none btn">Сохранить</div>
      <div @click="deleteUser" class="py-1 px-2 rounded-lg cursor-pointer text-base select-none btn btn-bad">{{ data.role !== 6 ? 'Уволить' : 'Отчислить' }}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { useUniversityStore } from '@/stores/universityStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { ROLES_NAME, type IUserGet, StatusCodes } from '@/helpers/constants';
import type { PropType } from 'vue';
export default{
  props:{
    data:{
      type: Object as PropType<IUserGet>,
      required: true,
    }
  },
  data() {
    return {
      enabledRoles: [2, 3, 4, 5, 6],
      userRole: this.data.role,
    }
  },
  computed:{
    ...mapStores(useUserInfoStore, useUniversityStore, useStatusWindowStore),

  },
  methods: {
    getRolesName(roleID: number): string{
      return ROLES_NAME[roleID];
    },
    getUserList(role: number){
      switch (role){
        case 2: return this.universityStore.decansList;
        case 3: return this.universityStore.educationDepartmentsList;
        case 4: return this.universityStore.zavCafsList;
        case 5: return this.universityStore.teachersList;
        case 6: return this.universityStore.studentsList;
        default: return [];
      }
    },
    updateUser(){
      if(this.userRole === this.data.role) return;

      //удаляем из прошлого списка
      const list = this.getUserList(this.data.role);
      for(let i = 0; i < list.length; i++){
        if(list[i].id === this.data.id){
          list.splice(i, 1);
          break;
        }
      }
      //добавляем в новый список
      this.getUserList(this.userRole).push(<IUserGet>{
        id: this.data.id,
        first_name: this.data.first_name,
        last_name: this.data.last_name,
        father_name: this.data.father_name,
        role: this.userRole,
        faculty_id: this.data.faculty_id,
        department_id: this.data.department_id,
        educational_direction: this.data.educational_direction,
        group_id: this.data.group_id,
        university_id: 1,
      });
      this.statusWindowStore.showStatusWindow(StatusCodes.success, `Должность изменена!`);
    },
    deleteUser(){
      const list = this.getUserList(this.data.role);
      for(let i = 0; i < list.length; i++){
        if(list[i].id === this.data.id){
          list.splice(i, 1);
          break;
        }
      }
      this.statusWindowStore.showStatusWindow(StatusCodes.success, `Пользователь ${this.data.role === 6 ? 'отсчислен' : 'уволен'}!`);
    }
  }
};
</script>