<template>
    <div>
        <Navbar/>
        <el-row>
            <el-col :span="8"></el-col>
            <el-col :span="8">
                <slot/>
                <Footer/>
            </el-col>
            <el-col :span="8"></el-col>
        </el-row>
    </div>
</template>