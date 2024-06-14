<template>
    <div>Logout</div>
</template>

<script setup>
const router = useRouter();
const token = useCookie('token');

token.value = null;
router.push("/home/login")

// pinia

</script>