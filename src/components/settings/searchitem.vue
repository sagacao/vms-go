<template>
    <div>
        <el-select v-model="selectVal" :placeholder="defaultgame || 请选择游戏ID" @change="currentVal">
            <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                <span style="float: left">{{ item.label }}</span>
                <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
            </el-option>
        </el-select>
    </div>
</template>

<script>
/* eslint-disable */
export default {
    name: 'SearchItem',
    props: {
        defaultgame : {
            type: String,
            required: true
        },
        selectVal: {
            type: String,
            required: true
        }
    },
    data() {
        return {
        options: [{
            value: '20201',
            label: '方块部落'
        }, {
            value: '20401',
            label: '春节红包'
        }],
        selectVal: ''
        }
    },
    created() {
        console.log("searchitem created")
    },
    methods: {
        getValue () {
            return this.selectVal
        },
        currentVal() {
            this.defaultgame = this.selectVal
        }
    }
}
</script>

<style scoped>
    .handle-box {
        margin-bottom: 50px;
    }
</style>