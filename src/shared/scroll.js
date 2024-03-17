export const scroll = () => {
  scrollToTheBottom.value.scrollTo({
    top: scrollToTheBottom.value.scrollHeight,
    behavior: "smooth",
  });
};
//la funzione scrolla in basso nella sezione in cui viene inserito il ref="scrollToTheBottom"
