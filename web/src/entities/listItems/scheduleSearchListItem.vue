<template>
  <div 
    @click="selectThis"
    class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900"
    :class="{'selected-list-item': isSelected}">
    {{ getName() }}
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
      return this.scheduleStore.selectedSheduleGroup === this.data.id;
    }
  },
  methods: {
    getName() {
      return `${this.data.surname === undefined ? '' : this.data.surname} ${this.data.name} ${this.data.thirdname === undefined ? '' : this.data.thirdname}`;
    },
    selectThis(){
      this.scheduleStore.selectedSheduleGroup = this.data.id;
    }
  }
};
</script>