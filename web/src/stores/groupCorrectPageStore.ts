import { defineStore } from "pinia";
import{ type TMaybeNumber } from "@/helpers/constants";

export const useGroupCorrectPageStore = defineStore('groupCorrectPage', {
  state() {
    return{
      selectedGroupID: null as TMaybeNumber,
    }
  }
});