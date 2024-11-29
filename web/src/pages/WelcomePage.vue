<template>
  <div class="w-full h-full welcome-bg">
    <div class="flex flex-col justify-center items-center w-full h-full bg-color-light-opacity">
      <div class="flex flex-col gap-y-8 md:gap-y-16 w-full md:w-[720px] opacity-100">
        <p class="welcome-logo" >VUZ+</p>
        <p class="text-center text-3xl mb:text-4xl md:text-5xl leading-normal">Удобная система организации учебного процесса</p>
        <div 
          @click="$router.push({name: 'LoginPage'})" 
          class="box-shadow flex flex-row justify-center items-center self-center py-2 mb:py-3 px-4 mb:px-6 text-2xl mb:text-3xl md:text-4xl rounded-xl text-color-light border-color bg-color-bold cursor-pointer select-none">
          
            Приступить
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">

export default{
  
}

</script>