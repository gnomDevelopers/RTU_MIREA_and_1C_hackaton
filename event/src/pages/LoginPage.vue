<template>
  <div class="flex h-screen">
    <div class="flex flex-col md:w-1/2 uus:w-full h-full items-start justify-start p-4">
      <p class="logo-hr mb-8">EventVUZ+</p>
      <div class="flex flex-col w-full h-full items-center justify-center">
        <div class="flex flex-col w-10/12 h-full items-center m-10">
          <div class="w-full h-full">
            <p class="auth-description">Войдите с помощью учётной записи VUZ+</p>
            <form>
              <div class="mb-8 h-full">
                <loginInput type="text"  text="Логин:" @input-change="checkLogin"/>
                <loginInput type="password" id="student-password" text="Ваш пароль:" @input-change="checkPassword"/>
              </div>
            </form>
            <submitButton value="Войти" class="btn" @click="sendLogin"/>
          </div>
        </div>
      </div>
    </div>
    <div class="login-bg w-1/2 uus:hidden md:block"></div>
  </div>
</template>

<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useUserInfoStore } from '@/stores/userInfoStore';
import loginInput from '../shared/loginInput.vue';
import submitButton from '../shared/submitButton.vue';
import { ValidUserLogin, ValidUserPassword } from '../helpers/validator';
import { type IValidAnswer, StatusCodes, type IAPI_Login_Request } from '../helpers/constants';
import { API_Login } from '@/api/api';
import FilesList from "@/entities/filesList.vue";
import { useUserFormStore } from '@/stores/userFormStore';
import { defineStore } from 'pinia';
import ProfilePage from "@/pages/ProfilePage.vue";


export default {

  components:{
    FilesList,
    loginInput,
    submitButton,
  },

  data(){
    return {
      userType: 'student', // Текущий выбранный тип пользователя
      currentStep: 1, // Шаг формы
      login: {value: '', error: ''} as IValidAnswer,
      password: {value: '', error: ''} as IValidAnswer,
      showPassword: false, // Для отображения пароля
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore, useUserInfoStore),

  },

  methods: {

    sendLogin(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data:IAPI_Login_Request = { email: this.login.value, password: this.password.value };
        API_Login(data)
            .then(response => {
              this.statusWindowStore.deteleStatusWindow(stID);
              this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');

              this.$router.push({name: 'ProfilePage'});
            })
            .catch(error => {
              this.statusWindowStore.deteleStatusWindow(stID);
              if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
              else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
            });
        return;
      }

      if(this.login.value === ''){
        if(this.login.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите логин!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.login.error);
      }
      if(this.password.value === ''){
        if(this.password.error === '')this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!');
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error);
      }

    },
    checkLogin(value: string){
      this.login = ValidUserLogin(value);
      if(value === '') this.login.error = '';
    },
    checkPassword(value: string){
      this.password = ValidUserPassword(value);
      if(value === '') this.password.error = '';
    },



  },
};
</script>

