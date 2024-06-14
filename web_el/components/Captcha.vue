<template>
    <el-image width="240" height="80" @click="refreshCaptcha" :src="`${apiRoutes.CAPTCHA}${captcha_id}`"</el-image>
    <el-input v-model="captcha_input" placeholder="Captcha"></el-input>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';
const apiRoutes = useApiRoutes();
const captcha_id = ref('');
const captcha_input = defineModel();

const getCaptcha = async () => {
    const res = await $fetch(apiRoutes.CAPTCHA);
    captcha_id.value = res;
};
const getCaptchaId = () => {
    return captcha_id;
}
const refreshCaptcha = () =>{
    getCaptcha();
};
onMounted(()=>{
    getCaptcha();
});

defineExpose({getCaptcha, getCaptchaId});
</script>