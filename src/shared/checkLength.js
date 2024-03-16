export function checkLength(name, length) {
  if (name.length > length) {
    return name.substring(0, length) + "...";
  }
  return name;
}
