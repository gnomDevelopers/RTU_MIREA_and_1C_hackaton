<template>
  <div 
    @click="selectThis"
    class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900"
    :class="{'selected-list-item': isSelected}">
    {{ data.name }}
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
export default {
  props: {
    data: {
      type: Object,
      required: true,
    }
  },
  computed: {
    ...mapStores(useScheduleStore),

    isSelected(){
      return this.scheduleStore.selectedSheduleTarget === this.data.name;
    }
  },
  methods: {
    selectThis(){
      this.scheduleStore.selectedSheduleTarget = this.data.name;
    }
  }
};
</script>