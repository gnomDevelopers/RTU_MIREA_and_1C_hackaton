<template>
  <div class="flex flex-col sm:flex-row items-start sm:items-center p-2 rounded-lg bg-color-light">
    <div class="flex-grow">
      <p class="text-lg">{{ data.surname }} {{ data.name }} {{ data.thirdname }}</p>
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
      <div v-if="userRole !== data.role" class="py-1 px-2 rounded-lg cursor-pointer text-base select-none btn">Сохранить</div>
      <div class="py-1 px-2 rounded-lg cursor-pointer text-base select-none btn btn-bad">{{ data.role !== 6 ? 'Уволить' : 'Отчислить' }}</div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { ROLES_NAME, ROLES_SET_PRORECTOR, ROLES_SET_DECAN_TO_PREPOD, type IUserGet } from '@/helpers/constants';
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
    ...mapStores(useUserInfoStore),

  },
  methods: {
    getRolesName(roleID: number): string{
      return ROLES_NAME[roleID];
    }
  }
};
</script>