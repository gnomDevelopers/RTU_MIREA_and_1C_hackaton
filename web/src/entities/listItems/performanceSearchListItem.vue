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
import { type PropType } from 'vue';
import { type IItemList } from '@/helpers/constants';

export default{
  props:{
    data:{
      type: Object as PropType<IItemList>,
      required: true,
    }
  },
  computed: {
    ...mapStores(usePerformancePageStore),

    isSelected(){
      return this.performancePageStore.selectedGroupID === this.data.id;
    }
  },
  methods: {
    selectGroup(){
      this.performancePageStore.selectedGroupID = this.data.id;
      this.performancePageStore.selectedDiscipline = null;
    }
  }
};
</script>