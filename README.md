# Todo list con Vue.js e GO (Fiber + Gorm)
## [DEMO](https://sf4nu.github.io/todo-vue-go/)

### Frontend Vue.js

Il frontend l'ho realizzato interamente con il framework Vue.js 
Permette a ciascun utente di avere le proprie categorie e task che può aggiungere con l'apposito input in alto:

![Immagine 2024-03-17 183522](https://github.com/SF4NU/todo-vue-go/assets/129513838/3a95ecd3-b1fb-4e69-9893-eda39a3fe4c5)
---
Inoltre, per le task si può mettere una spunta che segna se sono state completate: 

![Immagine 2024-03-17 184435](https://github.com/SF4NU/todo-vue-go/assets/129513838/44dbfe34-df58-4620-a5c0-c8010d084860)
---

### Backend (Golang)
1. Utilizza il framework Fiber per creare un server web.
2. Gestisce le richieste HTTP per recuperare, creare, aggiornare ed eliminare categorie e attività dal database PostgreSQL.
3. Si occupa anche di gestire la registrazione e l'autenticazione degli utenti.
4. Utilizza il pacchetto GORM per interagire con il database e gestire le migrazioni del database.
   
Esempio di codice:
   
![code](https://github.com/SF4NU/todo-vue-go/assets/129513838/67d45668-135e-4b2a-9114-8e40d9f69f31)

