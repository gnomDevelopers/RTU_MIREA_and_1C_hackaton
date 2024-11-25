import { defineStore } from "pinia";
import{ type TMaybeNumber } from "@/helpers/constants";

export const usePerformancePageStore = defineStore('performancePage', {
  state() {
    return{
      selectedGroupID: null as TMaybeNumber,
      selectedDiscipline: null as TMaybeNumber,
    }
  }
});