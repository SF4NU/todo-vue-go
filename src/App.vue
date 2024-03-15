<template>
  <div>
    <MainComponent />
  </div>
</template>

<script setup>
import MainComponent from "./components/MainComponent.vue";
</script>

<style>
:root {
  --vanilla: #eae2b7;
  --xanthous: #fcbf49;
  --orange: #f77f00;
  --fire: #d62828;
  --prussian: #003049;
}

/* * {
  outline: 1px solid red;
} */

body {
  margin: 0;
  font-family: NorthwestBold;
  background: rgb(209, 209, 209);
  background: radial-gradient(
    circle,
    rgba(209, 209, 209, 1) 45%,
    rgba(144, 144, 144, 1) 100%
  );
  font-family: Poppins, Arial;
}
</style>
