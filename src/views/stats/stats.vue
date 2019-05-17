<template>
    <div class="table">
        <div class="app-container">
            <div class="handle-box">
                <el-date-picker type="daterange" start-placeholder="开始日期" end-placeholder="结束日期" v-model="select_daterange" style="width: 60%;">
                </el-date-picker>
                <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                <el-col class="line" :span="20"></el-col>
                <el-button type="primary"  @click="handleAdd">新增<i class="el-icon-upload el-icon--right"></i></el-button>
            </div>
            <el-table :data="data" border class="table">
                <el-table-column prop="channel" label="渠道" width="80">
                </el-table-column>
                <el-table-column prop="game" label="游戏" width="80">
                </el-table-column>
                <el-table-column prop="newly" label="新增" width="80">
                </el-table-column>
                <el-table-column prop="tow_pr" label="次留" width="80">
                </el-table-column>
                <el-table-column prop="three_pr" label="三留" width="80">
                </el-table-column>
                <el-table-column prop="seven_pr" label="七留" width="80">
                </el-table-column>
                <el-table-column prop="retention" label="比例" width="80">
                </el-table-column>
                <el-table-column prop="logdate" label="日期" sortable width="120">
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <div class="pagination">
                <el-pagination background @current-change="handleCurrentChange" layout="prev, pager, next" :total="10">
                </el-pagination>
            </div>
        </div>

        <!-- 编辑弹出框 -->
        <el-dialog title="编辑" :visible.sync="editVisible" width="50%">
            <el-form ref="form" :model="form" label-width="100px">
                <el-form-item label="日期">
                    <el-date-picker type="date" placeholder="选择日期" v-model="form.logdate" value-format="yyyy-MM-dd" style="width: 100%;"></el-date-picker>
                </el-form-item>
                <el-form-item label="渠道">
                    <el-input v-model="form.channel"></el-input>
                </el-form-item>
                <el-form-item label="游戏">
                    <el-select v-model="form.game" placeholder="选择游戏" class="handle-select mr10">
                        <el-option key="1" label="方块(20201)" value="20201"></el-option>
                        <el-option key="2" label="点击(20801)" value="20801"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="新增">
                    <el-input v-model="form.newly"></el-input>
                </el-form-item>
                <el-form-item label="次留">
                    <el-input v-model="form.tow_pr"></el-input>
                </el-form-item>
                <el-form-item label="三留">
                    <el-input v-model="form.three_pr"></el-input>
                </el-form-item>
                <el-form-item label="七留">
                    <el-input v-model="form.seven_pr"></el-input>
                </el-form-item>
                <el-form-item label="比例">
                    <el-input v-model="form.retention"></el-input>
                </el-form-item>
            </el-form>
            <span slot="footer" class="dialog-footer">
                <el-button @click="editVisible = false">取 消</el-button>
                <el-button type="primary" @click="saveEdit">确 定</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script>
/* eslint-disable */ 
    export default {
        name: 'stats',
        data() {
            return {
                tableData: [],
                cur_page: 1,
                select_daterange:[],
                is_search: false,
                editVisible: false,
                form: {
                    channel: '',
                    game: '',
                    logdate: '',
                    newly: '',
                    tow_pr: '',
                    three_pr: '',
                    seven_pr: '',
                    retention: ''                },
                idx: -1
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
                this.$store.dispatch('stats/GetStats', { user, page, stime, etime}).then((res) => {
                    this.tableData = res.list;
                }).catch(err => {
                    console.log(err)
                    this.$message.error(err);
                });
            },
            handleAdd() {
                this.form = {
                    channel: '',
                    game: '',
                    logdate: '',
                    newly: '',
                    tow_pr: '',
                    three_pr: '',
                    seven_pr: '',
                    retention: ''
                }
                this.editVisible = true;
            },
            handleEdit(index, row) {
                this.idx = index;
                const item = this.tableData[index];
                // console.log(index)
                // console.log(item)
                // let edittime = Date.parse(item.logdate.replace(/-/g,'/'))
                this.form = {
                    channel: item.channel,
                    game: item.game,
                    logdate: item.logdate,
                    newly: item.newly,
                    tow_pr: item.tow_pr,
                    three_pr: item.three_pr,
                    seven_pr: item.seven_pr,
                    retention: item.retention
                }
                this.editVisible = true;
            },
            // 保存编辑
            saveEdit() {
                const user = this.$store.getters.name
                const formdata = this.form
                console.log(formdata)
                this.$store.dispatch('stats/SetStats', { user, formdata}).then((res) => {
                    this.$set(this.tableData, this.idx, this.form);
                    //this.editVisible = false;
                    this.$message.success(`执行成功`);
                }).catch(err => {
                    console.log(err)
                    this.$message.error(err);
                });

                this.editVisible = false;
            },
        }
    }

</script>

<style scoped>
    .handle-box {
        margin-bottom: 20px;
    }

    .handle-select {
        width: 160px;
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
