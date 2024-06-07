<template>
    <div>
        <h1>User Profile</h1>
        <p>User Name: {{ username }}</p>
        <button @click="logout">Log out</button>
    </div>
</template>

<script setup>
import useApiRoutes from '~/utils/apiRoutes';

const apiRoutes = useApiRoutes();
const router = useRouter();
const username = ref('');
async function setProfile(){
    const token = useCookie('token');
    try{
        const res = await $fetch(apiRoutes.PROFILE, {
            method: 'POST',
            headers: {
                Authorization: "Bearer: " + token.value
            }
        });
        username.value = res.msg;
    }catch(error){
        console.log("error", error);
        router.push('/home/login');
    }
};

function logout(){

};
onMounted(()=>{
    setProfile();
});
</script>