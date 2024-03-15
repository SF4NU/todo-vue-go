<template>
  <section class="main-section">
    <div class="main-div">
      <h3>Todo</h3>
      <CategoriesComponent />
      <img class="add-icon" src="./icons/add.svg" alt="add-new-icon" />
    </div>
    <!-- <TasksComponent /> -->
  </section>
</template>

<script setup>
import TasksComponent from "./Tasks/TasksComponent.vue";
import CategoriesComponent from "./Categories/CategoriesComponent.vue";
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
  width: clamp(280px, 50%, 350px);
  background-color: var(--vanilla);
  display: flex;
  flex-direction: column;
  align-items: center;
  border-radius: 10px;
  position: relative;
  box-shadow: 0px 0px 20px 5px rgba(0, 0, 0, 0.15);
}
h3 {
  font-size: 2rem;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-drag: none;
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
</style>
