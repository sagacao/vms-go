<template>
    <div class="table">
        <div class="app-container">
            <div class="handle-box">
                <el-row :gutter="10">
                    <el-col :span="4">
                        <xdh-select v-model="selected_game"> </xdh-select>
                    </el-col>
                    <el-col :span="2">
                        <el-button type="primary" icon="el-icon-search" @click="handleSearch">搜索</el-button>
                    </el-col>
                    <el-col :span="2" :offset="4">
                        <el-button type="primary"  @click="handleAdd">新增<i class="el-icon-upload el-icon--right"></i></el-button>
                    </el-col>
                </el-row>
            </div>
            <el-divider></el-divider>
            <el-divider></el-divider>
            <el-table :data="data" border class="table">
                <el-table-column prop="game" label="游戏" width="100">
                </el-table-column>
                <el-table-column prop="channel" label="渠道" width="160">
                </el-table-column>
                <el-table-column prop="appid" label="APPID" width="200">
                </el-table-column>
                <el-table-column prop="secret" label="秘钥" width="300">
                </el-table-column>
                <el-table-column prop="desc" label="描述" width="300">
                </el-table-column>
                <el-table-column label="操作" align="center">
                    <template slot-scope="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                        <el-button type="text" icon="el-icon-delete" @click="handleDel(scope.$index, scope.row)">删除</el-button>
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
                <el-form-item label="渠道">
                    <el-input v-model="form.channel">wx</el-input>
                </el-form-item>
                <el-form-item label="APPID">
                    <el-input v-model="form.appid"></el-input>
                </el-form-item>
                <el-form-item label="秘钥">
                    <el-input v-model="form.secret"></el-input>
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="form.desc"></el-input>
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
import XdhSelect from './xdhselect'
export default {
    components: { XdhSelect },
    name: 'gates',
    data() {
        return {
            tableData: [],
            cur_page: 1,
            is_search: false,
            selected_game : '',
            editVisible: false,
            actioning : "add",
            form: {
                game: this.selected_game,
                channel : '',
                appid: '',
                secret: '',
                desc: ''
            },
            idx: -1
        }
    },
    created() {

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
            this.handleSearch()
        },

        handleSearch() {
            this.is_search = true;
            const user = this.$store.getters.name
            const page = this.$route.path
            const game = this.selected_game
            console.log(this.selected_game)
            this.$store.dispatch('svrmgr/GetSvrAppInfo', { user, page, game}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    if (res.errorCode == 9100) {
                        this.$message.error('没有数据')
                    } else {
                        this.$message.error('服务器返回失败:' + res.errorCode)
                    } 
                    this.tableData = []
                } else {
                    // this.tableData = []
                    this.tableData = res.data
                    // for(var i = 0; i < res.data.length; i++) {
                    //     let item = res.data[i]
                    //     item['game'] = game
                    //     this.tableData.push(item)
                    // } 
                } 
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });
        },
        handleDel(index, row) {
            this.idx = index;
            this.$confirm('此操作将永久删除, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(() => {
                this.delCommit(index)
            }).catch((err) => {
                console.log(err)
                this.$message({
                    type: 'info',
                    message: '已取消删除' + err
                });
            });
        },
        delCommit(index) {
            const item = this.tableData[index];
            const formdata = {
                game : this.selected_game,
                appid : item.appid
            }
            // console.log(formdata)
            const user = this.$store.getters.name
            this.$store.dispatch('svrmgr/RemoveSvrAppInfo', { user, formdata}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    this.$message.error(`执行失败 :` + res.errorCode);
                } else {
                    this.$message.success(`执行成功`);
                    this.handleSearch()
                }
            }).catch(err => {
                this.$message.error(err);
            });
        },
        handleAdd() {
            if (this.selected_game === '') {
                this.$message.success(`请先选择游戏`)
                return
            }
            this.form = {
                game: this.selected_game,
                channel : 'wx',
                appid: '',
                secret : '',
                desc : ''
            }
            this.editVisible = true
            this.actioning = 'add'
        },
        handleEdit(index, row) {
            this.idx = index;
            const item = this.tableData[index];
            this.form = {
                game: this.selected_game,
                channel : item.channel,
                appid: item.appid,
                secret: item.secret,
                desc: item.desc
            }
            this.editVisible = true;
            this.actioning = 'edit'
        },
        // 保存编辑
        saveEdit() {
            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                channel : this.form.channel,
                appid : this.form.appid,
                secret : this.form.secret,
                desc: this.form.desc,
                action : this.actioning
            }
            console.log(formdata)
            this.$store.dispatch('svrmgr/SetSvrAppInfo', { user, formdata}).then((res) => {
                if (res.errorCode != 0) {
                    this.$message.error(`执行失败 :` + res.errorCode);
                } else {
                    this.$set(this.tableData, this.idx, this.form);
                    this.$message.success(`执行成功`);
                }
                //this.editVisible = false;
            }).catch(err => {
                console.log(err)
                this.$message.error(err);
            });

            this.editVisible = false;
            this.handleSearch()
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
