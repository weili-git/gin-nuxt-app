<template>
    <div>
        <van-cell-group>
            <van-nav-bar title="User Profile" left-arrow @click-left="onClickLeft" />
            <van-cell title="User Name">
                <template #default>
                    {{ username.first_name }}, {{ username.last_name }}
                </template>
            </van-cell>
            <van-cell title="Email">
                <template #default>
                    {{ username.email }}
                </template>
            </van-cell>
            <van-cell title="Birthday">
                <template #default>
                    {{ formatted_birthday }}
                </template>
            </van-cell>
        </van-cell-group>
    </div>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';

const apiRoutes = useApiRoutes();
const router = useRouter();
const username = ref('');
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
        username.value = res;
        console.log(username.value);
        const options = {
            year: 'numeric',
            month: 'long',
            day: 'numeric'
        };
        formatted_birthday.value = new Intl.DateTimeFormat('en-US', options).format(username.birthday);
    } catch (error) {
        console.log("error", error);
        router.push('/home/login');
    }
};

function logout() {

};
onMounted(() => {
    setProfile();
});
</script>