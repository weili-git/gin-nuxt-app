<template>
    <van-nav-bar title="This is the test page">
        <template #right>
            <van-space :size="10">
                <NuxtLink to="/home/login">Log In</NuxtLink>
                <NuxtLink to="/home/signup">Sign Un</NuxtLink>
            </van-space>
        </template>
    </van-nav-bar>
</template>