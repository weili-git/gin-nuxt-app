<template>
    <el-descriptions title="User Infor" border :column="1">
        <el-descriptions-item label="Name:">{{ user_info.first_name }}, {{ user_info.last_name }}</el-descriptions-item>
        <el-descriptions-item label="Email:">{{ user_info.email }}</el-descriptions-item>
        <el-descriptions-item label="Password:">{{ user_info.password }}</el-descriptions-item>
        <el-descriptions-item label="Birthday:">{{ formatted_birthday }}</el-descriptions-item>
    </el-descriptions>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';

const apiRoutes = useApiRoutes();
const router = useRouter();
const user_info = ref('');
const formatted_birthday = ref('');

async function setProfile() {
    const token = useCookie('token');
    try {
        const res = await $fetch(apiRoutes.PROFILE, {
            method: 'GET',
            headers: {
                Authorization: "Bearer:" + token.value
            }
        });
        user_info.value = res;
        const options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        };
        formatted_birthday.value = new Intl.DateTimeFormat('en-US', options).format(user_info.birthday);
    } catch (error) {
        console.log("error", error);
        router.push('/home/login');
    }
};

onMounted(() => {
    setProfile();
});
</script>