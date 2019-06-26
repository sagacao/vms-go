<template>
    <div class="table">
        <div class="app-container">
            <div class="handle-box">
                <el-select v-model="cur_channel"  placeholder="请选择" >
                    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                        <span style="float: left">{{ item.label }}</span>
                        <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                    </el-option>
                </el-select>
                <el-date-picker type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" v-model="select_daterange" style="width: 60%;">
                </el-date-picker>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                <el-col class="line" :span="20"></el-col>
            </div>
            <el-table :data="data" border class="table">
                <el-table-column prop="productCount" label="付费金额" width="220" sortable>
                </el-table-column>
                <el-table-column prop="count" label="付费笔数" width="220" sortable>
                </el-table-column>
                <el-table-column prop="sum" label="付费总额" width="220" sortable>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" :page-size="pageSize" layout="prev, pager, next" :total="totalnum">
                </el-pagination>
            </div>
        </div>
    </div>
</template>

<script>
/* eslint-disable */
import {parseTime} from '@/utils/index'
export default {
    name: 'paysum',
    data() {
        return {
            options:[
                {
                    value: 'xiaomi',
                    label: 'xiaomi'
                },
                {
                    value: 'vivo',
                    label: 'vivo'
                },
                {
                    value: 'wx',
                    label: 'wx'
                }
            ],
            tableData: [],
            totalnum : 0,
            cur_page: 1,
            pageSize: 30,
            cur_channel :'xiaomi',
            select_daterange:'',
            is_search: false
        }
    },
    created() {
        const end = new Date();
        const start = new Date();
        // start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
        this.select_daterange = [start, end];
        this.getData();
    },
    computed: {
        data() {
            return this.tableData;
        }
    },
    mounted(){
        //初始话下拉框的值
        // this.cur_channel=this.value;
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
            this.tableData = []
            this.totalnum = this.tableData.length
            this.is_search = true
            const user = this.$store.getters.name
            const page = this.$route.path
            // console.log(this.select_daterange)
            const sdate = this.select_daterange[0]
            const edate = this.select_daterange[1]
            
            // console.log(parseTime(sdate, '{y}-{m}-{d}'))
            const stime = parseTime(sdate, '{y}-{m}-{d}')//sdate.getFullYear() + '-' + (sdate.getMonth() + 1) + '-' + sdate.getDate()
            const etime = parseTime(edate, '{y}-{m}-{d}') //edate.getFullYear() + '-' + (edate.getMonth() + 1) + '-' + edate.getDate()
            const channel = this.cur_channel
            console.log("handleSearch", channel, stime, etime)
            this.$store.dispatch('svrmgr/GetPaySum', { user, page, channel, stime, etime}).then((res) => {
                if (res.list) {
                    this.tableData = res.list;
                    this.totalnum = this.tableData.length
                } else {
                    this.$message.info('没有数据')
                }
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
        width: 95%;
        font-size: 14px;
    }
    .red{
        color: #ff0000;
    }
    .mr10{
        margin-right: 10px;
    }
</style>
