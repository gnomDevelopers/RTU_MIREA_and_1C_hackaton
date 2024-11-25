<template>
  <div 
    @click="selectGroup"
    class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900"
    :class="{'selected-list-item' : isSelected}">
    {{ data.name }}
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useAttendacePageStore } from '@/stores/attendacePageStore';
import { useScheduleStore } from '@/stores/scheduleStore';
import { type IItemList } from '@/helpers/constants';
import { type PropType } from 'vue';

export default {
  props:{
    data: {
      type: Object as PropType<IItemList>,
      required: true,
    }
  },
  computed: {
    ...mapStores(useAttendacePageStore, useScheduleStore),

    isSelected(){
      return this.attendancePageStore.selectedGroupID === this.data.id;
    }
  },
  methods: {
    selectGroup(){
      this.attendancePageStore.selectedGroupID = this.data.id;
      this.scheduleStore.selectedClass = null;
    }
  }
};
</script>