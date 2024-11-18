<template>
  <div class="flex flex-col w-full h-full bg-color-ultralight scrollable">
    <div class="flex flex-col justify-evenly items-center gap-y-8 pt-8 md:gap-y-0 md:flex-row md:items-stretch md:pt-16">
      <div class="flex flex-col gap-y-4 md:gap-y-16 items-stretch">
        <p class="text-center text-3xl">Личный Кабинет</p>
        <div class="flex flex-col items-center">
          <img class="w-32 h-32" src="../assets/icons/icon-profile.svg"/>
          <p class="text-2xl mt-10">Админ Админов</p>
          <p class="text-lg text-sky-600">adminov_a_d</p>
        </div>
        <div class="flex flex-col gap-y-3 items-stretch py-3 px-6 bg-color-light rounded-lg profile-shadow">
          <div class="py-2 px-12 text-center rounded-lg cursor-pointer text-sky-500 bg-white hover:underline header-shadow active:text-sky-300">Сменить пароль</div>
          <div class="py-2 px-12 text-center rounded-lg cursor-pointer text-red-700 bg-white hover:underline header-shadow active:text-red-400">Выйти из аккаунта</div>
        </div>
      </div>
      <div class="flex flex-col gap-y-4 md:gap-y-16 pb-10">
        <p class="text-center text-3xl">Личная информация</p>
        <div class="flex justify-center items-center p-10 aspect-square rounded-xl bg-color-light profile-shadow">
          <div class="p-5 rounded-lg aspect-square bg-white header-shadow">
            <ul class="list-disc markers profile-list">
              <li>Вуз: <span>РТУ МИРЭА</span></li>
              <li>Должность: <span>Администратор</span></li>
              <li>ФИО: <span>Админов Админ Админович</span></li>
              <li>Институт: <span>ИПТИП</span></li>
              <li>Кафедра: <span>ИП</span></li>
              <li>Группа: <span>ЭФБО-01-23</span></li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">

export default{

};

</script>