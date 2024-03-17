export function checkLength(name, length) {
  if (name.length > length) {
    return name.substring(0, length) + "...";
  }
  return name;
}
//la funzione controlla se il nome è più lungo del valore nel parametro, se lo è taglia la parola (o la frase) e aggiunge 3 puntini
