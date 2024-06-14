<template>
    <el-header class="navbar">
        <el-menu mode="horizontal" background-color="#1976D2" text-color="#fff" active-text-color="#409EFF"
            :ellipsis="false">
            <el-menu-item index="home">
                <a href="/" class="nav-title">My App</a>
            </el-menu-item>

            <div style="flex-grow: 1;" />

            <el-menu-item index="user" class="avatar-menu-item" v-if="isLoggedin">
                <AvatarDropdown />
            </el-menu-item>

            <el-menu-item v-else>
                <a href="/home/login" style="text-decoration: none;font-size: medium;">Login in</a>
            </el-menu-item>
        </el-menu>
    </el-header>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';
const apiRoutes = useApiRoutes();
const isLoggedin = ref(false);


onMounted(async () => {
    const token = useCookie("token")
    if (token) {
        // fetch user info (name, avatar)
        const res = await $fetch(apiRoutes.PROFILE, {
            method: 'GET',
            headers: {
                Authorization: "Bearer:" + token.value
            }
        });
        if (res.id != 0) {
            isLoggedin.value = true;
        }
    }
}) 
</script>

<style scoped>
.navbar {
    height: 60px;
}

.nav-title {
    font-size: 24px;
    color: #fff;
    text-decoration: none;
}

.avatar-menu-item {
    line-height: 60px;
}
</style>