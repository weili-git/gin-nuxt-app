<template>
    <van-nav-bar title="This is the test page">
        <template #right>
            <van-space :size="10" v-if="isAuthenticated()">
                <!-- if not logged in -->
                <NuxtLink to="/home/login">Log In</NuxtLink>
                <NuxtLink to="/home/signup">Sign Un</NuxtLink>
            </van-space>
            <van-space :size="10" v-else>
                <!-- if logged in -->
                <NuxtLink to="/post/create">Create Post</NuxtLink> 
                <NuxtLink to="/home/profile">Profile</NuxtLink> 
                <NuxtLink to="/home/logout">Log Out</NuxtLink> 
                <van-dropdown-menu>
                    <template #title>
                        <van-image round width="2rem" height="2rem" src="https://fastly.jsdelivr.net/npm/@vant/assets/cat.jpeg"></van-image>
                    </template>
                    <van-dropdown-item v-model="value1" :options="items"/>
                </van-dropdown-menu>
                
            </van-space>
        </template>
    </van-nav-bar>
</template>

<script setup>
const value1 = ref(0);
const items = [
    {text: 'option 1', value: 0},
    {text: 'option 2', value: 1}
]

    function isAuthenticated() {
        return false;
    }
</script>