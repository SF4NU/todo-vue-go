import { format } from "date-fns";
import { computed } from "vue";

export const getCurrentTime = () => {
  const currentDate = new Date();
  const formattedDate = computed(() => format(currentDate, "dd/MM - HH:mm"));
  return formattedDate.value;
};
