export const wait = (ms) => {
  return new Promise((res, rej) => {
    setTimeout(() => {
      res();
    }, ms);
  });
};
//una promessa che serve come timer all'interno delle funzioni async/await
