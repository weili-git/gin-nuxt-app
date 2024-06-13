<template>
    <van-cell-group inset>
        <van-field v-model="email" label="Account" placeholder="Account"></van-field>
        <van-field v-model="password" label="Password" placeholder="Password"></van-field>
        <Captcha v-model="captcha_input" ref="catpcha_ref"/>
        <van-button type="primary" @click="login">Login</van-button>
        "Token:" {{ token }}
    </van-cell-group>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';
const apiRoutes = useApiRoutes();
const email = ref('');
const password = ref('');
const token = useCookie('token', {maxAge: 60 * 60});
const catpcha_ref = ref(null);
const captcha_input = ref('');

async function login() {
    try{
        const captcha_id = catpcha_ref.value.getCaptchaId();
        const verifyRes = await $fetch(apiRoutes.LOGIN, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: {
                "captcha_id": captcha_id.value,
                "captcha_value": captcha_input.value,
                "email": email.value,
                "password": password.value,
            }
        });
        if (verifyRes.msg === "Login OK!") {
            token.value = verifyRes.token;
            showToast("Logged in!");
        }
    }catch(error){
        console.error('Login failed:', error);
        showToast('Try again!')
        // clear captcha
        catpcha_ref.value.getCaptcha();
        captcha_input.value = '';
    }
};

</script>