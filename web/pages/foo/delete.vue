<template>
    <van-cell-group>
        <van-field v-model="user_id" label="ID" placeholder="ID"></van-field>
        <van-button type="primary" @click="handleClick">Delete</van-button>
    </van-cell-group>

</template>

<script setup>
    const user_id = ref('');
    async function handleClick() {
        const res = await $fetch(`http://127.0.0.1:8086/basic/delete/${user_id.value}`, { // important to use param
            method: "DELETE",
        });
        console.log(res);
    }

</script>