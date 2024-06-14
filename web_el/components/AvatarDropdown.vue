<template>
    <div class="avatar-dropdown">
        <el-dropdown @command="handleCommand">
            <span class="el-dropdown-link">
                <el-avatar :src="avatar" class="user-avatar"></el-avatar>
            </span>
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item command="profile">Profile</el-dropdown-item>
                    <el-dropdown-item command="settings">Settings</el-dropdown-item>
                    <el-dropdown-item command="logout">Logout</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
    </div>
</template>

<script setup>
const avatar = ref('/avatar.jpg')
const router = useRouter()

function handleCommand(command) {
    switch (command) {
        case 'profile':
            router.push('/home/profile')
            break;
        case 'settings':
            router.push('/home/settings')
            break;
        case 'logout':
            router.push('/home/logout')
            break;
    }
}
</script>

<style>
.avatar-dropdown {
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

.user-avatar {
    cursor: pointer;
}
</style>