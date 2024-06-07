<template>
    <div>
        <NavBar/>
        <van-row>
            <van-col span="8"></van-col>
            <van-col span="8">
                <slot/>
                <Footer/>
            </van-col>
            <van-col span="8"></van-col>
        </van-row>
    </div>
</template>