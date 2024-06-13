<template>
    <div>aaa</div>
</template>

<script setup>
const router = useRouter();
const token = useCookie('token');

token.value = null;
router.push("/home/login")
</script>