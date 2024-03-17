<template>
  <div class="user-auth-div">
    <div class="form-div-1">
      <h1>Registrati</h1>
      <form class="form-register" @submit.prevent="">
        <input v-model="userNameReg" type="text" placeholder="Nome utente" />
        <input
          v-model="passwordReg"
          type="password"
          placeholder="Password" /><button
          ref="registerButton"
          @click="createUser">
          {{ signUp }}
        </button>
      </form>
    </div>

    <div class="form-div-2">
      <h1>Accedi</h1>
      <form action="" class="form-login" @submit.prevent="">
        <input v-model="userNameLog" type="text" placeholder="Nome utente" />
        <input
          v-model="passwordLog"
          type="password"
          placeholder="Password" /><button
          ref="loginButton"
          @click="login($emit)">
          {{ signIn }}
        </button>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import axios from "axios";
import { wait } from "@/shared/wait";

const userNameReg = ref("");
const passwordReg = ref("");
const userNameLog = ref("");
const passwordLog = ref("");
const signIn = ref("Sign-In");
const signUp = ref("Sign-Up");
const loginButton = ref(null);
const registerButton = ref(null);

const createUser = async () => {
  if (userNameReg.value !== "" && passwordReg.value !== "") {
    try {
      const res = await axios.post(
        "https://go-fiber-prova-production.up.railway.app/registration",
        {
          username: userNameReg.value,
          password: passwordReg.value,
        }
      );
      userNameReg.value = "";
      passwordReg.value = "";

      if (res.status >= 200 || res.status <= 209) {
        registerButton.value.style.backgroundColor = "#3BD571";
        registerButton.value.style.color = "black";
        signUp.value = "Account Creato";
      }
      await wait(3000);
      registerButton.value.style.backgroundColor = "white";
      registerButton.value.style.color = "black";
      signUp.value = "Sign-up";
    } catch (error) {
      console.error(error);
      userNameReg.value = "";
      passwordReg.value = "";
      registerButton.value.style.backgroundColor = "red";
      registerButton.value.style.color = "white";
      signUp.value = "Nome utente già esistente!";
      await wait(3000);
      registerButton.value.style.backgroundColor = "white";
      registerButton.value.style.color = "black";
      signUp.value = "Sign-up";
    }
  }
};
const login = async (emit) => {
  if (userNameLog.value !== "" && passwordLog.value !== "") {
    try {
      const res = await axios.post(
        "https://go-fiber-prova-production.up.railway.app/login",
        {
          username: userNameLog.value,
          password: passwordLog.value,
        }
      );
      if (res.status >= 200 || res.status <= 209) {
        emit("getUserId", res.data.user_id);
        loginButton.value.style.backgroundColor = "#3BD571";
        loginButton.value.style.color = "black";
        signIn.value = "Login effettuato!";
      }
      userNameLog.value = "";
      passwordLog.value = "";
      await wait(3000);
      loginButton.value.style.backgroundColor = "white";
      loginButton.value.style.color = "black";
      signIn.value = "Sign-In";
    } catch (error) {
      console.error(error);
      userNameLog.value = "";
      passwordLog.value = "";
      loginButton.value.style.backgroundColor = "red";
      loginButton.value.style.color = "white";
      signIn.value = "Credenziali sbagliate!";
      await wait(3000);
      loginButton.value.style.backgroundColor = "white";
      loginButton.value.style.color = "black";
      signIn.value = "Sign-In";
    }
  }
};
</script>

<style scoped>
.user-auth-div {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  width: 100%;
  color: white;
}
.form-div-1,
.form-div-2 {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.form-login,
.form-register {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  row-gap: 10px;
}
.user-auth-div input {
  width: 80%;
  padding: 10px 10px;
  border-radius: 40px;
  border: none;
  font-family: Poppins, Arial;
  text-indent: 10px;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.15);
  transition: border-radius 0.4s ease;
}
.user-auth-div input:focus {
  border-radius: 0px;
}
.user-auth-div input:focus {
  outline: none;
}
.user-auth-div button {
  border: none;
  padding: 10px 15px;
  background-color: aliceblue;
  border-radius: 25px;
  box-shadow: inset 5px 5px 5px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: opacity 0.15s ease, transform 0.15s ease, scale 0.15s ease;
  font-family: Poppins, Arial;
}
.user-auth-div button:hover {
  opacity: 0.95;
  transform: translate(-0.5px, 0.5px);
}
.user-auth-div button:active {
  scale: 0.8;
}
</style>
