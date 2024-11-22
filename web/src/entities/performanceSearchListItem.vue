<template>
  <div class="p-2 text-lg bg-color-light rounded-xl cursor-pointer transition-colors border border-solid border-transparent hover:border-blue-900">
    {{ data.name }}
  </div>
</template>
<script lang="ts">
import { type PropType } from 'vue';
import { type IItemList } from '@/helpers/constants';

export default{
  props:{
    data:{
      type: Object as PropType<IItemList>,
      required: true,
    }
  }
};
</script>