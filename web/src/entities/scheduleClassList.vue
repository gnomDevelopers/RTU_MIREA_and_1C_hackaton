<template>
  <div class="flex flex-col items-center gap-y-4 bg-color-light rounded-xl p-4">
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">12.11.2024, Вторник</div>
      <div class="text-center cursor-default text-xl bg-white rounded-lg py-1">11 неделя</div>
    </div>
    <div class="flex flex-col gap-y-1 w-full items-stretch">
      <ScheduleClassListItem 
        v-for="(item, ind) in scheduleStore.scheduleTableDay"
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
  }
};
</script>