<template>
  <div 
    @click="selectGroup" 
    class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900"
    :class="{'selected-list-item' : isSelected}"
    >
    
    {{ data.name }}
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { usePerformancePageStore } from '@/stores/performancePageStore';

export default{
  props:{
    data:{
      type: Object, // {name: ...}
      required: true,
    }
  },
  computed: {
    ...mapStores(usePerformancePageStore),

    isSelected(){
      return this.performancePageStore.selectedGroup === this.data.name;
    }
  },
  methods: {
    selectGroup(){
      this.performancePageStore.selectedGroup = this.data.name;
      this.performancePageStore.selectedDiscipline = null;
    }
  }
};
</script>