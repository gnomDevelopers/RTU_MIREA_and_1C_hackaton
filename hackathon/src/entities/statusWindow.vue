<template>
  <div class="absolute self-center top-16 flex flex-col gap-y-1 z-50">
    <div
      v-for="item in statusWindowStore.statusWindowList"
      :key="item.id"
      class="flex flex-row gap-x-2 items-center py-1 px-3 rounded-md"
      :class="{
        'bg-amber-500': isLoading(item.status),
        'bg-green-600': isSuccess(item.status),
        'bg-red-500': isError(item.status),
        'bg-cyan-600': isInfo(item.status),
      }"
    >
      <div>
        <img
          v-if="isError(item.status)"
          class="w-7 h-7"
          src="../assets/icons/icon-error.svg"
        />
        <img
          v-else-if="isSuccess(item.status)"
          class="w-7 h-7"
          src="../assets/icons/icon-success.svg"
        />
        <img
          v-else-if="isLoading(item.status)"
          class="w-7 h-7"
          src="../assets/icons/icon-loading.svg"
        />
        <img
          v-else-if="isInfo(item.status)"
          class="w-7 h-7"
          src="../assets/icons/icon-info.svg"
        />
      </div>
      <div class="text-slate-50">
        {{ item.text }}
      </div>
    </div>
  </div>
</template>
<script lang="ts">

import { mapStores } from "pinia";
import { useStatusWindowStore, statusCodes } from "../stores/statusWindowStore";

export default {
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods: {
    isLoading(status: statusCodes): Boolean{
      return status === statusCodes.loading;
    },
    isSuccess(status: statusCodes): Boolean{
      return status === statusCodes.success;
    },
    isError(status: statusCodes): Boolean{
      return status === statusCodes.error;
    },
    isInfo(status: statusCodes): Boolean{
      return status === statusCodes.info;
    },
  }
};
</script>