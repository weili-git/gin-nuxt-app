<template>
    <van-cell-group inset>
        <van-field v-model="firstname" label="First Name" placeholder="First Name" />
        <van-field v-model="lastname" label="Last Name" placeholder="Last Name" />
        <van-field v-model="email" label="Email" placeholder="Email" />
        <van-field v-model="password" label="Password" placeholder="Password" />

        <van-field 
            v-model="birthday_format" 
            label="Birthday" 
            placeholder="Birthday"
            readonly
            @click="showCalendar = true"
        />
        <van-calendar v-model:show="showCalendar" @confirm="onConfirm" switch-mode="year-month" />
        <Captcha v-model="captcha_input" ref="captcha_ref"/>
        <van-button type="primary" @click="handleClick">Submit</van-button>
    </van-cell-group>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';
const apiRoutes = useApiRoutes();
const firstname= ref('');
const lastname= ref('');
const email= ref('');
const password= ref('');
const birthday= ref('');

const showCalendar= ref(false);
const birthday_format = ref('');

const captcha_ref = ref(null);
const captcha_input = ref('');

async function handleClick() {
    // check datatype of each field!
    try {
        const captcha_id = captcha_ref.value.getCaptchaId();
        const response = await $fetch(apiRoutes.SIGNUP,
            {
                method: "POST",
                body:{
                    "first_name": firstname.value, 
                    "last_name": lastname.value,
                    "email": email.value,
                    "password": email.value,
                    "birthday": birthday.value,
                    "captcha_id": captcha_id.value,
                    "captcha_value": captcha_input.value,
                }
            }
        );
        console.log('Response:', response.data);
        await navigateTo('/foo');
    }catch(error){
        console.error('Error:', error);
    }
}
const onConfirm = (date) => {
    showCalendar.value = false;
    birthday.value = date.toISOString();
    birthday_format.value = `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`;
    console.log(birthday.value);
}

    
</script>