<template>
  <section class="main-section">
    <div class="main-div">
      <h3 v-if="isLoggedIn">Todo</h3>
      <UserAuth v-if="!isLoggedIn" @getUserId="getUserId" />

      <div v-if="isLoggedIn" class="submit-div">
        <input @keyup.enter="sendInput" v-model="userInput" type="text" />
        <button @click="sendInput" type="button">Aggiungi</button>
      </div>
      <span v-if="!isSelected && isLoggedIn" class="sub-title-category"
        >Categorie:</span
      >
      <span v-else-if="isSelected && isLoggedIn" class="sub-title-category"
        >Compiti:</span
      >

      <CategoriesComponent
        v-if="!isSelected && isLoggedIn"
        :sendUserInput="sendUserInput"
        @getIdFromComponent="getId"
        :userId="userId"
        @backToAuth="
          backToAuth
        " /><!--Gli if controllano innanzitutto se l'utente è loggato (con isLoggedIn) dopodiché isSelected controlla se ha schiacciato sulla freccia che porta alle task relative alla categoria selezionata--><!--userId arriva dal componente di autenticazione,in pratica è l'id ricevuto in fase di login dello specifico user, con questo id si recuperano tutte le tabelle dello user loggato --><!--getIdFromComponent prende l'id della categoria per poi poterlo assegnare a categoryId e mandarlo alle task, lo fa utilizzando un $emit nel componente Tasks--><!--backToAuth serve semplicemente a fare il logout dello user-->
      <TasksComponent
        v-else-if="isSelected && isLoggedIn"
        :categoryId="categoryId"
        :sendUserInput="sendUserInput"
        @backToCategories="
          backToCategories
        " /><!--Se isSelected è true vengono mostrate le task--><!-- :categoryId si occupa di mandare l'id della categoria selezionata al componente Tasks per usarlo poi durante l'endpoint GET e trovare le task giuste--><!-- :sendUserInput si occupa di trasferire il testo presente nella casella di input alla relativa funzione che si occupa del metodo POST--><!--backToCategories serve a far tornare l'utente nella schermata delle categorie-->
    </div>
  </section>
</template>

<script setup>
import TasksComponent from "./Tasks/TasksComponent.vue";
import CategoriesComponent from "./Categories/CategoriesComponent.vue";
import UserAuth from "./UserAuth/UserAuth.vue";

import { ref } from "vue";

const userInput = ref("");
const sendUserInput = ref("");
const getId = (id) => {
  categoryId.value = id;
  isSelected.value = true; //Se isSelected è true vengono mostrate le task
};
const backToCategories = () => {
  //backToCategories serve a far tornare l'utente nella schermata delle categorie
  isSelected.value = false;
  categoryId.value = null;
};
const backToAuth = () => {
  //backToAuth serve semplicemente a fare il logout dello user
  isLoggedIn.value = false;
  categoryId.value = null;
  userId.value = null;
  isSelected.value = false;
};
const categoryId = ref(null);
const isSelected = ref(false);
const isLoggedIn = ref(false);
const userId = ref(null);

const getUserId = (id) => {
  userId.value = id;
  isLoggedIn.value = true;
};

const sendInput = () => {
  //si occupa di mandare il testo estratto dall'input alla variabile che poi trasferisce al componente tasks il nome della nuova tabella
  sendUserInput.value = userInput.value;
  userInput.value = "";
};
</script>

<style scoped>
.main-section {
  width: 100%;
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  margin: auto;
  margin-top: 180px;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-drag: none;
  position: relative;
}
.main-div {
  height: 600px;
  width: clamp(280px, 50%, 400px);
  background-color: var(--electric-violet);
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  position: relative;
  box-shadow: 0px 0px 20px 5px rgba(0, 0, 0, 0.15);
}
.main-div::after {
  content: "";
  position: absolute;
  bottom: 0;
  height: 60px;
  width: 100%;
  background-color: var(--electric-violet);
  border-radius: 10px;
}
h3 {
  font-size: 2rem;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-drag: none;
  color: white;
}
.add-icon {
  position: absolute;
  height: 70px;
  bottom: 25px;
  cursor: pointer;
  filter: drop-shadow(5px 5px 5px rgba(0, 0, 0, 0.1));
  transition: scale 0.15s ease, transform 2.5s ease;
  -webkit-user-drag: none;
}
.add-icon:hover {
  scale: 1.05;
  filter: brightness(1.5);
  transform: rotate(720deg);
}
.add-icon:active {
  scale: 0.95;
}
.submit-div {
  display: flex;
  column-gap: 10px;
}
.submit-div input {
  box-sizing: border-box;
  padding: 10px 0px;
  border-radius: 20px;
  border: none;
  text-indent: 20px;
  box-shadow: inset 3px 3px 5px rgba(0, 0, 0, 0.15);
  font-family: Poppins;
  font-weight: 500;
  color: lightslategray;
}
.submit-div input:focus {
  outline: none;
}
.submit-div button {
  border-radius: 20px;
  padding: 10px;
  border: none;
  font-family: Poppins;
  font-weight: 500;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.15);
  cursor: pointer;
  transition: scale 0.15s ease, opacity 0.15s ease;
}
.submit-div button:hover {
  scale: 1.05;
  opacity: 0.97;
}
.submit-div button:active {
  scale: 0.95;
  background-color: var(--electric-violet);
  color: white;
}
.sub-title-category {
  position: absolute;
  top: 165px;
  font-size: 20px;
  color: white;
  font-weight: 600;
}

@media (min-width: 250px) and (max-width: 1450px) and (max-height: 600px) {
  .main-section {
    margin-top: 25px;
  }
  .main-div {
    height: 500px;
  }
  .main-div::after {
    height: 20px;
  }
  .submit-div {
    column-gap: 5px;
  }
}
@media (min-width: 2500px) {
  .main-section {
    margin-top: 300px;
  }
  .main-div {
    scale: 1.5;
  }
}
@media (min-height: 600.5px) and (max-height: 860px) {
  .main-section {
    margin-top: 50px;
  }
  .main-div {
    height: 530px;
  }
  .main-div::after {
    height: 20px;
  }
  .submit-div input {
    width: 50%;
    margin-right: auto;
    margin-left: auto;
  }
  .submit-div button {
    width: 35%;
    margin-right: auto;
    margin-left: auto;
  }
  .submit-div {
    column-gap: 5px;
  }
}
</style>
