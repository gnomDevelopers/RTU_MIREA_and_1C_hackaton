<template>
  <div class="flex flex-col sm:flex-row items-start sm:items-center p-2 rounded-lg bg-color-light">
    <div class="flex-grow">
      <p class=" select-none">{{ data.name }}</p>
    </div>
    <div class="flex flex-row gap-x-2 flex-shrink-0 ">
      <div class="py-1 px-2 rounded-lg cursor-pointer text-sm btn">{{ getButtonText }}</div>
      <div class="py-1 px-2 rounded-lg cursor-pointer text-sm btn btn-bad">Уволить</div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useUserInfoStore } from '@/stores/userInfoStore';
import { ROLES_NAME, ROLES_SET_PRORECTOR, ROLES_SET_DECAN_TO_PREPOD, ROLES_SET_DECAN_TO_STUDENT, type IUsersList } from '@/helpers/constants';
import type { PropType } from 'vue';
export default{
  props:{
    data:{
      type: Object as PropType<IUsersList>,
      required: true,
    }
  },
  computed:{
    ...mapStores(useUserInfoStore),

    getButtonText(){
      let rolesSum;
      if(this.userInfoStore.role === 1){
        rolesSum = ROLES_SET_PRORECTOR.reduce((sum, item) => sum += item);
        if(this.data.role === 2) return 'Понизить до ' + ROLES_NAME[rolesSum - this.data.role];
        if(this.data.role === 4) return 'Повысить до ' + ROLES_NAME[rolesSum - this.data.role]; 
      }
      if(this.userInfoStore.role === 2){
        if(ROLES_SET_DECAN_TO_PREPOD.includes(this.data.role)){
          rolesSum = ROLES_SET_DECAN_TO_PREPOD.reduce((sum, item) => sum += item);
          if(this.data.role === 3) return 'Понизить до ' + ROLES_NAME[rolesSum - this.data.role];
          if(this.data.role === 4) return 'Повысить до ' + ROLES_NAME[rolesSum - this.data.role]; 
        }
        else{
          rolesSum = ROLES_SET_DECAN_TO_STUDENT.reduce((sum, item) => sum += item);
          if(this.data.role === 5) return 'Понизить до ' + ROLES_NAME[rolesSum - this.data.role];
          if(this.data.role === 6) return 'Повысить до ' + ROLES_NAME[rolesSum - this.data.role]; 
        }
      }
      if(this.userInfoStore.role === 3){
        // 
      }
      return '';
    }
  },
};
</script>