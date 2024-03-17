<template class="template">
  <div ref="scrollToTheBottom" class="category-div">
    <img
      @click="$emit('backToAuth')"
      class="backToCategories"
      src="../icons/back-arrow.svg"
      alt="" />
    <div
      v-for="(category, i) in data"
      :key="category.ID"
      :class="`sub-div-category`">
      <div
        v-if="isEditing !== i && !category.is_loading"
        class="sub-div-wrapper">
        <!--Ci sono 2 controlli: isEditing controlla se l'utente ha schiacciato sul pulsante per modificare il nome della categoria e category.is_loading è un campo che non esiste più all'interno della tabella categorie dato che mi sono accorto che non serviva cambiarla anche nel database, quindi di base category.is_loading è sempre falsa tranne quando specificato-->
        <span class="category-name">{{ checkLength(category.title, 17) }}</span
        ><!--checkLength spezza il nome della categoria se è più lunga di 17 caratteri-->
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
      <div
        v-else-if="isEditing === i && !isRemoving && !category.is_loading"
        class="edit-div">
        <!--Ci sono 2 controlli: isEditing controlla se l'utente ha schiacciato sul pulsante per modificare il nome della categoria, isRemoving controlla se l'utente ha schiacciato sul pulsante per eliminare la categoria-->
        <input
          class="edit-input"
          v-model="updatedTitle"
          @keyup.enter="editFinished(i, category.ID)"
          type="text" />
        <button class="edit-button" @click="editFinished(i, category.ID)">
          Aggiorna
        </button>
      </div>
      <div v-else-if="!category.is_loading" class="remove-div">
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
      <div v-else class="loader-wrapper">
        <div class="loader"></div>
      </div>
    </div>
  </div>
</template>

<script setup>
import axios from "axios";
import { ref, onBeforeMount, onUnmounted, watch } from "vue";
import { getCurrentTime } from "../../shared/getCurrentTime";
import { checkLength } from "../../shared/checkLength";
import { scroll } from "../../shared/scroll";

const userProps = defineProps(["sendUserInput", "userId"]); //sendUserInput è il contenuto dell'input per aggiungere nuove tabelle e userId è l'id dello user loggato ricevuto subito dopo il login
let isEditing = ref(null);
let data = ref(null);
const scrollToTheBottom = ref(null);
let updatedTitle = ref("");
let isRemoving = ref(false);

const isEditingCategory = (id) => {
  //controlla quale categoria è in corso di modifica
  isEditing.value = id;
};
const editFinished = (index, id) => {
  //aggiorna il nome della categoria dopo che viene schiacciato il pulsante aggiorna
  updateCategory(index, id);
  isEditing.value = null;
};

const updateCategory = async (index, id) => {
  //è la funzione del metodo PUT che prende il nuovo nome inserito nell'apposita casella di input
  if (updatedTitle.value !== "") {
    try {
      data.value[index].is_loading = true; //isLoading viene impostato a vero così che possa partire l'animazione di caricamento
      setTimeout(() => {
        //inizialmente il timeout era dopo il await del metodo put ma ci metteva troppo a mandare l'aggiornamento e poi liberare la variabile quindi con 300ms di timeout è quasi istantaneo e parte comunque la richiesta al server
        updatedTitle.value = "";
      }, 300);
      await axios.put(
        `https://go-fiber-prova-production.up.railway.app/categories/${id}`,
        {
          title: updatedTitle.value,
          last_modified: getCurrentTime(),
        }
      );
      fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle categorie
    } catch (error) {
      console.error(error);
    }
  }
};

const addCategory = async () => {
  //la funzione crea nuove categorie partendo dall'input dell'utente
  if (userProps.sendUserInput !== "" && userProps.sendUserInput !== " ") {
    //l'if serve a verificare se c'è effettivamente qualcosa nell'input per non far partire una richiesta vuota
    try {
      scroll(); //una funzione scroll che serve a scrollare in basso dove si trova la categoria appena creata

      const newCategoryTitle = {
        is_loading: true,
      };
      data.value.push(newCategoryTitle); //utilizza il push a differenza della funzione di update perché non c'è un index a cui fare riferimento
      const res = await axios.post(
        `https://go-fiber-prova-production.up.railway.app/categories`,
        {
          title: userProps.sendUserInput,
          last_modified: getCurrentTime(),
          user_id: userProps.userId,
        }
      );
      fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle categorie
    } catch (error) {
      console.error(error);
    }
  }
};

watch(
  () => userProps.sendUserInput, //questo watch controlla se l'input è cambiato, ma solo dopo che è stato inviato con un invio o con un click dato che cambia all'interno del componente main
  () => {
    addCategory();
  }
);

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

const deleteCategory = async (id, index) => {
  //si occupa di eliminare la categoria in base all'id, mentre l'index serve solo per sapere a quale div assegnare il caricamento
  try {
    isRemoving.value = false;
    isEditing.value = null;
    data.value[index].is_loading = true;
    await axios.delete(
      `https://go-fiber-prova-production.up.railway.app/categories/${id}`
    );
    await fetchData(); //richiama fetchData quando ha finito di aggiornare il nome, inoltre isLoading viene impostato di nuovo su falso dato che non esiste più un isLoading nella tabella delle categorie
  } catch (error) {
    console.error(error);
  }
};
const fetchData = async () => {
  //la funzione che si occupa di recuperare le categorie ,relative all'utente, dal database
  try {
    const res = await axios.get(
      `https://go-fiber-prova-production.up.railway.app/users/${userProps.userId}/categories`
    );
    data.value = res.data.sort((a, b) => a.ID - b.ID); //ordina i dati ricevuti in base all'id
  } catch (error) {
    console.error(error);
  }
};
onBeforeMount(() => {
  fetchData(); //si occupa di recuperare i dati appena il componente viene iniziato, penso che anche un onMounted avrebbe fatto la stessa cosa in questo caso
});
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
.loader-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
}
.loader {
  width: 40px;
  aspect-ratio: 1;
  color: #f03355;
  position: relative;
  background: conic-gradient(from 134deg at top, currentColor 92deg, #0000 0)
      top,
    conic-gradient(from -46deg at bottom, currentColor 92deg, #0000 0) bottom;
  background-size: 100% 50%;
  background-repeat: no-repeat;
  scale: 0.8;
}
.loader:before {
  content: "";
  position: absolute;
  inset: 0;
  --g: currentColor 14.5px, #0000 0 calc(100% - 14.5px), currentColor 0;
  background: linear-gradient(45deg, var(--g)),
    linear-gradient(-45deg, var(--g));
  animation: l7 1s infinite cubic-bezier(0.3, 1, 0, 1);
}
@keyframes l7 {
  33% {
    inset: -10px;
    transform: rotate(0deg);
  }
  66% {
    inset: -10px;
    transform: rotate(90deg);
  }
  100% {
    inset: 0;
    transform: rotate(90deg);
  }
}
</style>
