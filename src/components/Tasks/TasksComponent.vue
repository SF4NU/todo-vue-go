<template>
  <div ref="scrollToTheBottom" class="category-div">
    <img
      @click="$emit('backToCategories')"
      class="backToCategories"
      src="../icons/back-arrow.svg"
      alt="" />
    <div v-for="(task, i) in data" :key="i" :class="`sub-div-category`">
      <div v-if="isEditing !== i" class="sub-div-wrapper">
        <span class="category-name">{{
          checkLength(task.description, 20)
        }}</span>
        <div class="div-options">
          <img
            @click="isEditingTask(i)"
            class="options-icons edit-icon"
            src="../icons/edit.svg"
            alt="" />
          <img
            @click="askToDelete(i)"
            class="options-icons remove-category-icon"
            src="../icons/redbin.svg"
            alt="" />
          <input
            :id="i"
            @click="checkBox(i, task.ID, task.description)"
            :class="{ checkbox: true, 'checkbox-animate': task.completed }"
            type="checkbox" />
        </div>
        <span class="edit-time">Ultima modifica: {{ task.last_modified }}</span>
      </div>
      <div v-else-if="isEditing === i && !isRemoving" class="edit-div">
        <input class="edit-input" v-model="updatedDescription" type="text" />
        <button class="edit-button" @click="editFinished(i, task.ID)">
          Aggiorna
        </button>
      </div>
      <div v-else class="remove-div">
        <div>Stai per rimuovere "{{ checkLength(task.description, 10) }}"</div>
        <div>
          <img
            @click="deleteTask(task.ID, i)"
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
import { onBeforeMount, ref, watch } from "vue";
import axios from "axios";
import { getCurrentTime } from "@/shared/getCurrentTime";
import { checkLength } from "../../shared/checkLength";

const props = defineProps(["categoryId", "sendUserInput"]);
const data = ref(null);
let timeoutAddId = null;
let isEditing = ref(null);
let isRemoving = ref(false);
let updatedDescription = ref("");
const scrollToTheBottom = ref(null);
let watchVar = ref(0);

onBeforeMount(() => {
  fetchData();
});

const checkBox = (i, id, des) => {
  const check = document.getElementById(`${i}`);
  if (check.classList.contains("checkbox-animate")) {
    check.classList.remove("checkbox-animate");
    completeTask(i, id, false, des);
  } else {
    check.classList.add("checkbox-animate");
    completeTask(i, id, true, des);
  }
};

let timeoutCompleteId = null;

const completeTask = async (index, id, check, des) => {
  try {
    const completeTask = {
      description: des,
      completed: check,
      last_modified: getCurrentTime(),
    };
    data.value[index] = completeTask;
    await axios.put(
      `https://go-fiber-prova-production.up.railway.app/tasks/${id}`,
      {
        description: des,
        completed: check,
        last_modified: getCurrentTime(),
      }
    );
    timeoutCompleteId ? clearTimeout(timeoutCompleteId) : null;
    timeoutCompleteId = await setTimeout(() => {
      fetchData();
    }, 3000);

    fetchData();
  } catch (error) {
    console.error(error);
  }
};

const isEditingTask = (id) => {
  isEditing.value = id;
};
const editFinished = (index, id) => {
  updateTask(index, id);
  isEditing.value = null;
};

const updateTask = async (index, id) => {
  if (updatedDescription.value !== "") {
    try {
      const updateTaskDescription = {
        description: updatedDescription.value,
        completed: false,
        last_modified: getCurrentTime(),
      };
      data.value[index] = updateTaskDescription;

      await axios.put(
        `https://go-fiber-prova-production.up.railway.app/tasks/${id}`,
        {
          description: updatedDescription.value,
          completed: false,
          last_modified: getCurrentTime(),
        }
      );
      updatedDescription.value = "";
      fetchData();
    } catch (error) {
      console.error(error);
    }
  }
};

const fetchData = async () => {
  try {
    const res = await axios.get(
      `https://go-fiber-prova-production.up.railway.app/categories/${props.categoryId}/tasks`
    );
    data.value = res.data.sort((a, b) => a.ID - b.ID);
  } catch (error) {
    console.error(error);
  }
};

const addTask = async () => {
  if (props.sendUserInput !== "" && props.sendUserInput !== " ") {
    try {
      scroll();

      const newTaskDescription = {
        description: props.sendUserInput,
        last_modified: getCurrentTime(),
        completed: false,
        category_id: props.categoryId,
      };
      data.value.push(newTaskDescription);
      const res = await axios.post(
        `https://go-fiber-prova-production.up.railway.app/tasks`,
        {
          description: props.sendUserInput,
          last_modified: getCurrentTime(),
          completed: false,
          category_id: props.categoryId,
        }
      );
      timeoutAddId ? clearTimeout(timeoutAddId) : null;
      timeoutAddId = await setTimeout(() => {
        //aspetta prima di aggiornare di rimandare la richiesta GET nel caso l'utente stia ancora aggiungendo categorie
        fetchData();
      }, 2000);
    } catch (error) {
      console.error(error);
    }
  }
};

const askToDelete = (i) => {
  isEditing.value = i;
  isRemoving.value = true;
};

const doNotDelete = () => {
  isEditing.value = null;
  isRemoving.value = false;
};

let timeoutId = null;

const deleteTask = async (id, index) => {
  // elimina la task selezionata
  try {
    isRemoving.value = false;
    isEditing.value = null;
    data.value = data.value.filter((_, i) => i !== index); //dato che ci vogliono 1-2 secondi per mandare request al server e riceverne un'altra ho pensato che è meglio aggiornare prima quello che vede l'user e lasciare fare le richieste in background così da non rallentare nessuno
    await axios.delete(
      `https://go-fiber-prova-production.up.railway.app/tasks/${id}`
    );
    timeoutId ? clearTimeout(timeoutId) : null; // se viene schiacciato di nuovo il pulsante "deleteCategory" riparte il timer per evitare sovrapposizioni
    timeoutId = await setTimeout(() => {
      watchVar.value++; // quando questa variabile cambia riparte la funzione fetchData, il timer serve per quando vengono cancellate molto velocemente le task e quindi si evita il flash delle categorie che appaiono ogni volta che cambia watchVar
    }, 5000);
  } catch (error) {
    console.error(error);
  }
};

watch(
  () => props.sendUserInput,
  () => {
    addTask();
  }
);

const scroll = () => {
  scrollToTheBottom.value.scrollTo({
    top: scrollToTheBottom.value.scrollHeight,
    behavior: "smooth",
  });
};
</script>

<style scoped>
.backToCategories {
  position: absolute;
  top: 25px;
  left: 25px;
  height: 35px;
  cursor: pointer;
  transition: scale 0.15s ease;
}
.backToCategories:hover {
  filter: invert();
}
.backToCategories:active {
  scale: 0.8;
}
.checkbox,
.checkbox-animate {
  -webkit-appearance: none;
  -moz-appearance: none;
  appearance: none;
  height: 22px;
  width: 22px;
  background-color: white;
  border-radius: 50px;
  border: 1px solid black;
  cursor: pointer;
  position: relative;
  /* overflow: hidden; */
}
.checkbox-animate::after {
  content: "";
  width: 100%;
  height: 100%;
  position: absolute;
  background-color: yellowgreen;
  top: 0;
  left: 0;
  border-radius: 50px;
  animation: move1 0.6s linear 1 forwards;
}

@keyframes move1 {
  0% {
    opacity: 0;
    transform: translate(15px, -15px);
    scale: 1;
  }
  50% {
    opacity: 0.5;
    transform: translate(7px, -7px);
    scale: 0.5;
  }
  100% {
    opacity: 1;
    transform: translate(0px, 0px);
    scale: 0.8;
  }
}

/* .checkbox-clicked {
  background-color: yellowgreen;
} */
@keyframes identifier {
}
</style>