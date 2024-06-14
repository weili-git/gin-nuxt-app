<template>
    <div class="flex gap-4 mb-4">
        <span>Email: </span>
        <el-input v-model="email" placeholder="Account"></el-input>
        <span>Password: </span>
        <el-input v-model="password" label="Password" placeholder="Password"></el-input>
        <Captcha v-model="captcha_input" ref="catpcha_ref"/>
        <el-button type="primary" @click="login">Login</el-button>
        "Token:" {{ token }}
    </div>
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
            ElMessageBox.alert("Logged in!", "Success", {confirmButtonText: "ok"});
            
        }
    }catch(error){
        console.error('Login failed:', error);
        ElMessageBox.alert('Try again!', "Error", {confirmButtonText: "ok"})
        // clear captcha
        catpcha_ref.value.getCaptcha();
        captcha_input.value = '';
    }
};

</script>