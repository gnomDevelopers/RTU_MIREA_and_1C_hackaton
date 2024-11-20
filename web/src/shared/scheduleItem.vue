<template>
  <div class="flex flex-row justify-center items-stretch min-h-14">
    <p class="text-xl font-medium w-14 flex flex-row flex-shrink-0 justify-center items-center rounded-l-xl cursor-default bg-color-bold text-white">{{ index }}</p>
    <div class="flex flex-col flex-grow items-start gap-y-1 bg-white rounded-r-xl px-2">
      <p class="flex flex-row flex-wrap gap-x-1 text-lg">
        <span class="font-bold">{{ time }}</span>
        {{ title }}
      </p>
      <p class="flex flex-row flex-wrap gap-x-1 text-base">
        <span class="font-bold">{{ room }}</span>
        {{ group }}
      </p>
    </div>
  </div>
</template>
<script lang="ts">
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
  },

};
</script>