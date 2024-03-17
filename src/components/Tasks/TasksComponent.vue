<template>
  <div ref="scrollToTheBottom" class="category-div">
    <img
      @click="$emit('backToCategories')"
      class="backToCategories"
      src="../icons/back-arrow.svg"
      alt="" />
    <div v-for="(task, i) in data" :key="i" :class="`sub-div-category`">
      <div v-if="isEditing !== i && !task.is_loading" class="sub-div-wrapper">
        <!--Ci sono 2 controlli: isEditing controlla se l'utente ha schiacciato sul pulsante per modificare il nome della task e task.is_loading è un campo che non esiste più all'interno della tabella tasks dato che mi sono accorto che non serviva cambiarla anche nel database, quindi di base task.is_loading è sempre falsa tranne quando specificato-->
        <span class="category-name">{{
          checkLength(task.description, 20)
        }}</span
        ><!--checkLength spezza il nome della categoria se è più lunga di 20 caratteri-->
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
      <div
        v-else-if="isEditing === i && !isRemoving && !task.is_loading"
        class="edit-div">
        <input
          class="edit-input"
          v-model="updatedDescription"
          @keyup.enter="editFinished(i, task.ID)"
          type="text" />
        <button class="edit-button" @click="editFinished(i, task.ID)">
          Aggiorna
        </button>
      </div>
      <div v-else-if="!task.is_loading" class="remove-div">
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
      <div v-else class="loader-wrapper">
        <div class="loader"></div>
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
let isEditing = ref(null);
let isRemoving = ref(false);
let updatedDescription = ref("");
const scrollToTheBottom = ref(null);

onBeforeMount(() => {
  //si occupa di recuperare i dati appena il componente viene iniziato
  fetchData();
});

const checkBox = (i, id, des) => {
  //forse mi sono un po incasinato con il checkbox anche perché non ho usato il ref, in pratica questa funzione aggiunge o toglie la spunta sulla checkbox e allo stesso tempo si occupa di mandare una richiesta di aggiornamento al server per cambiare lo stato di completamento
  const check = document.getElementById(`${i}`);
  if (check.classList.contains("checkbox-animate")) {
    check.classList.remove("checkbox-animate");
    completeTask(i, id, false, des);
  } else {
    check.classList.add("checkbox-animate");
    completeTask(i, id, true, des);
  }
};

const completeTask = async (index, id, check, des) => {
  //la funzione si occupa di mandare una richiesta di aggiornamento al server per cambiare lo stato di completamento
  try {
    data.value[index].is_loading = true; //isLoading viene impostato a vero così che possa partire l'animazione di caricamento
    await axios.put(
      `https://go-fiber-prova-production.up.railway.app/tasks/${id}`,
      {
        description: des,
        completed: check,
        last_modified: getCurrentTime(),
      }
    );
    await fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle task
  } catch (error) {
    console.error(error);
  }
};

const isEditingTask = (id) => {
  //controlla quale categoria è in corso di modifica
  isEditing.value = id;
};
const editFinished = (index, id) => {
  //aggiorna il nome della categoria dopo che viene schiacciato il pulsante aggiorna
  updateTask(index, id);
  isEditing.value = null;
};

const updateTask = async (index, id) => {
  //è la funzione del metodo PUT che prende il nuovo nome inserito nell'apposita casella di input
  if (updatedDescription.value !== "") {
    try {
      data.value[index].is_loading = true; //isLoading viene impostato a vero così che possa partire l'animazione di caricamento

      setTimeout(() => {
        updatedDescription.value = ""; //inizialmente era dopo il await del metodo put ma ci metteva troppo a mandare l'aggiornamento e poi liberare la variabile quindi con 300ms di timeout è quasi istantaneo e parte comunque la risposta
      }, 300);
      await axios.put(
        `https://go-fiber-prova-production.up.railway.app/tasks/${id}`,
        {
          description: updatedDescription.value,
          completed: false,
          last_modified: getCurrentTime(),
        }
      );
      await fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle task
    } catch (error) {
      console.error(error);
    }
  }
};

const fetchData = async () => {
  //la funzione che si occupa di recuperare le tasks ,relative alla categoria, dal database
  try {
    const res = await axios.get(
      `https://go-fiber-prova-production.up.railway.app/categories/${props.categoryId}/tasks`
    );
    data.value = res.data.sort((a, b) => a.ID - b.ID); //ordina i dati ricevuti in base all'id
  } catch (error) {
    console.error(error);
  }
};

const addTask = async () => {
  //la funzione crea nuove task partendo dall'input dell'utente
  if (props.sendUserInput !== "" && props.sendUserInput !== " ") {
    //l'if serve a verificare se c'è effettivamente qualcosa nell'input per non far partire una richiesta vuota
    try {
      scroll(); //una funzione scroll che serve a scrollare in basso dove si trova la task appena creata
      const newTaskLoadingState = {
        is_loading: true,
      };
      data.value.push(newTaskLoadingState); //utilizza il push a differenza della funzione di update perché non c'è un index a cui fare riferimento
      const res = await axios.post(
        `https://go-fiber-prova-production.up.railway.app/tasks`,
        {
          description: props.sendUserInput,
          last_modified: getCurrentTime(),
          completed: false,
          category_id: props.categoryId,
        }
      );
      await fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle task
    } catch (error) {
      console.error(error);
    }
  }
};

const askToDelete = (i) => {
  //la funzione per il pulsante di conferma cancellazione
  isEditing.value = i;
  isRemoving.value = true;
};

const doNotDelete = () => {
  //la funzione per il pulsante per rifiutare la cancellazione
  isEditing.value = null;
  isRemoving.value = false;
};

const deleteTask = async (id, index) => {
  // elimina la task selezionata
  try {
    isRemoving.value = false;
    isEditing.value = null;
    data.value[index].is_loading = true; //isLoading viene impostato a vero così che possa partire l'animazione di caricamento

    // data.value = data.value.filter((_, i) => i !== index); //dato che ci vogliono 1-2 secondi per mandare request al server e riceverne un'altra ho pensato che è meglio aggiornare prima quello che vede l'user e lasciare fare le richieste in background così da non rallentare nessuno
    await axios.delete(
      `https://go-fiber-prova-production.up.railway.app/tasks/${id}`
    );
    await fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle task
  } catch (error) {
    console.error(error);
  }
};

watch(
  //questo watch controlla se l'input è cambiato, ma solo dopo che è stato inviato con un invio o con un click dato che cambia all'interno del componente main
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
}; //la funzione scrolla in basso nella sezione in cui viene inserito il ref="scrollToTheBottom"
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
  overflow: hidden;
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
</style>
