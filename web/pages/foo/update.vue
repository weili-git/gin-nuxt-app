<template>
    <van-cell-group inset>
        <van-field v-model="id" label="ID" placeholder="ID" />
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

        <van-button type="primary" @click="handleClick">Submit</van-button>
    </van-cell-group>
    
</template>

<script setup>
    const id = ref('');
    const firstname= ref('');
    const lastname= ref('');
    const email= ref('');
    const password= ref('');
    const birthday= ref('');

    const showCalendar= ref(false);
    const birthday_format = ref('');

    async function handleClick() {
        try{
            const res = await $fetch(`http://127.0.0.1:8086/basic/update/${id.value}`, {
                method: "POST",
                body: {
                    "id": parseInt(id.value, 10),
                    "first_name": firstname.value,
                    "last_name": lastname.value,
                    "email": email.value,
                    "password": email.value,
                    "birthday": birthday.value
                }
            });
        }catch(error){
            console.log("Update failed:", error);
        }
    }

    const onConfirm = (date) => {
        showCalendar.value = false;
        birthday.value = date;
        birthday_format.value = `${date.getFullYear()}/${date.getMonth() + 1}/${date.getDate()}`;
        console.log(birthday.value);
    }
</script>