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
  --fire: #d62828;
  --electric-violet: #4e34e0;
}

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
