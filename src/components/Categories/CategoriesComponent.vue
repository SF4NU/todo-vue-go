<template class="template">
  <div ref="scrollToTheBottom" class="category-div">
    <div
      v-for="(category, i) in data"
      :key="category.ID"
      :class="`sub-div-category`">
      <div v-if="isEditing !== i" class="sub-div-wrapper">
        <span class="category-name">{{ checkLength(category.title, 17) }}</span>
        <div class="div-options">
          <img
            @click="isEditingCategory(i)"
            class="options-icons edit-icon"
            src="../icons/edit.svg"
            alt="" />
          <img
            @click="askToDelete(i)"
            class="options-icons remove-category-icon"
            src="../icons/redbin.svg"
            alt="" />
          <img
            @click="$emit('getIdFromComponent', category.ID)"
            class="options-icons see-task-icon"
            src="../icons/arrow.svg"
            alt="" />
        </div>
        <span class="edit-time"
          >Ultima modifica: {{ category.last_modified }}</span
        >
      </div>
      <div v-else-if="isEditing === i && !isRemoving" class="edit-div">
        <input class="edit-input" v-model="updatedTitle" type="text" />
        <button class="edit-button" @click="editFinished(i, category.ID)">
          Aggiorna
        </button>
      </div>
      <div v-else class="remove-div">
        <div>Stai per rimuovere "{{ checkLength(category.title) }}"</div>
        <div>
          <img
            @click="deleteCategory(category.ID, i)"
            class="tick-remove tick-icon"
            src="../icons/tick.svg"
            alt="" />
          <img
            @click="doNotDelete()"
            class="tick-remove remove-icon"
            src="../icons/remove.svg"
            alt="" />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import axios from "axios";
import { ref, onBeforeMount, onUnmounted, watch } from "vue";
import { getCurrentTime } from "../../shared/getCurrentTime";
import { checkLength } from "../../shared/checkLength";

const userProps = defineProps(["sendUserInput"]);
let isEditing = ref(null);
let data = ref(null);
let watchVar = ref(0);
const scrollToTheBottom = ref(null);
let updatedTitle = ref("");
let isRemoving = ref(false);

let timeoutAddId = null;

const isEditingCategory = (id) => {
  isEditing.value = id;
};
const editFinished = (index, id) => {
  updateCategory(index, id);
  isEditing.value = null;
};

const updateCategory = async (index, id) => {
  if (updatedTitle.value !== "") {
    try {
      const updateCategoryTitle = {
        title: updatedTitle.value,
        last_modified: getCurrentTime(),
      };
      data.value[index] = updateCategoryTitle;

      await axios.put(
        `https://go-fiber-prova-production.up.railway.app/categories/${id}`,
        {
          title: updatedTitle.value,
          last_modified: getCurrentTime(),
        }
      );
      updatedTitle.value = "";
      fetchData();
    } catch (error) {
      console.error(error);
    }
  }
};

const addCategory = async () => {
  if (userProps.sendUserInput !== "" && userProps.sendUserInput !== " ") {
    try {
      scroll();

      const newCategoryTitle = {
        title: userProps.sendUserInput,
        last_modified: getCurrentTime(),
      };
      data.value.push(newCategoryTitle);
      const res = await axios.post(
        `https://go-fiber-prova-production.up.railway.app/categories`,
        {
          title: userProps.sendUserInput,
          last_modified: getCurrentTime(),
        }
      );
      timeoutAddId ? clearTimeout(timeoutAddId) : null;
      timeoutAddId = await setTimeout(() => {
        //aspetta prima di aggiornare di rimandare la richiesta GET nel caso l'utente stia ancora aggiungendo categorie
        fetchData();
      }, 5000);
    } catch (error) {
      console.error(error);
    }
  }
};

const scroll = () => {
  scrollToTheBottom.value.scrollTo({
    top: scrollToTheBottom.value.scrollHeight,
    behavior: "smooth",
  });
};

watch(
  () => userProps.sendUserInput,
  () => {
    addCategory();
  }
);

const askToDelete = (i) => {
  isEditing.value = i;
  isRemoving.value = true;
};

const doNotDelete = () => {
  isEditing.value = null;
  isRemoving.value = false;
};

let timeoutId = null;

const deleteCategory = async (id, index) => {
  // elimina la categoria selezionata
  try {
    isRemoving.value = false;
    isEditing.value = null;
    data.value = data.value.filter((_, i) => i !== index); //dato che ci vogliono 1-2 secondi per mandare request al server e riceverne un'altra ho pensato che è meglio aggiornare prima quello che vede l'user e lasciare fare le richieste in background così da non rallentare nessuno
    await axios.delete(
      `https://go-fiber-prova-production.up.railway.app/categories/${id}`
    );
    timeoutId ? clearTimeout(timeoutId) : null; // se viene schiacciato di nuovo il pulsante "deleteCategory" riparte il timer per evitare sovrapposizioni
    timeoutId = await setTimeout(() => {
      watchVar.value++; // quando questa variabile cambia riparte la funzione fetchData, il timer serve per quando vengono cancellate molto velocemente le categorie e quindi si evita il flash delle categorie che appaiono ogni volta che cambia watchVar
    }, 5000);
  } catch (error) {
    console.error(error);
  }
};
const fetchData = async () => {
  try {
    const res = await axios.get(
      "https://go-fiber-prova-production.up.railway.app/categories"
    );
    data.value = res.data.sort((a, b) => a.ID - b.ID);
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
    fetchData();
  }
);
</script>

<style>
.category-div {
  width: 90%;
  display: flex;
  flex-direction: column;
  row-gap: 10px;
  overflow-y: auto;
  scrollbar-width: none;
  margin-top: 60px;
}
.category-div:last-child {
  padding-bottom: 200px;
}
.sub-div-category {
  background-color: white;
  box-shadow: 0px 5px 5px rgba(0, 0, 0, 0.1);
  min-height: 70px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  text-indent: 10px;
  font-size: 15px;
  scroll-snap-align: start;
  font-weight: 500;
  position: relative;
}
.sub-div-category input {
  border: none;
  text-indent: 10px;
  font-family: Poppins;
  color: black;
  font-weight: 500;
  font-size: 15px;
}
.sub-div-category input:focus {
  outline: none;
}
.div-options {
  display: flex;
  align-items: center;
  justify-content: space-evenly;
  position: absolute;
  right: 10px;
}
.edit-icon:hover,
.remove-category-icon:hover,
.see-task-icon:hover {
  scale: 1.1;
  transition: scale 0.15s ease;
}
.edit-icon:active,
.remove-category-icon:active,
.see-task-icon:active {
  scale: 0.9;
  transition: scale 0.15s ease;
}

.options-icons {
  height: 25px;
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
.sub-div-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
}
.edit-input {
  background-color: red;
  width: 70%;
  background-color: rgb(218, 218, 218);
  height: 40px;
  border-radius: 25px;
}
.edit-input:focus {
  text-indent: 16px;
}
.edit-div {
  display: flex;
  column-gap: 15px;
  margin-left: 10px;
}
.edit-button {
  border: none;
  background-color: var(--electric-violet);
  color: white;
  font-family: Poppins;
  border-radius: 25px;
  padding: 0px 10px;
  margin-right: 5px;
  cursor: pointer;
  box-shadow: inset -5px -5px 5px rgba(0, 0, 0, 0.25);
  transition: opacity 0.15s ease;
}
.edit-button:hover {
  opacity: 0.9;
}
.edit-button:active {
  scale: 0.95;
}
.tick-remove {
  height: 40px;
  cursor: pointer;
}
.remove-div {
  display: flex;
  flex-direction: column;
  font-size: 12px;
  justify-content: center;
  align-items: center;
  width: 100%;
  margin-top: 10px;
}
.tick-icon {
  transition: scale 0.15s ease;
}
.tick-icon:hover {
  filter: contrast(1.5);
}
.remove-icon {
  transform: rotate(45deg);
  transition: transform 0.5s ease, scale 0.1s ease;
}
.remove-icon:hover {
  transform: rotate(360deg);
  filter: contrast(0.75);
}
.remove-icon:active,
.tick-icon:active {
  scale: 0.8;
}
</style>
