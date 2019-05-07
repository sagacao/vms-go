<template>
    <div class="table">
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item><i class="el-icon-lx-cascades"></i> 游戏数据</el-breadcrumb-item>
                <el-breadcrumb-item>基本数据</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="handle-box">
                <el-date-picker type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" v-model="select_daterange" style="width: 60%;">
                </el-date-picker>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                <el-col class="line" :span="20"></el-col>
            </div>
            <el-table :data="data" border class="table">
                <el-table-column prop="channel" label="渠道" width="120">
                </el-table-column>
                <el-table-column prop="game" label="游戏" width="120">
                </el-table-column>
                <el-table-column prop="newly" label="新增" width="100">
                </el-table-column>
                <el-table-column prop="tow_pr" label="次留" width="100">
                </el-table-column>
                <el-table-column prop="three_pr" label="三留" width="100">
                </el-table-column>
                <el-table-column prop="seven_pr" label="七留" width="100">
                </el-table-column>
                <el-table-column prop="logdate" label="日期" sortable >
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="10">
                </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
/* eslint-disable */ 
    export default {
        name: 'logger',
        data() {
            return {
                tableData: [],
                cur_page: 1,
                select_daterange:'',
                is_search: false
            }
        },
        created() {
            const end = new Date();
            const start = new Date();
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
            this.select_daterange = [start, end];
            this.getData();
        },
        computed: {
            data() {
                return this.tableData;
            }
        },
        mounted(){
            // const end = new Date();
            // const start = new Date();
            // start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
            // this.select_daterange = [start, end];
        },
        methods: {
            // 分页导航
            handleCurrentChange(val) {
                this.cur_page = val;
                this.getData();
            },
            getData() {
                this.handleSearch();
            },
            handleSearch() {
                this.is_search = true;
                const user = this.$store.getters.name
                const page = this.$route.path
                // console.log(this.select_daterange)
                const sdate = this.select_daterange[0]
                const edate = this.select_daterange[1]
                // console.log(this.select_daterange[0].getTime())
                const stime = sdate.getFullYear() + '-' + (sdate.getMonth() + 1) + '-' + sdate.getDate()
                const etime = edate.getFullYear() + '-' + (edate.getMonth() + 1) + '-' + edate.getDate()
                this.$store.dispatch('GetLogger', { user, page, stime, etime}).then((res) => {
                    this.tableData = res.data.list;
                }).catch(err => {
                    console.log(err)
                    this.$message.error(err);
                });
            },
        }
    }

</script>

<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 120px;
    }

    .handle-input {
        width: 300px;
        display: inline-block;
    }
    .del-dialog-cnt{
        font-size: 16px;
        text-align: center
    }
    .table{
        width: 100%;
        font-size: 14px;
    }
    .red{
        color: #ff0000;
    }
    .mr10{
        margin-right: 10px;
    }
</style>
