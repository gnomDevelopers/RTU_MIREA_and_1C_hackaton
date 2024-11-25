<template>
  <div class="flex flex-col items-center w-full mb:w-auto gap-y-4 bg-color-light rounded-xl p-4">
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">25.11.2024, Понедельник</div>
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">13 неделя</div>
    </div>
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <ScheduleClassListItem 
        v-for="(item, ind) in scheduleStore.scheduleTableDay"
        :key="uniqueID++"
        :index="ind + 1" 
        :time="`${item.time} ${item.type}`" 
        :title="item.title" 
        :room="item.place" 
        :group="item.groups.join(', ')"
        :canAddFaculties 
      />
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
import ScheduleClassListItem from '@/shared/scheduleClassListItem.vue';

export default {
  props:{
    canAddFaculties:{
      type:Boolean,
      required: false,
      default: false,
    }
  },
  components:{
    ScheduleClassListItem,
  },
  computed:{
    ...mapStores(useScheduleStore),
  },
  data(){
    return {
      uniqueID: 1,
    }
  }
};
</script>