<template class="template">
  <section class="category-section">
    <div class="category-div">
      <div v-for="(category, i) in data" :key="i" class="sub-div-category">
        <span v-if="checkInput" class="category-name">{{
          category.title
        }}</span>
        <!-- <input
          v-else-if="!checkInput"
          type="text"
          v-model="category"
          @keyup.enter="checkInput = !checkInput" /> -->
        <div class="div-options">
          <img class="options-icons" src="../icons/edit.svg" alt="" />
          <img
            @click="deleteCategory(category.ID)"
            class="options-icons"
            src="../icons/redbin.svg"
            alt="" />
          <img class="options-icons" src="../icons/arrow.svg" alt="" />
        </div>
        <span class="edit-time">Ultima modifica: 00:00</span>
      </div>
    </div>
  </section>
</template>

<script setup>
import axios from "axios";
import { ref, onBeforeMount, onUnmounted, watch } from "vue";

let checkInput = ref(true);
let category = ref("");
let data = ref(null);
let watchVar = ref(0);

const deleteCategory = (id) => {
  console.log(id);
  const deleteTheCategory = async () => {
    try {
      await axios.delete(
        `https://go-fiber-prova-production.up.railway.app/categories/${id}`
      );
      watchVar.value++;
      console.log(watchVar);
      data.value = data.value.filter((category) => category.ID !== id);
    } catch (error) {
      console.error(error);
    }
  };
  deleteTheCategory();
};
const fetchData = async () => {
  try {
    const res = await axios.get(
      "https://go-fiber-prova-production.up.railway.app/categories"
    );
    data.value = res.data;
    console.log(res.data);
  } catch (error) {
    console.error(error);
  }
};
onBeforeMount(() => {
  fetchData();
});
onUnmounted(() => {
  data.value = null;
});

watch(
  () => watchVar.value,
  () => {
    console.log("changed");
    fetchData();
  }
);
</script>

<style scoped>
.category-div {
  width: 90%;
  display: flex;
  flex-direction: column;
  row-gap: 10px;
  overflow-y: auto;
  scrollbar-width: none;
}
.category-div:last-child {
  padding-bottom: 10px;
}
.sub-div-category {
  background-color: white;
  box-shadow: 0px 5px 5px rgba(0, 0, 0, 0.1);
  min-height: 70px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  text-indent: 15px;
  border-left: 10px var(--fire) solid;
  font-size: 18px;
  scroll-snap-align: start;
  font-weight: 500;
  position: relative;
}
.sub-div-category input {
  border: none;
  text-indent: 13px;
  font-family: Poppins;
  color: black;
  font-weight: 500;
  font-size: 18px;
}
.sub-div-category input:focus {
  outline: none;
}
.category-section {
  width: 100%;
  height: 340px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
}
.div-options {
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  position: absolute;
  right: 20px;
}
.options-icons {
  height: 30px;
  cursor: pointer;
}
.edit-time {
  font-size: 10px;
  font-weight: 400;
  color: rgb(153, 153, 153);
  position: absolute;
  right: 20px;
  bottom: 3px;
}
</style>
