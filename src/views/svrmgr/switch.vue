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
            <el-table :data="data" border class="table" :default-sort = "{prop: 'funcname', order: 'ascending'}" >
                <el-table-column prop="game" label="游戏" sortable width="100">
                </el-table-column>
                <el-table-column prop="funcname" label="名字" sortable width="180">
                </el-table-column>
                <el-table-column property="funcswitch" label="状态" width="160" align="center">
                    <template slot-scope="scope">
                        <el-switch active-color="#13ce66" inactive-color="#ff4949" active-value="1" inactive-value="0" v-model="scope.row.funcswitch" @change="handleEdit(scope.$index,scope.row)">
                        </el-switch>
                    </template>
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
                <el-form-item label="名字">
                    <el-select v-model="form.funcname" placeholder="请选择" >
                        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value">
                            <span style="float: left">{{ item.label }}</span>
                            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="form.funcswitch" placeholder="请选择" >
                        <el-option v-for="item in funcstatus" :key="item.value" :label="item.label" :value="item.value">
                            <span style="float: left">{{ item.label }}</span>
                            <span style="float: right; color: #8492a6; font-size: 13px">{{ item.value }}</span>
                        </el-option>
                    </el-select>
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
    name: 'switch',
    data() {
        return {
            options:[
                {
                    value: 'invite',
                    label: 'invite'
                },
                {
                    value: 'push',
                    label: 'push'
                },
                {
                    value: 'forcar',
                    label: 'forcar'
                },
                {
                    value: 'redirect',
                    label: 'redirect'
                },
                {
                    value: 'clickTips',
                    label: 'clickTips'
                },
                {
                    value: 'unlockLevelMode',
                    label: 'unlockLevelMode'
                },
                {
                    value: 'banner',
                    label: 'banner'
                },
                {
                    value: 'mute',
                    label: 'mute'
                },
                {
                    value: 'comic',
                    label: 'comic'
                }
            ],
            funcstatus:[
                {
                    value: '1',
                    label: '开'
                },
                {
                    value: '0',
                    label: '关'
                }
            ],
            tableData: [],
            cur_page: 1,
            is_search: false,
            selected_game : '',
            editVisible: false,
            actioning : "add",
            form: {
                game: this.selected_game,
                funcname : 'push',
                funcswitch: '1'
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
            this.$store.dispatch('svrmgr/GetSvrSwitch', { user, page, game}).then((res) => {
                console.log(res)
                if (res.errorCode != 0) {
                    if (res.errorCode == 9100) {
                        this.$message.error('没有数据')
                    } else {
                        this.$message.error('服务器返回失败:' + res.errorCode)
                    }
                } else {
                    this.tableData = []
                    // this.tableData = res.data
                    for(var i = 0; i < res.data.length; i++) {
                        if (res.data[i].funcname == 'maxPumpTimes' || res.data[i].funcname == 'openVideoModal') {
                            continue
                        }
                        let item = res.data[i]
                        item['game'] = game
                        item['funcswitch'] = item.funcswitch + ''
                        this.tableData.push(item)
                    } 
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
                funcname : item.funcname,
                funcswitch: item.funcswitch
            }
            // console.log(formdata)
            const user = this.$store.getters.name
            this.$store.dispatch('svrmgr/RemoveSvrSwitch', { user, formdata}).then((res) => {
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
                funcname : 'push',
                funcswitch: '1'
            }
            this.editVisible = true
            this.actioning = 'add'
        },
        handleEdit(index, row) {
            this.idx = index;
            const item = this.tableData[index];
            this.form = {
                game: this.selected_game,
                funcname : item.funcname,
                funcswitch: item.funcswitch
            }
            // this.editVisible = true;
            this.actioning = 'edit'
            this.saveEdit()
        },
        // 保存编辑
        saveEdit() {
            const user = this.$store.getters.name
            const formdata = {
                game : this.selected_game,
                funcname : this.form.funcname,
                funcswitch: parseInt(this.form.funcswitch),
                action : this.actioning
            }
            console.log(formdata)
            this.$store.dispatch('svrmgr/SetSvrSwitch', { user, formdata}).then((res) => {
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
