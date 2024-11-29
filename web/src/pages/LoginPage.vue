<template>
  <div class="flex flex-grow flex-col gap-y-7 items-center justify-center bg-slate-50 login-bg">
    <div class="w-full text-center text-3xl mb:text-4xl">Вход в систему</div>
    <div class="flex flex-col gap-y-6 w-11/12 mb:w-96  py-5 px-7 rounded-lg login-shadow bg-slate-50 text-gray-900">
      <div class="logo">VUZ+</div>
      <div class="flex flex-col gap-y-5">
        <loginInput type="text" text="Почта" @input-change="checkLogin"/>
        <loginInput type="password" text="Пароль" @input-change="checkPassword"/>
        <div class=" text-sm text-sky-500 w-full text-right cursor-pointer">Забыли пароль?</div>
      </div>
      <submitButton value="Войти" class="btn" @click="sendLogin"/>
    </div>
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

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      login: {value: '', error: ''} as IValidAnswer,
      password: {value: '', error: ''} as IValidAnswer,
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore, useUserInfoStore),
  },
  methods: {
    async sendLogin(){
      if(this.login.value !== '' && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data:IAPI_Login_Request = { email: this.login.value, password: this.password.value };

        try{
          const response = await API_Login(data);
          
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация успешна!');

          await this.userInfoStore.onAuthorized(response); // процесс авторизации

          this.$router.push({name: 'MainPage'});//редирект
        }catch (error: any){
          this.statusWindowStore.deteleStatusWindow(stID);

          if(error.status === 500 || error.status === 400) this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Неверный логин или пароль!');
          else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при авторизации!');
        }

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
  }
}

</script>