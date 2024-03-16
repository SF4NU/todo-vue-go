export const wait = (ms) => {
  return new Promise((res, rej) => {
    setTimeout(() => {
      res();
    }, ms);
  });
};
