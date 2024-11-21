<template>
  <div class="flex flex-row gap-x-4 w-full justify-between px-6 text-2xl">
    <div @click="previousMonth" class="cursor-pointer select-none">&lt;</div>
    <div class="cursor-default">{{ getMonthName }}, {{ currentYear }}</div>
    <div @click="nextMonth" class="cursor-pointer select-none">&gt;</div>
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
          <td v-for="day in week" :key="day.day" :class="{'text-gray-400': !day.isThisMonth}">{{ day.day }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script lang="ts">
import { type Day } from '@/helpers/constants';
export default{
  props:{
    month:{
      type: Number,
      required: true,
    }
  },
  data(){
    return{
      currentMonth: [] as Day[][],
      currentMonthIndex: 11,
      currentYear: 2024,
      monthNames: ['Январь','Февраль','Март','Апрель','Май','Июнь','Июль','Август','Сентябрь','Октябрь','Ноябрь','Декабрь'],
    }
  },
  computed:{
    getMonthName(){
      return this.monthNames[this.currentMonthIndex];
    },
  },
  mounted(){
    this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
  },
  methods:{
    nextMonth(){
      this.currentMonthIndex++;
      if(this.currentMonthIndex === 12){
        this.currentMonthIndex = 0;
        this.currentYear++;
      }
      this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
    },
    previousMonth(){
      this.currentMonthIndex--;
      if(this.currentMonthIndex === -1){
        this.currentMonthIndex = 11;
        this.currentYear--;
      }
      this.currentMonth = this.loadMonth(this.currentYear, this.currentMonthIndex);
    },
    loadMonth(year: number, month: number): Day[][] { // месяц начинется с 0
      const firstDay = new Date(year, month).getDay();
      const monthSize = new Date(year, month, 0).getDate();
      const diffStart = (firstDay + 7 - 1) % 7 - 1;

      const weeks: Day[][] = [];
      for(let j = 0; j < 6; j++){

        const curWeek:Day[] = [];
        for(let i = 0; i < 7; i++){
          const d = j * 7 + i;
          const day = new Date(year, month, d - diffStart);

          let dayMonth = 0;
          if(d <= diffStart) dayMonth = (month - 1 + 12) % 12;
          else {
            if(d > diffStart + monthSize) dayMonth = (month + 1) % 12;
            else dayMonth = month;
          }

          curWeek.push({year: day.getFullYear(), month: dayMonth, day: day.getDate(), isThisMonth: month === dayMonth});
        }

        weeks.push(curWeek);
      }
      
      return weeks;
    }
  }
};
</script>