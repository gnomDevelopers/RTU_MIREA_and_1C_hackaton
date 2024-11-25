<template>
  <div class="flex flex-row justify-start items-stretch min-h-14">
    <p class="text-xl font-medium w-14 flex flex-row flex-shrink-0 justify-center items-center rounded-l-xl cursor-default bg-color-bold text-white">{{ index }}</p>
    <div @click="selectClass" class="flex flex-row flex-grow items-stretch bg-white rounded-r-xl px-2 cursor-pointer">
      
      <div class="flex flex-col flex-grow items-start gap-y-1">
        <p class="flex flex-row flex-wrap gap-x-1 text-lg">
          <span class="font-bold">{{ time }}</span>
          {{ title }}
        </p>
        <p class="flex flex-row flex-wrap gap-x-1 text-base">
          <span class="font-bold">{{ room }}</span>
          {{ group }}
        </p>
      </div>

      <div class="self-center cursor-pointer" v-if="title === '' && canAddFaculties">
        <svg class="w-6 h-6" viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M16 2V30M2 16H30" stroke="#063C73" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"/>
        </svg>
      </div>

    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useScheduleStore } from '@/stores/scheduleStore';
export default{
  props: {
    index: {
      type: Number,
      required: true,
    },
    time: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    room: {
      type: String,
      required: true,
    },
    group: {
      type: String,
      required: true,
    },
    canAddFaculties:{
      type:Boolean,
      required: false,
      default: false,
    }
  },
  computed: {
    ...mapStores(useScheduleStore),
  },
  methods: {
    selectClass(){
      this.scheduleStore.selectedClass = 1; //
    }
  }
};
</script>