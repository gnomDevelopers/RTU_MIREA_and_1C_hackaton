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

export default {
  props:{
    data: {
      type: Object,
      required: true,
    }
  },
  computed: {
    ...mapStores(useAttendacePageStore, useScheduleStore),

    isSelected(){
      return this.attendancePageStore.selectedGroup === this.data.name;
    }
  },
  methods: {
    selectGroup(){
      this.attendancePageStore.selectedGroup = this.data.name;
      this.scheduleStore.selectedClass = null;
    }
  }
};
</script>