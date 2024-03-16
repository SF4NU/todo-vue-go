<template>
  <section class="main-section">
    <div class="main-div">
      <h3>Todo</h3>

      <div class="submit-div">
        <input @keyup.enter="sendInput" v-model="userInput" type="text" />
        <button @click="sendInput" type="button">Aggiungi</button>
      </div>
      <span v-if="!isSelected" class="sub-title-category">Categorie:</span>
      <span v-else-if="isSelected" class="sub-title-category">Compiti:</span>

      <CategoriesComponent
        v-if="!isSelected"
        :sendUserInput="sendUserInput"
        @getIdFromComponent="getId" />
      <TasksComponent
        v-else-if="isSelected"
        :categoryId="categoryId"
        :sendUserInput="sendUserInput"
        @backToCategories="backToCategories" />
    </div>
  </section>
</template>

<script setup>
import TasksComponent from "./Tasks/TasksComponent.vue";
import CategoriesComponent from "./Categories/CategoriesComponent.vue";
import { ref, watch } from "vue";

const userInput = ref("");
const sendUserInput = ref("");
const getId = (id) => {
  categoryId.value = id;
  isSelected.value = true;
};
const backToCategories = () => {
  isSelected.value = false;
  categoryId.value = null;
};
const categoryId = ref(null);
const isSelected = ref(false);

const sendInput = () => {
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
</style>
