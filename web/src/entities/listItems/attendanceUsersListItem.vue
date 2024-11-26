<template>
  <div class="flex flex-row items-stretch flex-nowrap min-h-8 rounded-lg bg-white">
    <div class="flex flex-row justify-center flex-shrink-0 items-center w-14 rounded-l-lg bg-color-bold">
      <p class="text-xl text-white">{{ index }}</p>
    </div>
    <div class="flex flex-row items-center px-2 flex-grow">
      <p class="text-xl">{{ data.surname }} {{ data.name }} {{ data.thirdname }}</p>
    </div>
    <div class="flex flex-row self-center gap-x-1 px-2 py-1">
      <div 
        @click="setAttendance(0)" 
        class="w-6 flex flex-row justify-center items-center rounded cursor-pointer unselected-attendace"
        :class="{'selected-attendace': getSelectedAttendaceType === 0}">
        <img class="w-5 h-5" src="../../assets/icons/icon-plus.svg"/>
      </div>
      <div 
        @click="setAttendance(1)" 
        class="w-6 flex flex-row justify-center items-center rounded cursor-pointer unselected-attendace"
        :class="{'selected-attendace': getSelectedAttendaceType === 1}">
        <p class="text-xl text-white">Н</p>
      </div>
      <div 
        @click="setAttendance(2)" 
        class="w-6 flex flex-row justify-center items-center rounded cursor-pointer unselected-attendace"
        :class="{'selected-attendace': getSelectedAttendaceType === 2}">
        <p class="text-xl text-white">У</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useAttendacePageStore } from '@/stores/attendacePageStore';
import { type IGroupAttendance, type IUserGet } from '@/helpers/constants';
import { type PropType } from 'vue';

export default {
  props: {
    index: {
      type: Number,
      required: true,
    },
    data: {
      type: Object as PropType<IUserGet>,
      required: true,
    }
  },
  computed: {
    ...mapStores(useAttendacePageStore),

    getSelectedAttendaceType(){
      for(let item of this.attendancePageStore.attendanceGroupMembers){
        if(item.user.id === this.data.id) return item.attendace;
      }
      return 0;
    }
  },
  methods: {
    setAttendance(attendaceType: number){
      // for(let item of this.attendancePageStore.attendanceGroupMembers){
      //   if(item.user.id === this.data.id) {
      //     item.attendace = attendaceType;
      //     return;
      //   }
      // }
    }
  }
};
</script>