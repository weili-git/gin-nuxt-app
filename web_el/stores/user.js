import { defineStore } from "pinia";
import { ref } from 'vue';

export const useUserStore = defineStore('user', () => {
    const token = ref('');
    const user = ref({});
    function reset() {
        token.value = "";
        user.value = {};
    }
    return { token, user, reset };
}, {
    persist: true
});