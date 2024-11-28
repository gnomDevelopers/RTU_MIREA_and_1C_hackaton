<template>
  <div>
    <div class="flex flex-row gap-x-4 px-6 justify-between items-center text-2xl">
      <div @click="previousMonth" class="cursor-pointer select-none text-3xl p-1">&lt;</div>
      <div class="cursor-default">{{ getMonthName }}, {{ currentYear }}</div>
      <div @click="nextMonth" class="cursor-pointer select-none text-3xl p-1">&gt;</div>
    </div>
    <div class="text-xl p-4 pt-0 rounded-xl bg-color-light">
      <table class="schedule">
        <thead>
          <tr>
            <td>Пн</td>
            <td>Вт</td>
            <td>Ср</td>
            <td>Чт</td>
            <td>Пт</td>
            <td>Сб</td>
            <td>Вс</td>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(week, index) in currentMonth" :key="index">
            <td v-for="day in week" :key="day.day" :class="{'text-gray-400': day.month !== currentMonthIndex, 'selected-day' : isSameDay(day)}" @click="selectDay(day)">{{ day.day }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script lang="ts">
import { type Day, MONTH_NAMES } from '@/helpers/constants';

const today = new Date();

export default{
  emits:['selectDay'],
  props:{
    month:{
      type: Number,
      required: true,
    }
  },
  data(){
    return{
      currentMonth: [] as Day[][],
      currentMonthIndex: this.month,
      currentYear: 2024,

      selectedDay: {year: today.getFullYear(), month: today.getMonth(), day: today.getDate()} as Day,
    }
  },
  computed:{
    getMonthName(){
      return MONTH_NAMES[this.currentMonthIndex];
    },
  },
  mounted(){
    this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
  },
  methods:{
    nextMonth(){
      if(this.currentYear > new Date().getFullYear() + 1) return; // ограничения по году
      this.currentMonthIndex++;
      if(this.currentMonthIndex === 12){
        this.currentMonthIndex = 0;
        this.currentYear++;
      }
      this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
    },
    previousMonth(){
      if(this.currentYear < new Date().getFullYear() - 1) return; // ограничения по году
      this.currentMonthIndex--;
      if(this.currentMonthIndex === -1){
        this.currentMonthIndex = 11;
        this.currentYear--;
      }
      this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
    },
    loadMonth(year: number, month: number): Day[][] { // месяц начинется с 0
      const firstDay = new Date(year, month).getDay();
      const monthSize = new Date(year, month+1, 0).getDate();
      const diffStart = (firstDay + 7 - 1) % 7;

      const weeks: Day[][] = [];
      for(let j = 0; j < 6; j++){

        const curWeek:Day[] = [];
        for(let i = 1; i <= 7; i++){
          const d = j * 7 + i;
          const day = new Date(year, month, d - diffStart);

          let dayMonth = 0;
          if(d <= diffStart) dayMonth = (month - 1 + 12) % 12;
          else {
            if(d > diffStart + monthSize) dayMonth = (month + 1) % 12;
            else dayMonth = month;
          }

          curWeek.push({year: day.getFullYear(), month: dayMonth, day: day.getDate()});
        }
        weeks.push(curWeek);
      }
      
      return weeks;
    },
    selectDay(day: Day){
      this.selectedDay = day;
      this.$emit('selectDay', day);
    },
    isSameDay(day: Day){
      if(
        this.selectedDay.year === day.year && 
        this.selectedDay.month === day.month && 
        this.selectedDay.day === day.day) 
        return true;
      return false;
    }
  }
};
</script>