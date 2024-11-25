<template>
  <div 
    @click="selectGroup"
    class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900"
    :class="{'selected-list-item': isSelected}">
    {{ data.name }}
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useGroupCorrectPageStore } from '@/stores/groupCorrectPageStore';
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
    ...mapStores(useGroupCorrectPageStore),

    isSelected(){
      return this.groupCorrectPageStore.selectedGroupID === this.data.id;
    }
  },
  methods: {
    selectGroup(){
      this.groupCorrectPageStore.selectedGroupID = this.data.id;
    }
  }
};
</script>