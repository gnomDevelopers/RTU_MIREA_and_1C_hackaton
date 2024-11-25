<template>
  <div 
    @click="selectDiscipline"
    class="flex flex-row items-stretch min-h-14 cursor-pointer border-2 border-solid rounded-xl overflow-hidden transition-colors border-transparent hover:border-blue-900"
    :class="{'selected-list-item' : isSelected}">
    <div class="flex flex-col justify-center items-center w-10  bg-color-bold">
      <p class="text-xl text-color-light">{{ data.id }}</p>
    </div>
    <div class="flex flex-col flex-grow justify-center px-2 bg-white ">
      <p class="text-lg">{{ data.name }}</p>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { usePerformancePageStore } from '@/stores/performancePageStore';
import { type ISubject } from '@/helpers/constants';
import { type PropType } from 'vue';

export default{
  props: {
    data: {
      type: Object as PropType<ISubject>,
      required: true,
    }
  },
  computed: {
    ...mapStores(usePerformancePageStore),

    isSelected(){
      return this.performancePageStore.selectedDiscipline === this.data.id;
    }
  },
  methods: {
    selectDiscipline(){
      this.performancePageStore.selectedDiscipline = this.data.id;
    }
  }
};
</script>